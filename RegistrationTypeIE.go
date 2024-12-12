package nas

// 9.11.3.7 5GS registration type [M V 1/2]
type RegistrationType struct {
	value uint8
}

func NewRegistrationType(forValue bool, typeValue uint8) (v RegistrationType) {
	if forValue {
		v.value = 1<<3 + typeValue&0x07
	} else {
		v.value = typeValue & 0x07
	}
	return
}

func (ie *RegistrationType) encode() (wire []byte, err error) {
	wire = []byte{ie.value & 0x0f}
	return
}

func (ie *RegistrationType) decode(wire []byte) (err error) {
	if len(wire) != 1 {
		err = ErrInvalidSize
		return
	}
	ie.value = wire[0] & 0x0f
	return
}

func (ie *RegistrationType) GetFor() bool {
	return (ie.value&0x0f)>>3 == 1
}

func (ie *RegistrationType) GetType() uint8 {
	return ie.value & 0x07
}
