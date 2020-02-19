package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"

	"time"
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

	go RecvFunc(conn)

	for {
		fmt.Printf("client : ")
		in := bufio.NewReader(os.Stdin)
		data, err := in.ReadString('\n')
		if err != nil {
			panic(err)
		}

		conn.Write([]byte(data))
		in.Reset(os.Stdin)
		time.Sleep(2 * time.Second)
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

		fmt.Printf("total size : %d\n", len(recvBuf))
		fmt.Printf("data from server : %s", string(recvBuf))

	}

}
