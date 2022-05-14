package main

import (
	"github.com/alefejonatha/Http/restapi/database"
	"github.com/alefejonatha/Http/restapi/server"
)

func main() {
	database.StartDB()
	defer database.CloseConn()

	server := server.NewServer()
	server.Run()
}
