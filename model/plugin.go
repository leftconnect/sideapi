package model

import "time"

type Plugin struct {
	ID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
