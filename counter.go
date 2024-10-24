package nas

import (
	"encoding/binary"
)

const (
	NAS_COUNT_WINDOW uint8 = 20
)

/*
TS 33.501 6.4.3.1

	COUNT (32 bits) := 0x00 || NAS COUNT (24 bits)
	NAS COUNT (24 bits) := NAS OVERFLOW (16 bits) || NAS SQN (8 bits)
*/

type Counter uint32

func newCounter(overflow uint16, sqn uint8) (c Counter) {
	c.set(overflow, sqn)
	return
}

func (c *Counter) bytes() (b []byte) {
	b = make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(*c))
	return
}

func (c *Counter) set(overflow uint16, sqn uint8) {
	v := uint32(sqn)           //initialize with sqn
	v |= uint32(overflow) << 8 //fill in overflow
	*c = Counter(v)
}

func (c *Counter) setSqn(sqn uint8) {
	*c = Counter((uint32(*c) & 0xffffff00) | uint32(sqn))
}

func (c *Counter) inc() {
	v := uint32(*c) + 1
	*c = Counter(v & 0x00ffffff) //mask last 24 bits
}

func (c *Counter) sqn() uint8 {
	return uint8(*c & 0x000000ff)
}

func (c *Counter) overflow() uint16 {
	return uint16((*c & 0x00ffff00) >> 8)
}

func (c *Counter) setOverflow(overflow uint16) {
	*c = Counter((uint32(*c) & 0xff0000ff) | (uint32(overflow) << 8))
}
