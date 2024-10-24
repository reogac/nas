package nas

// 9.11.4.12 QoS flow descriptions [O TLV-E 6-65538]
type QosFlowDescriptions struct {
	Bytes []byte
}

func (ie *QosFlowDescriptions) encode() (wire []byte, err error) {
	wire = ie.Bytes
	return
}

func (ie *QosFlowDescriptions) decode(wire []byte) (err error) {
	ie.Bytes = make([]byte, len(wire))
	copy(ie.Bytes, wire)
	return
}
