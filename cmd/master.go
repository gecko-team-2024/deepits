package main

import (
	"deepits/internal/database"
	"deepits/internal/server"
)

func main() {
	db := database.NewDatabase()
	server.StartServer(db)
}
