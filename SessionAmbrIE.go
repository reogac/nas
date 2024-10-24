package nas

import (
	"encoding/binary"
)

// 9.11.4.14 Session-AMBR [M LV 7]
type SessionAmbr struct {
	Bytes [6]byte
}

func NewSessionAmbr(ulU uint8, ulR uint16, dlU uint8, dlR uint16) *SessionAmbr {
	ambr := &SessionAmbr{}
	ambr.Bytes[0] = ulU
	binary.BigEndian.PutUint16(ambr.Bytes[1:], ulR)
	ambr.Bytes[3] = dlU
	binary.BigEndian.PutUint16(ambr.Bytes[4:], dlR)
	return ambr
}

func (ie *SessionAmbr) encode() (wire []byte, err error) {
	wire = ie.Bytes[:]
	return
}

func (ie *SessionAmbr) decode(wire []byte) (err error) {
	if len(wire) != 6 {
		err = ErrInvalidSize
		return
	}
	copy(ie.Bytes[:], wire)
	return
}
