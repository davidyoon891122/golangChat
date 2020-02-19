package tools

type Header struct {
	Length     int
	DataLength int
	Process    int16
	Service    int16
}

func InitHeader() *Header {
	header := &Header{}

	return header
}

func (h *Header) headerPacker(length int, dataLength int, process int, service int) {
	h.putLength(length)
	h.putDataLength(dataLength)
	h.putProcess(process)
	h.putService(service)
}

func (h *Header) putLength(length int) {
	h.Length = length
}

func (h *Header) putDataLength(dataLength int) {
	h.DataLength = dataLength
}

func (h *Header) putProcess(process int) {
	h.Process = int16(process)
}

func (h *Header) putService(service int) {
	h.Service = int16(service)
}
