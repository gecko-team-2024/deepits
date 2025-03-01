package server

import (
	"deepits/internal/database"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
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

// chay http server
func StartHTTPServer(db *database.Database) {
	r := gin.Default()

	r.GET("/get/:key", func(c *gin.Context) {
		key := c.Param("key")
		value, exists := db.Get(key)
		if exists {
			c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
		}
	})

	r.POST("/set", func(c *gin.Context) {
		var json struct {
			Key   string `json:"key" binding:"required"`
			Value string `json:"value" binding:"required"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Set(json.Key, json.Value, 0)
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	r.DELETE("/delete/:key", func(c *gin.Context) {
		key := c.Param("key")
		db.Delete(key)
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	r.StaticFile("/", "./web/index.html") // Phục vụ tệp index.html
	r.Run(":8080")                        // Chạy HTTP server trên cổng 8080
}
