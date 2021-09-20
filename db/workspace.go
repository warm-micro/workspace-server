package db

import (
	"gorm.io/gorm"
)

// 워크스페이스 맴버
type Member struct {
	gorm.Model
	Username   string
	Workspaces []*Workspace `gorm:"many2many:workspaces_members"`
}

// 워크스페이스 
type Workspace struct {
	gorm.Model
	Name    string    `form:"name"`
	Members []*Member `gorm:"many2many:workspaces_members"`
	Code    string    `gorm:"unique"`
}
