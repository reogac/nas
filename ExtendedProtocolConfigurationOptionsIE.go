package nas

import (
	"encoding/binary"
	"math"
)

// 9.11.4.6 Extended protocol configuration options [O TLV-E 4-65538]
type ExtendedProtocolConfigurationOptions struct {
	units []PcoUnit
}

func (ie *ExtendedProtocolConfigurationOptions) Units() []PcoUnit {
	return ie.units
}

func (ie *ExtendedProtocolConfigurationOptions) AddUnit(unit PcoUnit) {
	ie.units = append(ie.units, unit)
}

func (ie *ExtendedProtocolConfigurationOptions) encode() (wire []byte, err error) {
	var extension uint8 = 1
	var spare uint8 = 0
	var configurationProtocol uint8 = 0
	var buf []byte

	wire = []byte{(extension << 7) | (spare << 6) | (configurationProtocol)}
	for _, u := range ie.units {
		if buf, err = u.toBytes(); err != nil {
			return
		}
		wire = append(wire, buf...)
	}

	return
}

func (ie *ExtendedProtocolConfigurationOptions) decode(wire []byte) (err error) {
	if len(wire) < 1 {
		err = ErrIncomplete
		return
	}
	var unit PcoUnit
	var unitLength int
	offset := 1
	wireLen := len(wire)
	for offset < wireLen {
		if offset+3 > wireLen { //make sure we have Id and len of an unit
			err = ErrIncomplete
			return
		}
		unit.Id = binary.BigEndian.Uint16(wire[offset : offset+2])
		unitLength = int(wire[offset+2])
		offset += 3
		if offset+unitLength > wireLen {
			err = ErrIncomplete
			return
		}
		unit.Content = make([]byte, unitLength)
		copy(unit.Content, wire[offset:offset+unitLength])
		ie.units = append(ie.units, unit)
	}
	return
}

type PcoUnit struct {
	Id      uint16
	Content []byte
}

func (u *PcoUnit) toBytes() (buf []byte, err error) {
	if len(u.Content) > math.MaxUint8 {
		err = ErrInvalidSize
		return
	}
	buf = []byte{0, 0, uint8(len(u.Content))}
	binary.BigEndian.PutUint16(buf, u.Id)
	buf = append(buf, u.Content...)
	return
}
