package error

import "github.com/gin-gonic/gin"

var ValidationError = []string{}

func HandleErr(c *gin.Context, err error) error {
	if err != nil {
		c.Error(err)
	}
	return err
}
