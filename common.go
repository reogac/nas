package nas

import (
	"encoding/binary"
)

type Bytes []byte

func NewBytes(bytes []byte) *Bytes {
	v := Bytes(make([]byte, len(bytes)))
	copy(v[:], bytes)
	return &v
}

func (ie *Bytes) decode(wire []byte) error {
	*ie = make([]byte, len(wire))
	copy(*ie, wire)
	return nil
}

func (ie *Bytes) encode() ([]byte, error) {
	return []byte(*ie), nil
}

type Uint8 uint8

func NewUint8(v uint8) *Uint8 {
	tmp := Uint8(v)
	return &tmp
}

func (ie *Uint8) Value() uint8 {
	return uint8(*ie)
}

func (ie *Uint8) encode() ([]byte, error) {
	return []byte{uint8(*ie)}, nil
}

func (ie *Uint8) decode(wire []byte) error {
	if len(wire) != 1 {
		return ErrInvalidSize
	}
	*ie = Uint8(wire[0])
	return nil
}

type Uint16 uint16

func NewUint16(v uint16) *Uint16 {
	tmp := Uint16(v)
	return &tmp
}

func (ie *Uint16) encode() ([]byte, error) {
	var wire [2]byte
	binary.BigEndian.PutUint16(wire[:], uint16(*ie))
	return wire[:], nil
}

func (ie *Uint16) decode(wire []byte) error {
	if len(wire) != 2 {
		return ErrInvalidSize
	}
	*ie = Uint16(binary.BigEndian.Uint16(wire))
	return nil
}
