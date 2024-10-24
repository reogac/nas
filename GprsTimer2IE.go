package nas

// 9.11.2.4 GPRS timer 2 [O TLV 3]
type GprsTimer2 struct {
	Value uint8
}

func (ie *GprsTimer2) encode() (wire []byte, err error) {
	wire = []byte{ie.Value}
	return
}

func (ie *GprsTimer2) decode(wire []byte) (err error) {
	if len(wire) != 1 {
		err = ErrInvalidSize
		return
	}
	ie.Value = wire[0]
	return
}
