package nas

// 9.11.3.7 5GS registration type [M V 1/2]
type RegistrationType struct {
	Value uint8
}

func (ie *RegistrationType) encode() (wire []byte, err error) {
	wire = []byte{ie.Value & 0x07}
	return
}

func (ie *RegistrationType) decode(wire []byte) (err error) {
	if len(wire) != 1 {
		err = ErrInvalidSize
		return
	}
	ie.Value = wire[0] & 0x07
	return
}
