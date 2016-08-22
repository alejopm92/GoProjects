package models

//Modelo del libros

type Book struct {

	Id        string `json:"id"`
	Nombre    string `json:"nombre"`
	Autor     string `json:"autor"`
	Fecha     string `json:"fecha"`
	Editorial string `json:"editorial"`

}
