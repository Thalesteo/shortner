package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required, uuid, primarykey"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	Name      string    `db:"name" json:"name" validate:"min=3,max=32"`
	Email     string    `db:"email" json:"email" validate:"min=9,max=256"`
	Password  string    `db:"password" json:"password" validate:"min=8"`
}
