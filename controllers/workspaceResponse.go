package controllers

import "wm/workspace/db"

type MemberResponse struct {
	Id          uint
	Username    string
	Nickname    string
	Email       string
	PhoneNumber string
}

type WorksapceResponse struct {
	Id      uint
	Name    string
	Members []MemberResponse
}

func NewWorkspaceResponse(workspace db.Workspace) *WorksapceResponse {
	return &WorksapceResponse{Id: workspace.ID, Name: workspace.Name, Members: []MemberResponse{}}
}
