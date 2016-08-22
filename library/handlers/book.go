package handlers

import (
	"github.com/alejopm92/GoProjects/library/repository"
	"github.com/kataras/iris"
)

func GetAll(c *iris.Context) {
	libraryRepository := repository.BooksRepository{}
	books, err := libraryRepository.GetAll()

	//Manejo de errores
	if err != nil {

	}
	c.JSON(200, books)
}
