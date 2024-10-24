package nas

import (
	"encoding/hex"
	"fmt"
)

type MappedSNssai struct {
	Sst uint8
	Sd  []byte //0 or 3 bytes
}

// 9.11.2.8 S-NSSAI [O TLV 3-10]
type SNssai struct {
	Sst    uint8
	Sd     []byte //0 or 3 bytes
	Mapped *MappedSNssai
}

func (ie *SNssai) Set(sst uint8, sd string) error {
	ie.Sst = sst //always set Sst so error can be ignored
	if buf, err := hex.DecodeString(sd); err != nil {
		return fmt.Errorf("decode Sd from hex string fails: %+v", err)
	} else if len(buf) != 3 {
		return fmt.Errorf("Sd must be 3 bytes")
	} else {
		ie.Sd = buf
	}
	return nil
}

func (ie *SNssai) SetMapped(sst uint8, sd string) error {
	ie.Mapped = &MappedSNssai{
		Sst: sst, //always set Sst so error can be ignored
	}
	if buf, err := hex.DecodeString(sd); err != nil {
		return fmt.Errorf("decode Sd from hex string fails: %+v", err)
	} else if len(buf) != 3 {
		return fmt.Errorf("Sd must be 3 bytes")
	} else {
		ie.Mapped.Sd = buf
	}
	return nil
}

func (ie *SNssai) GetSd() string {
	if len(ie.Sd) > 0 {
		return hex.EncodeToString(ie.Sd)
	}
	return ""
}

func (ie *SNssai) GetMappedSd() string {
	if ie.Mapped != nil && len(ie.Mapped.Sd) > 0 {
		return hex.EncodeToString(ie.Mapped.Sd)
	}
	return ""
}

func (ie *SNssai) encode() (wire []byte, err error) {
	wire = []byte{ie.Sst}

	if len(ie.Sd) > 0 && len(ie.Sd) != 3 {
		err = fmt.Errorf("Sd must be 3 bytes")
		return
	}

	if len(ie.Sd) > 0 {
		wire = append(wire, ie.Sd...)
	}

	if ie.Mapped != nil {
		wire = append(wire, ie.Mapped.Sst)
		if len(ie.Mapped.Sd) > 0 && len(ie.Mapped.Sd) != 3 {
			err = fmt.Errorf("Sd must be 3 bytes")
			return
		}

		if len(ie.Mapped.Sd) > 0 {
			wire = append(wire, ie.Mapped.Sd...)
		}
	}
	return
}

func (ie *SNssai) decode(wire []byte) (err error) {
	switch len(wire) {
	case 1: //only Sst
		ie.Sst = wire[0]
	case 2: //Sst & Mapped-Sst
		ie.Sst = wire[0]
		ie.Mapped = &MappedSNssai{
			Sst: wire[1],
		}
	case 4: //Sst & Sd
		ie.Sst = wire[0]
		ie.Sd = make([]byte, 3)
		copy(ie.Sd, wire[1:4])
	case 5: //Sst & Sd & Mapped-Sst
		ie.Sst = wire[0]
		ie.Sd = make([]byte, 3)
		copy(ie.Sd, wire[1:4])
		ie.Mapped = &MappedSNssai{
			Sst: wire[4],
		}

	case 8: //Full
		ie.Sst = wire[0]
		ie.Sd = make([]byte, 3)
		copy(ie.Sd, wire[1:4])
		ie.Mapped = &MappedSNssai{
			Sst: wire[4],
			Sd:  make([]byte, 3),
		}
		copy(ie.Mapped.Sd, wire[5:8])
	default:
		err = fmt.Errorf("Wrong SNssai size (%d), expected value: [1, 2, 4, 5 or 8]", len(wire))
	}

	return
}
