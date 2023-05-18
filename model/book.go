package model

import "time"

type Book struct {
	BookId    uint `gorm:"primaryKey"`
	Title     string
	AuthorId  uint
	Author    Author `gorm:"foreignKey:AuthorId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
