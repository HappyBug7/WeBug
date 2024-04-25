package model

type Class struct {
	Name     string `json:"name" form:"name"`
	Day      int    `json:"day" form:"day"`
	Class    int    `json:"class" form:"class"`
	Duration string `json:"duration" form:"duration"`
	Place    string `json:"place" form:"place"`
}
