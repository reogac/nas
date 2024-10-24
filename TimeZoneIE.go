package nas

// 9.11.3.52 Time zone [O TV 2]
type TimeZone struct {
	Value uint8
}

func (ie *TimeZone) encode() (wire []byte, err error) {
	wire = []byte{ie.Value}
	return
}

func (ie *TimeZone) decode(wire []byte) (err error) {
	if len(wire) != 1 {
		err = ErrInvalidSize
		return
	}
	ie.Value = wire[0]
	return
}
