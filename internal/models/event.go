package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name     string `json:"name"`
	Checked  bool   `json:"checked"`
	CameraID uint   `gorm:"index" json:"cameraID"`
	Camera   Camera `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"camera"`
}
