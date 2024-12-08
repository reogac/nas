package nas

// 9.11.3.34 security algorithms [M V 1]
type SecurityAlgorithms struct {
	value uint8
}

func NewSecurityAlgorithms(intAlg, encAlg uint8) SecurityAlgorithms {
	v := intAlg & 0x0f
	v += (encAlg & 0x0f) << 4
	return SecurityAlgorithms{
		value: v,
	}
}

func (ie *SecurityAlgorithms) encode() (wire []byte, err error) {
	wire = []byte{ie.value}
	return
}

func (ie *SecurityAlgorithms) decode(wire []byte) (err error) {
	if len(wire) != 1 {
		err = ErrInvalidSize
		return
	}
	ie.value = wire[0]
	return
}

func (ie *SecurityAlgorithms) IntAlg() uint8 {
	return ie.value & 0x0f
}

func (ie *SecurityAlgorithms) EncAlg() uint8 {
	return ie.value & 0xf0 >> 4
}
