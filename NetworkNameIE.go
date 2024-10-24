package nas

// 9.11.3.35 Network name [O TLV 3-n]
type NetworkName struct {
	Bytes []byte
}

func (ie *NetworkName) encode() (wire []byte, err error) {
	wire = ie.Bytes
	return
}

func (ie *NetworkName) decode(wire []byte) (err error) {
	ie.Bytes = make([]byte, len(wire))
	copy(ie.Bytes, wire)
	return
}
