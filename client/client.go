package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"

	"../tools"
	"./menu"
)

var (
	IP      string
	PORT    string
	recvBuf []byte
)

func main() {
	IP = "10.131.150.171"
	PORT = ":13302"
	conn, err := net.Dial("tcp", IP+PORT)

	defer conn.Close()

	if err != nil {
		panic(err)
	}
	for {
		packedLoginData := menu.DisplayMenu()

		conn.Write(packedLoginData)

		loginBuf := make([]byte, 4096)

		n, err := conn.Read(loginBuf)
		if err != nil {
			panic(err)
		}

		msg, res, service := tools.Unpack(loginBuf[:n])
		fmt.Println("MSG from server : ", msg)
		fmt.Println("from server res : ", res)
		fmt.Println("Server : ", service)

		if res == true {
			break
		}
	}

	go RecvFunc(conn)
	for {

		fmt.Printf("client : ")
		in := bufio.NewReader(os.Stdin)
		data, err := in.ReadString('\n')
		if err != nil {
			panic(err)
		}

		packedData := tools.Pack(data, 0, 2, 0)
		fmt.Println(packedData)
		conn.Write(packedData)
		in.Reset(os.Stdin)
		// time.Sleep(2 * time.Second)
	}
}

func RecvFunc(conn net.Conn) {
	for {
		recvBuf = make([]byte, 0, 4096)
		tmp := make([]byte, 256)
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("read error; error : %v \n", err)
			}
			break
		}
		recvBuf = append(recvBuf, tmp[:n]...)

		msg, res, service := tools.Unpack(recvBuf)
		fmt.Println("from server res : ", res)
		fmt.Println("Server : ", service)
		if res == false {
			packedLoginData := menu.DisplayMenu()

			conn.Write(packedLoginData)
		}

		fmt.Printf("total size : %d\n", len(recvBuf))
		fmt.Printf("data from server : %s", msg)

	}

}
