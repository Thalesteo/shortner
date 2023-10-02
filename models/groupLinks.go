package models

import (
	"time"

	"gorm.io/gorm"
)

type GroupLinks struct {
	gorm.Model
	Name    string
	Url     string
	Status  bool
	UserId  int `gorm:"references:ID"`
	Links   []Link
	MaxTime time.Time
}
