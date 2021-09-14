package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"wm/workspace/db"

	"github.com/gin-gonic/gin"
)

func ListWorkspace(c *gin.Context) {
	var member db.Member
	username, err := getUserIdFromJWT(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "jwt user id is wrong",
			"body":    "",
		})
		return
	}
	db.DB.Where("username = ?", username).First(&member)

	var workspaces []db.Workspace
	db.DB.Model(&member).Preload("Members").Association("Workspaces").Find(&workspaces)
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
		"body":    workspaces,
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
	fmt.Println(workspaceId)
	var workspace db.Workspace
	if err := db.DB.Where("ID = ?", workspaceId).First(&workspace).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong workspace Id",
			"body":    workspaceId,
		})
		return
	}
	db.DB.Delete(&workspace)
	c.JSON(http.StatusOK, gin.H{
		"message": "workspace deleted",
		"body":    workspace,
	})
}

func CreateInviteCode(c *gin.Context) {
	workspaceId := c.PostForm("workspaceId")
	var workspace db.Workspace

	err := db.DB.Where("ID = ?", workspaceId).First(&workspace).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong workspace id",
			"body":    nil,
		})
		return
	}
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(8999) + 1000
	workspace.Code = fmt.Sprint(code)
	err = db.DB.Save(&workspace).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	for err != nil {
		rand.Seed(time.Now().UnixNano())
		code := rand.Intn(8999) + 1000
		workspace.Code = fmt.Sprint(code)
		err = db.DB.Create(&workspace).Error
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "invited code generated",
		"body":    workspace.Code,
	})
}

func AcceptInvite(c *gin.Context) {
	inviteCode := c.PostForm("code")
	var workspace db.Workspace
	err := db.DB.Where("code = ?", inviteCode).First(&workspace).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"meesage": "wrong code",
			"body":    inviteCode,
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
	err = db.DB.Where("username = ?", username).First(&member).Error
	if err != nil {
		member = db.Member{Username: username}
		db.DB.Create(&member)
	}
	workspace.Members = append(workspace.Members, &member)
	db.DB.Save(&workspace)
	c.JSON(http.StatusOK, gin.H{
		"message": "invite accepted",
		"body":    workspace,
	})
}
