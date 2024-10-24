package nas

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

// 9.11.3.8 5GS tracking area identity [O TV 7]
type TrackingAreaIdentity struct {
	PlmnId PlmnId
	Tac    uint32
}

func (ie *TrackingAreaIdentity) Set(mcc, mnc, tac string) (err error) {
	if err = ie.PlmnId.Set(mcc, mnc); err != nil {
		return
	}
	var tacBytes []byte
	if tacBytes, err = hex.DecodeString(tac); err != nil {
		return
	}
	if len(tacBytes) != 3 {
		return fmt.Errorf("TAC msut be 3 bytes")
	}
	ie.Tac = tacFromBytes(tacBytes)
	return
}

func (ie *TrackingAreaIdentity) encode() (wire []byte, err error) {
	wire = append(ie.PlmnId.bytes[:], tacBytes(ie.Tac)...)
	return
}

func (ie *TrackingAreaIdentity) decode(wire []byte) (err error) {
	if len(wire) < 6 {
		return ErrIncomplete
	} else if len(wire) > 6 {
		return ErrTail
	}
	ie.readBytes(wire)
	return
}

// wire must be 6 bytes
func (tai *TrackingAreaIdentity) readBytes(wire []byte) {
	copy(tai.PlmnId.bytes[:], wire[0:3])
	tai.Tac = tacFromBytes(wire[3:6])
}

func tacBytes(tac uint32) []byte {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], tac)
	return buf[1:]
}

// buf must be 3 bytes
func tacFromBytes(wire []byte) uint32 {
	var buf [4]byte
	copy(buf[1:], wire[0:3])
	return binary.BigEndian.Uint32(buf[:])
}
