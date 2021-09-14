package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"wm/workspace/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()

		latency := time.Since(t)
		log.Print(latency)
		status := c.Writer.Status()
		log.Print(status)
		path := c.FullPath()
		log.Println(path)
		method := c.Request.Method
		log.Println(method)
		data := url.Values{
			"api":     {path},
			"status":  {strconv.Itoa(status)},
			"latency": {latency.String()},
			"method":  {method},
		}

		resp, err := http.PostForm(config.LOGGER_SERVICE+"/log", data)

		if err != nil {
			log.Fatal(err)
		}

		var res map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&res)
	}
}

func JwtFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenString = tokenString[7:]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.SECRET), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims)
		} else {
			fmt.Println(err.Error())
		}
		c.Next()
	}
}
