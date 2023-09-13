package main

import (
	"bases31/api"
	"bases31/storage/mongodb"
	"net/http"
)

type server struct {
	api *api.API
}

func main() {
	var srv server
	//db1 := memdb.New()
	db2 := mongodb.InitDb()
	//db3 := postgresdb.InitDb()

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(db2)
	http.ListenAndServe(":8080", srv.api.Router())
}
