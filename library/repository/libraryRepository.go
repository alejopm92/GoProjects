package repository

import (
	"github.com/alejopm92/GoProjects/library/models"
	"log"
)

type BooksRepository struct {
}

func (l *BooksRepository) GetAll() ([]models.Book, error) {
	books := make([]models.Book, 0)
	rows, err := Connection.Query("SELECT * FROM book")
	if err!= nil {
		return books,err
	}

	for rows.Next() {
		book := models.Book{}
		err = rows.Scan( &book.Id, &book.Nombre, &book.Autor, &book.Editorial, &book.Fecha)

		if err != nil{
			log.Fatal("No se pudo mi amigo", err)
		}
		books=append(books, book)
	}

	return books, nil
}
