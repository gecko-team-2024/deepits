package main

import (
	"deepits/internal/database"
	"fmt"
	"net"
	"sync"
)

//Cau truc du an dang lam viec voi Replication, tuc la Slave se lien tuc nhan du lieu tu Master.
//Dieu nay co the tao ra nhieu string buffet, connection object gay ap luc len bo nho va GC

//#Huong toi uu
//- Tai su dung buffet khi doc lenh tu Master
//- Tai su dung connection object trong he thong

//Trong Slave, thay vi tao buffet moi cho moi lan doc lenh, ta se su dung sync.Pool de tai su dung bo nho

// Pool tai su dung buffet string
var buffetPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1024) //tao buffet moi khi can
	},
}

func main() {
	db := database.NewDatabase()

	//ket noi den Master (co the thay doi dia chi IP ne can)
	masterConn, err := net.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Cannot connect to master:", err)
		return
	}
	defer masterConn.Close()

	fmt.Println("Connected to Master at 127.0.0.1:6379")

	for {
		buffet := buffetPool.Get().([]byte) //lay buffet tu Pool
		defer buffetPool.Put(buffet)        //tra bufet ve Pool sau khi dung

		n, err := masterConn.Read(buffet) //doc du lieu tu buf

		if err != nil {
			fmt.Println("Error reading from master:", err)
			return
		}
		//giai ma du lieu thanh lenh
		var cmd, key, value string
		fmt.Sscanf(string(buffet[:n]), "%s %s %s\n", cmd, key, value)

		//xu ly lenh
		if cmd == "SET" {
			db.Set(key, value, 0)
		}
	}

	// var cmd, key, value string
	// for {
	// 	_, err := fmt.Fscanf(masterConn, "%s %s %s\n", cmd, key, value)
	// 	if err != nil {
	// 		return
	// 	}
	// 	if cmd == "SET" {
	// 		db.Set(key, value, 0)
	// 	}
	// }
}
