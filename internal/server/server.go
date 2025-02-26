package server

import (
	"deepits/internal/database"
	"fmt"
	"net"
)

// chay tcp server
func StartServer(db *database.Database) {
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is running on :6379")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go handleConnection(conn, db)
	}
}

// xu ly ket noi tu client
func handleConnection(conn net.Conn, db *database.Database) {
	defer conn.Close()

	var cmd, key, value string
	for {
		_, err := fmt.Fscanf(conn, "%s %s %s\n", &cmd, &key, &value)
		if err != nil {
			return
		}
		if cmd == "SET" {
			db.Set(key, value, 0)
		}
		if cmd == "GET" {
			db.Set(key, value, 0)
		}
	}
}
