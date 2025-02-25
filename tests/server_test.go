package tests

import (
	"bufio"
	"deepits/internal/database"
	"deepits/internal/server"
	"net"
	"testing"
	"time"
)

// ket noi den server va gui lenh
func sendCommand(t *testing.T, command string) string {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(command + "\n"))
	if err != nil {
		t.Fatalf("Failed to send command: %v", err)
	}

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	return response
}

// test server set va get
func TestServerSetGet(t *testing.T) {
	go func() {
		db := database.NewDatabase()
		server.StartServer(db) // chay server trong goroutine
	}()
	time.Sleep(5 * time.Second) // doi server khoi dong

	// gui lenh SET & GET
	setResp := sendCommand(t, "SET name khanh")
	if setResp != "OK\n" {
		t.Errorf("Expected 'OK', got '%s'", setResp)
	}

	getResp := sendCommand(t, "GET name")
	if getResp != "khanh\n" {
		t.Errorf("Expected 'khanh', got '%s'", getResp)
	}
}

// test delete command qua TCP
func TestServerDelete(t *testing.T) {
	go func() {
		db := database.NewDatabase()
		server.StartServer(db) // chay server trong goroutine
	}()
	time.Sleep(2 * time.Second) // doi server khoi dong

	sendCommand(t, "SET key1 value1")
	sendCommand(t, "DEL key1")

	getResp := sendCommand(t, "GET key1")
	if getResp != "(nil)\n" {
		t.Errorf("Expected '(nil)', got '%s'", getResp)
	}
}
