package nas

// 9.11.3.5 5GS network feature support [O TLV 3-5]
type NetworkFeatureSupport struct {
	bytes [2]byte
}

func (ie *NetworkFeatureSupport) encode() (wire []byte, err error) {
	wire = ie.bytes[:]
	return
}

func (ie *NetworkFeatureSupport) decode(wire []byte) (err error) {
	if len(wire) < 2 {
		return ErrIncomplete
	} else if len(wire) > 2 {
		return ErrTail
	}

	copy(ie.bytes[:], wire)
	return
}

// octet 0, bit 7
func (ie *NetworkFeatureSupport) GetMPSI() bool {
	return getBit(ie.bytes[0], 7) == 1
}
func (ie *NetworkFeatureSupport) SetMPSI(f bool) {
	if f {
		ie.bytes[0] = setBit(ie.bytes[0], 7)
	} else {
		ie.bytes[0] = clearBit(ie.bytes[0], 7)
	}
}

// octet 0, bit 6
func (ie *NetworkFeatureSupport) GetIWKN26() bool {
	return getBit(ie.bytes[0], 6) == 1
}
func (ie *NetworkFeatureSupport) SetIWKN26(f bool) {
	if f {
		ie.bytes[0] = setBit(ie.bytes[0], 6)
	} else {
		ie.bytes[0] = clearBit(ie.bytes[0], 6)
	}
}

// octet 0, bit 5,4
func (ie *NetworkFeatureSupport) GetEMF() uint8 {
	return (ie.bytes[0] & bitMask(4, 2)) >> 4
}
func (ie *NetworkFeatureSupport) SetEMF(v uint8) {
	mask := uint8(0x03) //2 bits
	ie.bytes[0] = (ie.bytes[0] & ^(mask << 4)) | ((v & mask) << 4)
}

// octet 0, bit 3,2
func (ie *NetworkFeatureSupport) GetEMC() uint8 {
	return (ie.bytes[0] & bitMask(2, 2)) >> 2
}

func (ie *NetworkFeatureSupport) SetEMC(v uint8) {
	mask := uint8(0x03)
	ie.bytes[0] = (ie.bytes[0] & ^(mask << 2)) | ((v & mask) << 2)
}

// octet 0, bit 1
func (ie *NetworkFeatureSupport) GetIMSVoPSN3GPP() bool {
	return getBit(ie.bytes[0], 1) == 1
}
func (ie *NetworkFeatureSupport) SetIMSVoPSN3GPP(f bool) {
	if f {
		ie.bytes[0] = setBit(ie.bytes[0], 1)
	} else {
		ie.bytes[0] = clearBit(ie.bytes[0], 1)
	}
}

// octet 0, bit 0
func (ie *NetworkFeatureSupport) GetIMSVoPS3GPP() bool {
	return getBit(ie.bytes[0], 0) == 0
}
func (ie *NetworkFeatureSupport) SetIMSVoPS3GPP(f bool) {
	if f {
		ie.bytes[0] = setBit(ie.bytes[0], 0)
	} else {
		ie.bytes[0] = clearBit(ie.bytes[0], 0)
	}
}

// octet 1, bit 1
func (ie *NetworkFeatureSupport) GetMCSI() bool {
	return getBit(ie.bytes[1], 1) == 1
}

func (ie *NetworkFeatureSupport) SetMCSI(f bool) {
	if f {
		ie.bytes[1] = setBit(ie.bytes[1], 1)
	} else {
		ie.bytes[1] = clearBit(ie.bytes[1], 1)
	}
}

// octet 1, bit 0
func (ie *NetworkFeatureSupport) GetEMCN() bool {
	return getBit(ie.bytes[1], 0) == 0
}

func (ie *NetworkFeatureSupport) SetEMCN(f bool) {
	if f {
		ie.bytes[1] = setBit(ie.bytes[1], 0)
	} else {
		ie.bytes[1] = clearBit(ie.bytes[1], 0)
	}
}
