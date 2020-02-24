package tools

import (
	"encoding/binary"
	"fmt"
)

var index int = 0
var recvBytes []byte
var UserID string

//[]byte 초기화
func Unpack(recvBytesG []byte) (string, bool, string) {
	recvBytes = recvBytesG
	readHeader()
	//execute specific function by Service

	if header.Error == 1 {
		index = 0
		return "Error", false, "Error"
	}

	if header.Service == 1 {
		// Service 1 is the login function.
		readLogin()
		var res bool
		var code int
		UserID = login.UserID
		res, code = runLogin()
		index = 0
		return LoginCode[code], res, "Login"
	} else if header.Service == 2 {
		// Service 2 is the chat function
		msg := readChat()
		fmt.Println(msg)
		index = 0
		return msg, true, "Chat"
	} else if header.Service == 0 {
		msg := readChat()
		fmt.Println(msg)
		index = 0
		return msg, false, "Chat"
	}

	return "", false, ""
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

func readByte() byte {
	ret := recvBytes[index : index+1]

	return ret[0]

}

func setError() {

}
