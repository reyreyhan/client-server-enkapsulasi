package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type WinLose struct {
	Hero_id int32
	WinLose bool
}

func main() {
	fmt.Println("Server")

	//Membuat Koneksi Server TCP
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	//Cek Koneksi Sudah Terhubung / Belum
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Error")
	}

	//Decode Kiriman Dari Client
	dec := gob.NewDecoder(conn)
	winLose := &WinLose{}
	dec.Decode(winLose)
	fmt.Println("Hero ID : ", winLose.Hero_id, " - ", winLose.WinLose)

	//Encode Untuk Mengirim Ke Client
	sendToClient := winLose
	encoder := gob.NewEncoder(conn)
	encoder.Encode(sendToClient)
}
