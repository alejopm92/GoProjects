package main

import (
	"github.com/alejopm92/GoProjects/library/handlers"
	"github.com/kataras/iris"
	"database/sql"
	"github.com/alejopm92/GoProjects/library/repository"
)

func main() {

	db, err := sql.Open("mysql", "user:password@/database")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	repository.Connection = db

	iris.Get("/api/library/", handlers.GetAll)
	iris.Listen(":9090")
}
