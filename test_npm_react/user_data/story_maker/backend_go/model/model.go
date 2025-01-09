package model

type Story struct {
	Id     int		 `json:"id"`
	Name   string	 `json:"name" binding:"required"`
	Author string 	 `json:"author"`
}

type Model interface {
	GetStories () []Story
	AddStory (Story) Story
	DeleteStories ([]int)
}