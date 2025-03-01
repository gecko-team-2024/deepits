package main

import (
	"deepits/internal/database"
	"deepits/internal/server"
)

func main() {
	db := database.NewDatabase()
	go server.StartHTTPServer(db)
	server.StartServer(db)
}
