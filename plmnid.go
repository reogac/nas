package nas

import (
	"encoding/hex"
	"fmt"
)

type PlmnId struct {
	bytes [3]byte
}

func (id *PlmnId) decode(wire []byte) (err error) {
	if len(wire) != 3 {
		err = ErrInvalidSize
		return
	}
	copy(id.bytes[:], wire)
	return
}

func (id *PlmnId) encode() (wire []byte, err error) {
	wire = id.bytes[:]
	return
}

func (id *PlmnId) Parse(s string) (err error) {
	if len(s) != 5 && len(s) != 6 {
		return fmt.Errorf("PlmnId must be 5 or 6 digits")
	}

	return id.Set(s[0:3], s[3:])
}

func (id *PlmnId) String() string {
	mcc, mnc := id.Get()
	return fmt.Sprintf("%s%s", mcc, mnc)
}

func (id *PlmnId) Set(mcc, mnc string) (err error) {
	//encode from string here
	var plmnId string
	if len(mcc) != 3 {
		err = fmt.Errorf("Mcc length must be 3")
		return
	}

	if len(mnc) != 3 && len(mnc) != 2 {
		err = fmt.Errorf("Mnc length must be 2 or 3")
		return
	}
	plmnId = mcc + mnc
	for _, c := range plmnId {
		if c-'0' >= 0 && c-'0' <= 9 {
			continue
		}
		if c-'a' >= 0 && c-'a' <= 5 {
			continue
		}

		err = fmt.Errorf("\"%c\" is not a hex character", c)
		return
	}

	if len(mnc) == 2 {
		plmnId = plmnId + "f"
	}

	tmp, _ := hex.DecodeString(plmnId) //must be 3 byte
	id.bytes[0] = (tmp[0] << 4) | (tmp[0] >> 4)
	id.bytes[1] = ((tmp[2] & 0x0f) << 4) | ((tmp[1] & 0xf0) >> 4)
	id.bytes[2] = (tmp[2] & 0xf0) | (tmp[1] & 0x0f)
	return
}

func (id *PlmnId) Get() (mcc, mnc string) {
	//decode from string here
	mccDigit1 := id.bytes[0] & 0x0f
	mccDigit2 := (id.bytes[0] & 0xf0) >> 4
	mccDigit3 := (id.bytes[1] & 0x0f)

	mncDigit1 := (id.bytes[2] & 0x0f)
	mncDigit2 := (id.bytes[2] & 0xf0) >> 4
	mncDigit3 := (id.bytes[1] & 0xf0) >> 4

	plmnIdBytes := []byte{(mccDigit1 << 4) | mccDigit2, (mccDigit3 << 4) | mncDigit1, (mncDigit2 << 4) | mncDigit3}
	plmnId := hex.EncodeToString(plmnIdBytes)
	mcc = plmnId[0:3]
	if plmnId[5] == 'f' {
		mnc = plmnId[3:5]
	} else {
		mnc = plmnId[3:6]
	}

	return
}
