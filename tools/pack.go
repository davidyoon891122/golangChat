package tools

import (
	"bytes"
	"encoding/binary"
)

var headerLength int = 13
var header *Header
var chat *Chat
var login *Login
var dataLength int
var (
	dataBuffer  bytes.Buffer
	headerBytes []byte
	bodyBytes   []byte
)

func PackHeader() {
	packInt(header.Length)
	packInt(header.DataLength)
	packShort(header.Process)
	packShort(header.Service)
	headerBytes = append(headerBytes, dataBuffer.Bytes()...)
}

func Pack(data interface{}, process int, service int, err int) []byte {
	var totalBytes []byte

	switch data.(type) {
	case string:
		chat = InitChat()
		chat.chatPacker(data.(string))
		chatWrap()
		header = InitHeader()
		header.headerPacker(headerLength, dataLength, process, service, err)
		PackHeader()
		totalBytes = append(headerBytes, bodyBytes...)
		headerBytes = nil
		bodyBytes = nil
		dataBuffer.Reset()
		return totalBytes
	case *Login:
		login = InitLogin()
		login = data.(*Login)
		loginWrap()
		header = InitHeader()
		header.headerPacker(headerLength, dataLength, process, service, err)
		PackHeader()
		totalBytes = append(headerBytes, bodyBytes...)
		bodyBytes = nil
		headerBytes = nil
		dataBuffer.Reset()
		return totalBytes
	}
	return nil
}

func loginWrap() {
	packString(login.UserID)
	packString(login.Password)
	dataLength = dataBuffer.Len() + headerLength
	bodyBytes = append(bodyBytes, dataBuffer.Bytes()...)
	dataBuffer.Reset()
}

func chatWrap() {
	packString(chat.msg)
	dataLength = dataBuffer.Len() + headerLength
	bodyBytes = append(bodyBytes, dataBuffer.Bytes()...)
	dataBuffer.Reset()
}

func packInt(data int) {
	buff := make([]byte, 4)
	binary.BigEndian.PutUint32(buff, uint32(data))
	dataBuffer.Write(buff)
}

func packShort(data interface{}) {
	switch data.(type) {
	case int:
		buff := make([]byte, 2)
		binary.BigEndian.PutUint16(buff, uint16(data.(int)))
		dataBuffer.Write(buff)
	case int16:
		buff := make([]byte, 2)
		binary.BigEndian.PutUint16(buff, uint16(data.(int16)))
		dataBuffer.Write(buff)
	}
}

func packString(data string) {
	strLength := len(data)
	packShort(strLength)
	dataBuffer.Write([]byte(data))
}
