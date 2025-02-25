package database

import (
	"fmt"
	"net"
	"os"
	"sync"
)

// dinh ngia mot item trong database
type Item struct {
	Value      string
	Expiration int64
}

//#Replication (Sao chep du lieu) giup mot database chinh (Master) dong bo du lieu voi
//mot hoac nhieu database phu (Slave). Dieu nay giup:
//1. Tang kha nang chiu tai: Doc du lieu tu nhieu Slave thay vi chi mot Master
//2. Cai thien do tin cay: Neu Master bi loi, mot Slave khac co the tro thanh Master
//3. Backup du lieu: Tranh mat du lieu neu mot node bi hong

//Dung TCP de dong bo du lieu giua Master va cac Slave

// dinh nghia database (map + mutex)
type Database struct {
	Data   map[string]Item
	slaves []net.Conn // danh sach cac slaves
	Mu     sync.RWMutex
}

// khoi tao database moi
func NewDatabase() *Database {
	db := &Database{Data: make(map[string]Item)}
	db.LoadFromFile()
	return db
}

// luu du lieu vao database va gui den cac Slave
func (db *Database) Set(key string, value string, ttl int64) {
	db.Mu.Lock()
	defer db.Mu.Unlock()
	db.Data[key] = Item{Value: value, Expiration: 0}
	db.AppendToFile(fmt.Sprintf("SET %s %s\n", key, value))
	db.replicate(fmt.Sprintf("SET %s %s\n", key, value))
}

// lay du lieu tu database
func (db *Database) Get(key string) (string, bool) {
	db.Mu.RLock()
	defer db.Mu.RUnlock()
	val, exists := db.Data[key]
	return val.Value, exists
}

// xoa key khoi database + gui den Slave
func (db *Database) Delete(key string) {
	db.Mu.Lock()
	defer db.Mu.Unlock()
	delete(db.Data, key)
	db.AppendToFile(fmt.Sprintf("DEL %s\n", key))
	db.replicate(fmt.Sprintf("DEL %s\n", key))
}

func (db *Database) AppendToFile(command string) {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()
	file.WriteString(command)
}

func (db *Database) LoadFromFile() {
	file, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	var cmd, key, value string
	for {
		_, err := fmt.Fscanf(file, "%s %s %s\n", &cmd, &key, &value)
		if err != nil {
			break
		}
		if cmd == "SET" {
			db.Data[key] = Item{Value: value, Expiration: 0}
		}
	}
}

// them Slave vao he thong
func (db *Database) AddSlave(slaveAddr string) {
	conn, err := net.Dial("tcp", slaveAddr)
	if err != nil {
		fmt.Println("Cannot connect to slave:", err)
		return
	}
	db.slaves = append(db.slaves, conn)
}

// gui lenh den tat ca cac Slave
func (db *Database) replicate(command string) {
	for _, conn := range db.slaves {
		_, err := conn.Write([]byte(command))
		if err != nil {
			fmt.Println("Error sending to slave:", err)
		}
	}
}
