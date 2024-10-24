package nas

// 9.11.3.1 5GMM capability [O TLV 3-15]
type GmmCapability struct {
	Bytes []byte
}

func (ie *GmmCapability) encode() (wire []byte, err error) {
	wire = ie.Bytes
	return
}

func (ie *GmmCapability) decode(wire []byte) (err error) {
	ie.Bytes = make([]byte, len(wire))
	copy(ie.Bytes, wire)
	return
}
