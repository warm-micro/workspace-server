package db

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	UserId uint
}

type Workspace struct {
	gorm.Model
	Name    string
	Members []*Member `gorm:"many2many:workspaces_members"`
}
