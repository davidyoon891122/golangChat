package tools

import (
	"encoding/binary"
)

var index int = 0
var recvBytes []byte

func Unpack(recvBytesG []byte) string {
	recvBytes = recvBytesG
	readHeader()
	ret := ReadChat()
	index = 0
	return ret
}

func readHeader() {
	header = InitHeader()
	header.Length = readInt()
	header.DataLength = readInt()
	header.Process = readShort()
	header.Service = readShort()
}

func readBody() {
	ReadChat()

}

func (h *Header) GetLength() int {
	return h.Length
}

func (h *Header) GetDataLength() int {
	return h.DataLength
}

func (h *Header) GetProcess() int16 {
	return h.Process
}

func (h *Header) GetService() int16 {
	return h.Service
}

func ReadChat() string {
	chat := readString()

	return chat
}

func GetChat() {

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
