package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"wm/workspace/config"
	"wm/workspace/db"

	"github.com/gin-gonic/gin"
)

func ListWorkspace(c *gin.Context) {
	userId, _ := c.GetQuery("userId")
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	jwt := c.GetHeader("Authorization")
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

	if response.StatusCode != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong userId",
			"body":    userId,
		})
		return
	}
	var member db.Member
	username, err := getUserIdFromJWT(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "jwt user id is wrong",
			"body":    "",
		})
		return
	}
	db.DB.Preload("Workspaces").Where("username = ?", username).First(&member)

	c.JSON(http.StatusOK, gin.H{
		"message": "test",
		"body":    member.Workspaces,
	})
}

func CreateWorkspace(c *gin.Context) {
	var workspace db.Workspace
	c.Bind(&workspace)
	username, err := getUserIdFromJWT(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "jwt user id is wrong",
			"body":    "",
		})
		return
	}
	var member db.Member
	if err := db.DB.Where("username = ? ", username).First(&member).Error; err != nil {
		member = db.Member{Username: username}
		db.DB.Create(&member)
	}

	workspace.Members = append(workspace.Members, &member)
	db.DB.Create(&workspace)

	c.JSON(http.StatusOK, gin.H{
		"message": workspace.Name + " created",
		"body":    workspace,
	})
}

func CheckWorkspace(c *gin.Context) {
	workspaceId, ok := c.GetQuery("workspaceId")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "workspace Id is required",
			"body":    "",
		})
		return
	}

	var workspace db.Workspace
	var check bool

	if err := db.DB.Where("ID = ?", workspaceId).First(&workspace).Error; err != nil {
		check = false
	} else {
		check = true
	}

	c.JSON(http.StatusOK, gin.H{
		"message": check,
	})
}

func DeleteWorkspace(c *gin.Context) {
	workspaceId := c.Param("workspaceId")
	var workspace db.Workspace
	if err := db.DB.Where("ID = ?", workspaceId).First(&workspace).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong sprint Id",
			"body":    workspaceId,
		})
		return
	}
	db.DB.Delete(&workspace)
	c.JSON(http.StatusOK, gin.H{
		"message": "sprint deleted",
		"body":    workspace,
	})
}
