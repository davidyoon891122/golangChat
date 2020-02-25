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

	// loginFunc(conn)

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

		msg, res, service := tools.Unpack(recvBuf)

		fmt.Println("Unpacked res:", res)
		fmt.Println("MSG to client : ", msg)
		if service == "Login" {
			if res == true {

				fmt.Println("UserID : ", tools.UserID)
				client := Client{
					UserID: tools.UserID,
					Conn:   conn,
				}
				ClientPool[address] = client
				packedData := tools.Pack(msg, 0, 2, 0)
				conn.Write(packedData)
			} else if res == false {
				packedData := tools.Pack("login failed", 0, 0, 1)
				conn.Write(packedData)

			} else {
				fmt.Println("Login res exception")
			}

		} else if service == "Error" {
			packedData := tools.Pack("", 0, 0, 1)
			conn.Write(packedData)

		} else if service == "Chat" {
			fmt.Println(res)

			fmt.Printf("%v : %s", ClientPool[address].UserID, msg)
			msgWithID := fmt.Sprintf("%v : %s", ClientPool[address].UserID, msg)
			for k, v := range ClientPool { //broad casting
				fmt.Printf("key : %s, value : %v", k, v.UserID)
				packedData := tools.Pack(msgWithID, 0, 2, 0)
				v.Conn.Write(packedData)
			}
		}
	}
	defer fmt.Printf("client is disconnected : %v\n", ClientPool[address].UserID)
	defer delete(ClientPool, address)
}

// func loginFunc(conn net.Conn) {
// 	var msg string
// 	fmt.Println("Login Process")
// 	msg = "Please send your ID and password"
// 	packedMsg := tools.Pack(msg, 0, 1)
// 	conn.Write(packedMsg)
// }
