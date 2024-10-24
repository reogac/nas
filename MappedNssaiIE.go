package nas

// 9.11.3.31B Mapped NSSAI [O TLV 3-42]
type MappedNssai struct {
	Bytes []byte
}

func (ie *MappedNssai) encode() (wire []byte, err error) {
	wire = ie.Bytes
	return
}

func (ie *MappedNssai) decode(wire []byte) (err error) {
	ie.Bytes = make([]byte, len(wire))
	copy(ie.Bytes, wire)
	return
}
