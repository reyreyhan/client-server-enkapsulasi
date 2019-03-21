package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

type WinLose struct {
	Hero_id int32
	WinLose bool
}

func main() {
	fmt.Println("Client")
	//Membuat Struct
	winLoseEncode := WinLose{Hero_id: 40, WinLose: true}

	fmt.Println("start client")

	//Koneksi TCP ke Server
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("Connection error", err)
	}

	//Encode Struct
	encoder := gob.NewEncoder(conn)
	encoder.Encode(winLoseEncode)

	//Decode Struct Dari Server
	dec := gob.NewDecoder(conn)
	winLose := &WinLose{}
	dec.Decode(winLose)
	fmt.Println("Hero ID : ", winLose.Hero_id, " - ", winLose.WinLose)
	fmt.Println("done")
}
