package nas

import (
	"encoding/binary"
	"fmt"
)

func twoBytes(v uint16) []byte {
	var buf [2]byte
	binary.BigEndian.PutUint16(buf[:], v)
	return buf[:]
}

func fromTwoBytes(buf []byte) (v uint16, err error) {
	if len(buf) == 2 {
		v = binary.BigEndian.Uint16(buf)
	} else {
		err = ErrInvalidSize
	}
	return
}

func getIei(v uint8) uint8 {
	if v >= 0x80 {
		return v & 0xf0 >> 4
	}
	return v
}

func checkLengthBounds(v, l, u uint16) bool {
	if v < l { //check lower bound
		log.Errorf("fail check %d<%d\n", v, l)
		return false
	}
	//NOTE: must check a corner case where l=u=0
	//u = 0 means no upper bound
	if u != 0 && u >= l { //check if upper bound is valid
		if v > u {
			log.Errorf("fail check %d>%d\n", v, u)
			return false
		}
	}
	return true
}

// decode LV, LV-E
func decodeLV(buf []byte, extend bool, l, u uint16, v Decoder) (consumed int, err error) {
	bufSize := len(buf)
	lenSize := 1 //TLV
	if extend {
		lenSize = 2 //TLV-E
	}
	if lenSize > bufSize {
		err = ErrIncomplete
		return
	}
	lenValue := uint16(buf[0])
	if extend {
		lenValue = binary.BigEndian.Uint16(buf[0:2])
	}

	log.Tracef("DecodeLV %d byte\n", lenValue)

	if !checkLengthBounds(lenValue, l, u) {
		err = ErrInvalidSize
		return
	}

	consumed = int(uint16(lenSize) + lenValue)
	if consumed > bufSize {
		consumed = bufSize
		err = ErrIncomplete
		return
	}
	err = v.decode(buf[lenSize:consumed])
	return
}

// encode LV, LV-E
func encodeLV(extend bool, l, u uint16, v Encoder) (wire []byte, err error) {

	var buf []byte
	if buf, err = v.encode(); err != nil {
		return
	}
	vLen := uint16(len(buf))

	if !checkLengthBounds(vLen, l, u) {
		err = ErrInvalidSize
		return
	}

	log.Tracef("EncodeLV %d byte\n", vLen)
	if extend {
		prefix := [2]byte{0, 0}
		binary.BigEndian.PutUint16(prefix[:], vLen)
		wire = append(prefix[:], buf...)
	} else {
		wire = append([]byte{uint8(vLen)}, buf...)
	}
	return
}

// lastBit 0[rightmost][leftmost]:7; numBits:0:8
func bitMask(lastBit, numBits uint8) uint8 {
	return (uint8(1)<<numBits - uint8(1)) << lastBit
}

// pos must be 0:7
func getBit(v uint8, pos uint8) uint8 {
	return (v & (uint8(1) << pos)) >> pos
}

// pos must be 0:7
func setBit(v uint8, pos uint8) uint8 {
	return (v | (uint8(1) << pos))
}

// pos must be 0:7
func clearBit(v uint8, pos uint8) uint8 {
	mask := ^(uint8(1) << pos)
	return v & mask
}

// decimal string to byte array
func decimalBytes(s string) (buf []byte, err error) {
	tmp := make([]byte, len(s))
	for i, c := range s {
		d := uint8(c - '0')
		if d >= 0 && d <= 9 {
			tmp[i] = d
		} else {
			err = fmt.Errorf("\"%c\" is not decimal number", c)
			return
		}
	}
	buf = tmp
	return
}
