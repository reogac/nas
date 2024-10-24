package nas

// 9.11.4.10 PDU address [O TLV 11]
type PduAddress struct {
	bytes []byte
}

func NewPduAddress(t uint8, content []byte) *PduAddress {
	v := &PduAddress{
		bytes: make([]byte, len(content)+1),
	}
	v.bytes[0] = t
	copy(v.bytes[1:], content)
	return v
}

func (ie *PduAddress) AddressType() uint8 {
	return ie.bytes[0] & 0x0f
}

func (ie *PduAddress) Content() []byte {
	return ie.bytes[1:]
}

func (ie *PduAddress) encode() (wire []byte, err error) {
	wire = make([]byte, len(ie.bytes))
	copy(wire, ie.bytes)
	return
}

func (ie *PduAddress) decode(wire []byte) (err error) {
	if len(wire) < 1 {
		err = ErrInvalidSize
		return
	}
	ie.bytes = make([]byte, len(wire))
	copy(ie.bytes, wire)
	return
}
