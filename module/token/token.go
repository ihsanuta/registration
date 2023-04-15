//go:generate mockery --name=TokenModule
package token

import (
	"fmt"
	"io/ioutil"
	"log"
	"registration/app/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenModule interface {
	GenerateTokenJWT(user model.User) (string, error)
	Validate(token string) (map[string]interface{}, error)
}

type tokenModule struct{}

func NewTokenModule() TokenModule {
	return &tokenModule{}
}

type TokenData struct {
	Data model.User `json:"data"`
	jwt.RegisteredClaims
}

func (t *tokenModule) GenerateTokenJWT(user model.User) (string, error) {
	prvKey, err := ioutil.ReadFile("cert/id_rsa")
	if err != nil {
		log.Fatalln(err)
	}

	now := time.Now().UTC()
	claims := make(jwt.MapClaims)
	claims["dat"] = user                      // Our custom data.
	claims["exp"] = now.Add(time.Hour).Unix() // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix()                // The time at which the token was issued.
	claims["nbf"] = now.Unix()                // The time before which the token must be disregarded.

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(prvKey))
	if err != nil {
		return "", fmt.Errorf("error parsing RSA private key: %v", err)
	}

	// encoded string
	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return tokenString, nil
}

func (t *tokenModule) Validate(token string) (map[string]interface{}, error) {
	pubKey, err := ioutil.ReadFile("cert/id_rsa.pub")
	if err != nil {
		log.Fatalln(err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(pubKey)
	if err != nil {
		return nil, fmt.Errorf("validate: parse key: %w", err)
	}

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("validate: invalid")
	}

	data := claims["dat"].(map[string]interface{})
	return data, nil
}
