package controllers

import (
	"fmt"
	"net/http"
	"wm/workspace/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func getUserIdFromJWT(tokenString string) (string, error) {
	tokenString = tokenString[7:]
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.SECRET), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	username, _ := claims["sub"].(string)

	return username, nil
}
