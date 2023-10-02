package models

import "time"

type User struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Email      *string
	Password   string
	GroupLinks []GroupLinks
	CreatedAt  time.Time
}
