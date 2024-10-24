package nas

// extended protocol description
const (
	EPD_5GSM uint8 = 0x2E
	EPD_5GMM uint8 = 0x7E
)

type MmHeader struct {
	sec     uint8
	msgType uint8
}

func (h *MmHeader) SetSecurityHeader(v uint8) {
	h.sec = v
}

func (h *MmHeader) GetSecurityHeader() uint8 {
	return h.sec
}

// wire must be 2 bytes (EPD is removed)
func mmHeaderFromBytes(wire []byte) *MmHeader {
	return &MmHeader{
		sec:     wire[0],
		msgType: wire[1],
	}
}

// three bytes including EPD
func (h *MmHeader) headerBytes() []byte {
	return []byte{EPD_5GMM, h.sec, h.msgType}
}

type SmHeader struct {
	sessionId uint8
	pti       uint8
	msgType   uint8
}

// wire must be 3 bytes, EPD is removed
func smHeaderFromBytes(wire []byte) *SmHeader {
	return &SmHeader{
		sessionId: wire[0],
		pti:       wire[1],
		msgType:   wire[2],
	}
}

// 4bytes including EPD
func (h *SmHeader) headerBytes() []byte {
	return []byte{EPD_5GSM, h.sessionId, h.pti, h.msgType}
}

func (h *SmHeader) GetPti() uint8 {
	return h.pti
}

func (h *SmHeader) GetSessionId() uint8 {
	return h.sessionId
}

func (h *SmHeader) SetPti(v uint8) {
	h.pti = v
}

func (h *SmHeader) SetSessionId(v uint8) {
	h.sessionId = v
}
