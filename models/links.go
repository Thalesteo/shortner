package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	Name    string
	Url     string
	Qrcode  string
	Status  bool
	GroupId int `gorm:"references:ID"`
}
