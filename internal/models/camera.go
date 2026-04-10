package models

import "gorm.io/gorm"

type Camera struct {
	gorm.Model
	Name     string `json:"name"`
	RTSP     string `json:"rtsp"`
	Username string `json:"username"`
	Password string `json:"password"`
	IP       string `json:"ip"`
	Status   string `json:"status"`
}
