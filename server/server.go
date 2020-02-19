package main

import (
	"fmt"
	"io"
	"net"
)

var (
	IP               string
	PORT             string
	CurrentDirectory string
	ClientPool       []Client
)

type Client struct {
	Address string
	Conn    net.Conn
}

func main() {
	IP = "10.131.150.171"
	PORT = ":13302"

	ln, err := net.Listen("tcp", IP+PORT)

	defer ln.Close()
	if err != nil {
		panic(err)
	}

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
	client := conn.RemoteAddr().String()
	clientS := Client{
		Address: client,
		Conn:    conn,
	}
	ClientPool = append(ClientPool, clientS)

	fmt.Printf("ClientPool : %v\n", ClientPool)
	fmt.Printf("Connected Client : %s \n", client)
	recvBuf = make([]byte, 0, 4096)
	tmp := make([]byte, 256)
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("read error : %v", err)
			}
			break // if the error is End of File error, then break
		}
		recvBuf = append(recvBuf, tmp[:n]...)
		fmt.Printf("total size : %d \n", len(recvBuf))
		fmt.Printf("data from client: %s", string(recvBuf))

		conn.Write(recvBuf)

	}

}

func DeleteClinet() {

}
