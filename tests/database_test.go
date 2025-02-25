package tests

import (
	"deepits/internal/database"
	"os"
	"testing"
	"time"
)

// xoa file log truoc khi chay test
func cleanup() {
	os.Remove("log.txt")
}

func TestPersistence(t *testing.T) {
	cleanup()
	db := database.NewDatabase()

	//ghi du lieu vao database va file log
	db.Set("name", "khanh", 0)
	db.Set("age", "20", 0)

	//kiem tra xem du lieu co dung khong
	val, exists := db.Get("name")
	if !exists || val != "khanh" {
		t.Errorf("Expected 'khanh', got '%s'", val)
	}

	val, exists = db.Get("age")
	if !exists || val != "20" {
		t.Errorf("Expected '20', got '%s'", val)
	}

	//tao database moi de kiem tra load tu file log
	db2 := database.NewDatabase()

	val, exists = db2.Get("name")
	if !exists || val != "khanh" {
		t.Errorf("Expected 'khanh', got '%s'", val)
	}

	val, exists = db2.Get("age")
	if !exists || val != "20" {
		t.Errorf("Expected '20', got '%s'", val)
	}
}

// test SET va GET key
func TestSetAndGet(t *testing.T) {
	db := database.NewDatabase()
	db.Set("username", "khanh", 0)
	value, exists := db.Get("username")
	if !exists || value != "khanh" {
		t.Errorf("Expected 'khanh', got '%s'", value)
	}
}

// test key khong ton tai
func TestExpiration(t *testing.T) {
	db := database.NewDatabase()
	db.Set("temp", "expired", 1) //het han sau 1 giay

	time.Sleep(2 * time.Second)
	_, exists := db.Get("temp")

	if exists {
		t.Errorf("Expected key 'temp' to be expired")
	}
}

// test delete key
func TestDelete(t *testing.T) {
	db := database.NewDatabase()
	db.Set("delete_me", "test", 0)
	db.Delete("delete_me")

	_, exists := db.Get("delete_me")
	if exists {
		t.Errorf("Expected key 'delete_me' to be deleted")
	}
}
