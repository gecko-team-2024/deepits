package tests

import (
	"deepits/internal/command"
	"deepits/internal/database"
	"testing"
)

// test set va get command
func TestSetGetCommand(t *testing.T) {
	db := database.NewDatabase()

	result := command.ExecuteCommand(db, "SET username khanh")
	if result != "OK" {
		t.Errorf("Expected 'OK', got '%s'", result)
	}

	result = command.ExecuteCommand(db, "GET username")
	if result != "khanh" {
		t.Errorf("Expected 'khanh', got '%s'", result)
	}
}

// test delete command
func TestDeleteCommand(t *testing.T) {
	db := database.NewDatabase()

	command.ExecuteCommand(db, "SET key1 value1")
	command.ExecuteCommand(db, "DEL key1")

	result := command.ExecuteCommand(db, "GET key1")
	if result != "(nil)" {
		t.Errorf("Expected '(nil)', got '%s'", result)
	}
}

// test unknown command
func TestUnknownCommand(t *testing.T) {
	db := database.NewDatabase()

	result := command.ExecuteCommand(db, "INVALID_COMMAND")
	expected := "ERR: Unknown command"
	if result != expected {
		t.Errorf("Expected 'ERR: Unknown command', got '%s'", result)
	}
}
