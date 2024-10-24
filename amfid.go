package nas

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type AmfId struct {
	bytes [3]byte
}

// decimal format region-set-pointer; not empty
func (id *AmfId) ParseDecimals(s string) (err error) {
	//reset
	id.bytes[0] = 0
	id.bytes[1] = 0
	id.bytes[2] = 0

	parts := strings.Split(s, "-")
	if len(parts) > 3 {
		return fmt.Errorf("Invalid AmfId")
	}
	var v uint64
	if len(parts) >= 1 {
		if v, err = strconv.ParseUint(parts[0], 10, 8); err != nil {
			err = fmt.Errorf("Fail to parse region: %+v", err)
			return
		}
		id.SetRegion(uint8(v))
	}
	if len(parts) >= 2 {
		if v, err = strconv.ParseUint(parts[1], 10, 8); err != nil {
			err = fmt.Errorf("Fail to parse set: %+v", err)
			return
		}
		id.SetSet(uint16(v))
	}
	if len(parts) == 3 {
		if v, err = strconv.ParseUint(parts[2], 10, 8); err != nil {
			err = fmt.Errorf("Fail to parse pointer: %+v", err)
			return
		}
		id.SetPointer(uint8(v))
	}
	return
}

// hex string (6bytes)
func (id *AmfId) Parse(s string) (err error) {
	if buf, err := hex.DecodeString(s); err != nil {
		return fmt.Errorf("AmfId not a valid hex string: %+v", err)
	} else if len(buf) != 3 {
		return fmt.Errorf("AmfId must be 3 bytes")
	} else {
		copy(id.bytes[:], buf)
	}
	return nil
}

// hex string
func (id *AmfId) String() string {
	return hex.EncodeToString(id.bytes[:])
}

// decimal formated string with separators
func (id *AmfId) DecimalString() string {
	region, set, pointer := id.Get()
	return fmt.Sprintf("%d-%d-%d", region, set, pointer)
}

func (id *AmfId) Get() (amfRegion uint8, amfSet uint16, amfPointer uint8) {
	amfRegion = id.GetRegion()
	amfSet = id.GetSet()
	amfPointer = id.GetPointer()
	return
}

func (id *AmfId) Set(amfRegion uint8, amfSet uint16, amfPointer uint8) {
	id.SetRegion(amfRegion)
	id.SetSet(amfSet)
	id.SetPointer(amfPointer)
}

func (id *AmfId) GetRegion() uint8 {
	return id.bytes[0]
}

func (id *AmfId) SetRegion(v uint8) {
	id.bytes[0] = v
}

func (id *AmfId) GetSet() uint16 {
	return (uint16(id.bytes[1]) << 2) + (uint16(id.bytes[2]) >> 6)
}

func (id *AmfId) SetSet(v uint16) {
	id.bytes[1] = uint8((v & 0x3fff) >> 2) //first 8bits
	id.bytes[2] &= 0x3f                    //zero first 2bits
	id.bytes[2] += (uint8(v&0x03) << 6)    //then fill with last 2 bits of v
}

func (id *AmfId) GetPointer() uint8 {
	return id.bytes[2] & 0x3f //last 6bits
}

func (id *AmfId) SetPointer(v uint8) {
	id.bytes[2] = (id.bytes[2] & 0xc0) + (v & 0x3f)
}

// create AmfId from a hex string
func ParseAmfId(s string) (*AmfId, error) {
	id := new(AmfId)
	if err := id.Parse(s); err != nil {
		return nil, err
	}
	return id, nil
}
