package main

import (
	"fmt"
	"io"
	"net"

	"../tools"
)

var (
	IP               string
	PORT             string
	CurrentDirectory string
	ClientPool       map[string]Client
)

// map[IP]Client {ID, Conn}
type Client struct {
	UserID string
	Conn   net.Conn
}

func main() {
	IP = "10.131.150.171"
	PORT = ":13302"
	ln, err := net.Listen("tcp", IP+PORT)
	fmt.Println("chat server is working now...")

	defer ln.Close()
	if err != nil {
		panic(err)
	}
	ClientPool = make(map[string]Client)
	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Printf("Accept Error; error : %v", err)
			continue
		}

		go Handler(conn)
	}

}

func Handler(conn net.Conn) {
	var recvBuf []byte
	defer conn.Close()
	address := conn.RemoteAddr().String()
	client := Client{
		UserID: "",
		Conn:   conn,
	}
	ClientPool[address] = client

	fmt.Printf("ClientPool : %v\n", ClientPool)
	fmt.Printf("Connected Client address: %s \n", address)
	for {
		recvBuf = make([]byte, 0, 4096)
		tmp := make([]byte, 256)
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("read error : %v", err)
			}
			break // if the error is End of File error, then break
		}
		recvBuf = append(recvBuf, tmp[:n]...)

		msg := tools.Unpack(recvBuf)
		fmt.Printf("%v : %s", address, msg)
		for k, v := range ClientPool { //broad casting
			fmt.Printf("key : %s, value : %v", k, v.UserID)
			v.Conn.Write(recvBuf)
		}
	}

	defer delete(ClientPool, address)
}
