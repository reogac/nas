package nas

// 9.11.4.11 PDU session type [O TV 1]
type PduSessionType struct {
	Value uint8
}

func (ie *PduSessionType) encode() (wire []byte, err error) {
	wire = []byte{ie.Value}
	return
}

func (ie *PduSessionType) decode(wire []byte) (err error) {
	if len(wire) != 1 {
		err = ErrInvalidSize
		return
	}
	ie.Value = wire[0]
	return
}
