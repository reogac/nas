package nas

// 9.11.4.13 QoS rules [M LV-E 6-65538]
type QosRules struct {
	Bytes []byte
}

func (ie *QosRules) encode() (wire []byte, err error) {
	wire = ie.Bytes
	return
}

func (ie *QosRules) decode(wire []byte) (err error) {
	ie.Bytes = make([]byte, len(wire))
	copy(ie.Bytes, wire)
	return
}
