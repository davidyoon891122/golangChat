package tools

import (
	"bytes"
	"encoding/binary"
)

var headerLength int = 12
var header *Header
var (
	dataBuffer bytes.Buffer
)

func PackHeader() {
	packInt(header.Length)
	packInt(header.DataLength)
	packShort(header.Process)
	packShort(header.Service)

}

func Pack(data string) []byte {
	dataBuffer.Reset()
	dataLength := len(data) + headerLength
	header = InitHeader()
	header.headerPacker(headerLength, dataLength, 0, 0)
	PackHeader()
	bodyPacker(data)

	return dataBuffer.Bytes()
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
