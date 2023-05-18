package usecase

import (
	"myapp/model"
)

func CreateAuthor(name string, age int) model.Author {
	author := model.Author{}
	author.Name = name
	author.Age = age
	model.DB.Create(&author)
	return author
}

func GetAuthor(id int) (error, model.Author) {
	author := model.Author{}
	err := model.DB.First(&author, uint(id)).Error
	return err, author
}
