package tools

import (
	"encoding/binary"
	"fmt"
)

var index int = 0
var recvBytes []byte

func Unpack(recvBytesG []byte) string {
	recvBytes = recvBytesG
	readHeader()
	fmt.Println(header.DataLength)
	fmt.Println(header.Service)

	readBody()
	return ""
}

func readBody() {
	if header.Service == 1 {
		readLogin()
		runLogin()
	} else if header.Service == 2 {
		ret := readChat()

		fmt.Println(ret)
		index = 0
	}
}

func readInt() int {
	ret := binary.BigEndian.Uint32(recvBytes[index:])
	index += 4
	return int(ret)
}

func readShort() int16 {
	ret := binary.BigEndian.Uint16(recvBytes[index:])
	index += 2
	return int16(ret)
}

func readString() string {
	strLength := int(readShort())
	str := recvBytes[index : index+strLength]
	index += strLength

	return string(str)
}
