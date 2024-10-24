package nas

// 9.11.3.12 Additional 5G security information [O TLV 3]
type AdditionalSecurityInformation struct {
	value uint8
}

func NewAdditionalSecurityInformation(retransmission, hoDerivation bool) *AdditionalSecurityInformation {
	ie := &AdditionalSecurityInformation{}
	ie.SetRetransmission(retransmission)
	ie.SetHoDerivation(hoDerivation)
	return ie
}

func (ie *AdditionalSecurityInformation) encode() (wire []byte, err error) {
	wire = []byte{ie.value}
	return
}

func (ie *AdditionalSecurityInformation) decode(wire []byte) (err error) {
	if len(wire) != 1 {
		err = ErrInvalidSize
		return
	}
	ie.value = wire[0]
	return
}

func (ie *AdditionalSecurityInformation) SetRetransmission(flag bool) {
	if flag {
		setBit(ie.value, 1)
	} else {
		clearBit(ie.value, 1)
	}
}
func (ie *AdditionalSecurityInformation) SetHoDerivation(flag bool) {
	if flag {
		setBit(ie.value, 0)
	} else {
		clearBit(ie.value, 0)
	}

}
