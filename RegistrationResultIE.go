package nas

// 9.11.3.6 5GS registration result [M LV 2]
type RegistrationResult struct {
	value uint8
}

func (ie *RegistrationResult) encode() (wire []byte, err error) {
	wire = []byte{ie.value}
	return
}

func (ie *RegistrationResult) decode(wire []byte) (err error) {
	if len(wire) != 1 {
		err = ErrInvalidSize
		return
	}
	ie.value = wire[0]
	return
}

func (ie *RegistrationResult) GetSmsAllowed() bool {
	return getBit(ie.value, 3) == 1 //4th bit from right
}

func (ie *RegistrationResult) SetSmsAllowed(flag bool) {
	if flag {
		ie.value = setBit(ie.value, 3)
	} else {
		ie.value = clearBit(ie.value, 3)
	}
}

func (ie *RegistrationResult) GetResult() uint8 {
	return ie.value & 0x07 //last 3 bits
}

func (ie *RegistrationResult) SetResult(v uint8) {
	mask := uint8(0x07)
	ie.value = (ie.value & ^mask) | (v & mask)
}
