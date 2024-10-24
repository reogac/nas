package nas

import (
	"fmt"
)

// TS 24.501 9.11.3.32
const (
	TscNative uint8 = 0x00
	TscMapped uint8 = 0x01
)

// TS 24.501 9.11.3.32
const (
	NasKeySetIdentifierNoKeyIsAvailable uint8 = 0x07
)

// 9.11.3.32 key set identifier [M V 1/2]
type KeySetIdentifier struct {
	Tsc uint8
	Id  uint8
}

func (ie *KeySetIdentifier) encode() (wire []byte, err error) {
	v := ie.Id & 0x07 // get last 3 bits
	switch ie.Tsc {
	case TscNative:
	case TscMapped:
		v |= 0x08 //set bit 4
	default:
		err = fmt.Errorf("Invalid Tsc")
	}
	wire = []byte{v}
	return
}

func (ie *KeySetIdentifier) decode(wire []byte) (err error) {
	if len(wire) < 1 {
		err = fmt.Errorf("Ksi empty")
		return
	}
	ie.Id = wire[0] & 0x07       //get last 3 bits
	ie.Tsc = wire[0] & 0x08 >> 4 //get bit 4
	return
}

func (ie *KeySetIdentifier) String() string {
	if ie.Tsc == TscNative {
		return fmt.Sprintf("native:%d", ie.Id)
	} else {
		return fmt.Sprintf("mapped:%d", ie.Id)
	}
}
