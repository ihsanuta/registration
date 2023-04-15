package handler

import (
	"fmt"
	"net/http"
	"registration/app/model"
	"registration/module/internalvalidate"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Registrasi godoc
// @Summary Registrasi
// @Schemes
// @Tags User
// @Description Registrasi
// @Accept json
// @Param   requestBody  body   model.UserRequest   true  "Registrasi User Request Body"
// @Produce json
// @Success 201 {object} model.Response
// @Router /api/v1/registrasi [post]
func (h *handler) Register(c *gin.Context) {
	var form model.UserRequest

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(form)
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		msgerr := err.(validator.ValidationErrors)[0]

		// from here you can create your own error messages in whatever language you wish
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%s error %s %s", msgerr.Field(), msgerr.Tag(), msgerr.Param())})
		return
	}

	err = internalvalidate.ValidatePassword(form.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = internalvalidate.ValidatePhone(form.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.usecase.User.Register(form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, model.Response{
		Status:  http.StatusOK,
		Message: "success",
	})
}

// Login godoc
// @Summary Login
// @Schemes
// @Tags User
// @Description Login
// @Accept json
// @Param   requestBody  body   model.Login   true  "Login User Request Body"
// @Produce json
// @Success 200 {object} model.Response
// @Router /api/v1/login [post]
func (h *handler) Login(c *gin.Context) {
	var form model.Login

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := internalvalidate.ValidatePhone(form.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.usecase.User.Login(form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    token,
	})
}

// Get Data User godoc
// @Summary Get Data User
// @Schemes
// @Tags User
// @Description Get Data User
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} model.Response
// @Router /api/v1/user [get]
func (h *handler) GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	user, ok := ctx.Value("user-session").(map[string]interface{})
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session invalid"})
		return
	}

	phone := user["phone"].(string)
	token, err := h.usecase.User.GetByPhone(phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    token,
	})
}

// Login godoc
// @Summary Login
// @Schemes
// @Tags User
// @Description Login
// @Accept json
// @Security ApiKeyAuth
// @Param   requestBody  body   model.UserUpdate   true  "Update User Request Body"
// @Produce json
// @Success 200 {object} model.Response
// @Router /api/v1/user [put]
func (h *handler) Update(c *gin.Context) {
	var form model.UserUpdate

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(form)
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		msgerr := err.(validator.ValidationErrors)[0]

		// from here you can create your own error messages in whatever language you wish
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%s error %s %s", msgerr.Field(), msgerr.Tag(), msgerr.Param())})
		return
	}

	ctx := c.Request.Context()
	user, ok := ctx.Value("user-session").(map[string]interface{})
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session invalid"})
		return
	}

	phone := user["phone"].(string)
	token, err := h.usecase.User.UpdateName(phone, form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    token,
	})
}
