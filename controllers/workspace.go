package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"wm/workspace/config"

	"github.com/gin-gonic/gin"
)

type CheckResponse struct {
	message bool
}

func ListWorkspace(c *gin.Context) {
	userId, _ := c.GetQuery("userId")
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	jwt := c.GetHeader("Authorization")
	fmt.Println("jwt: ", jwt)
	req, err := http.NewRequest("GET", config.ACCOUNT_SERVICE+"/user/exists?userId="+userId, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Add("Authorization", jwt)
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()

	var checkResponse map[string]interface{}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(body), &checkResponse)

	fmt.Println("message: ", checkResponse)
	if response.StatusCode != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong userId",
			"body":    userId,
		})
		return
	}

	// var workspaces []db.Workspace

	c.JSON(http.StatusOK, gin.H{
		"message": "test",
		"body":    userId,
	})
}

func CreateWorkspace(c *gin.Context) {

}
