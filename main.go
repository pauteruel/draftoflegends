package main

import (
	"log"
	"net/http"
)

const (
	host     = "localhost"
	port     = 32768
	user     = "postgres"
	password = "ninofaker4"
	dbname   = "draftoflegends"
)
var dataBase *DataBase

func main() {

	dataBase =NewDataBaseInstance(host,port,user,password,dbname)
	error := dataBase.Init()
	if error != nil {
		panic(error)
	}
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
