package nas

// 9.11.3.44 PDU session status [O TLV 4-34]
type PduSessionStatus struct {
	bytes  [2]byte
	Spares []byte
}

func (ie *PduSessionStatus) encode() (wire []byte, err error) {
	wire = ie.bytes[:]
	if len(ie.Spares) > 0 {
		wire = append(wire, ie.Spares...)
	}
	return
}

func (ie *PduSessionStatus) decode(wire []byte) (err error) {
	if len(wire) < 2 {
		return ErrIncomplete
	}
	copy(ie.bytes[:], wire[0:3])
	if len(wire) > 2 {
		ie.Spares = make([]byte, len(wire)-2)
		copy(ie.Spares[:], wire[2:])
	}
	return
}

func (ie *PduSessionStatus) Set(flags [16]bool) {
	var pos uint8
	var row int
	for i, f := range flags {
		row = i / 8
		pos = uint8(i % 8)
		if f {
			ie.bytes[row] = setBit(ie.bytes[row], pos)
		} else {
			ie.bytes[row] = clearBit(ie.bytes[row], pos)
		}
	}
}

func (ie *PduSessionStatus) Get() (flags [16]bool) {
	var pos uint8
	var row int
	for i := 0; i < 16; i++ {
		row = i / 8
		pos = uint8(i % 8)
		flags[i] = getBit(ie.bytes[row], pos) == 1
	}
	return
}
