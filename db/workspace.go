package db

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Username   string
	Workspaces []*Workspace `gorm:"many2many:workspaces_members"`
}

type Workspace struct {
	gorm.Model
	Name    string    `form:"name"`
	Members []*Member `gorm:"many2many:workspaces_members"`
	Code    string    `gorm:"unique"`
}
