package nas

// 9.11.2.5 GPRS timer 3 [O TLV 3]
type GprsTimer3 struct {
	Value uint8
}

func (ie *GprsTimer3) encode() (wire []byte, err error) {
	wire = []byte{ie.Value}
	return
}

func (ie *GprsTimer3) decode(wire []byte) (err error) {
	if len(wire) != 1 {
		err = ErrInvalidSize
		return
	}
	ie.Value = wire[0]
	return
}

// just for testing, need to correct this
func NewGprsTimer3(u uint8, v uint8) *GprsTimer3 {
	//TODO:
	return &GprsTimer3{
		Value: v,
	}
}
