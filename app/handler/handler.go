package handler

import (
	"context"
	"fmt"
	"net/http"
	usecase "registration/app/usecase"
	"registration/config"
	tk "registration/module/token"
	"strings"
	"sync"

	_ "registration/docs"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler interface{}

var once = &sync.Once{}

type handler struct {
	usecase *usecase.Usecase
}

// @title           Auth Service Swagger API
// @version         1.0
// @description     Auth Service Swagger API
// @termsOfService  http://swagger.io/terms/
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /

func Init(usecase *usecase.Usecase) Handler {
	var h *handler
	once.Do(func() {
		h = &handler{
			usecase: usecase,
		}
		h.Serve()
	})
	return h
}

func (h *handler) Serve() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	group := router.Group("/api/v1")
	group.POST("/registrasi", h.Register)
	group.POST("/login", h.Login)
	group.GET("/user", h.authenticateMiddleware, h.GetUser)
	group.PUT("/user", h.authenticateMiddleware, h.Update)

	serverString := fmt.Sprintf("%s:%s", config.AppConfig["host"], config.AppConfig["port"])
	router.Run(serverString)
}

var tkn = tk.NewTokenModule()

func (h *handler) authenticateMiddleware(c *gin.Context) {
	tokenString := GetTokenFromGinContext(c)

	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token required"})
		c.Abort()
		return
	}

	session, err := tkn.Validate(tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if session == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}

	// user := session.(model.User)
	currentContext := context.WithValue(c.Request.Context(), "user-session", session)
	c.Request = c.Request.WithContext(currentContext)
	c.Next()
}

func GetTokenFromGinContext(c *gin.Context) string {
	authorizationHeader := c.GetHeader("authorization")

	authorizationValues := strings.SplitN(authorizationHeader, " ", 2)

	if len(authorizationValues) < 2 || strings.ToLower(authorizationValues[0]) != "bearer" {
		return ""
	}

	return strings.TrimSpace(authorizationValues[1])
}
