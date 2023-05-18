package model

import "time"

type Author struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
