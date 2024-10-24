package nas

// 9.11.3.30 LADN information [O TLV-E 12-1715]
type LadnInformation struct {
	Bytes []byte
}

func (ie *LadnInformation) encode() (wire []byte, err error) {
	wire = ie.Bytes
	return
}

func (ie *LadnInformation) decode(wire []byte) (err error) {
	ie.Bytes = make([]byte, len(wire))
	copy(ie.Bytes, wire)
	return
}
