package nas

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

type Guti struct {
	PlmnId PlmnId
	AmfId  AmfId
	Tmsi   uint32
}

func (id *Guti) getIdentityType() uint8 {
	return MobileIdentity5GSType5gGuti //GUTI type
}
func (id *Guti) encode() (wire []byte, err error) {
	wire = make([]byte, 11)
	copy(wire[1:4], id.PlmnId.bytes[:])
	copy(wire[4:7], id.AmfId.bytes[:])
	binary.BigEndian.PutUint32(wire[7:11], id.Tmsi)
	wire[0] = id.getIdentityType()
	return
}

func (id *Guti) decode(wire []byte) (err error) {
	// identity type not included
	if len(wire) != 11 {
		err = ErrInvalidSize
		return
	}
	copy(id.PlmnId.bytes[:], wire[1:4])
	copy(id.AmfId.bytes[:], wire[4:7])
	id.Tmsi = binary.BigEndian.Uint32(wire[7:11])

	return
}

func (id *Guti) String() string {
	buf, _ := id.encode()
	plmnId := id.PlmnId.String()
	return plmnId + hex.EncodeToString(buf[4:])
}

func (id *Guti) Tmsi5GS() string {
	tmsi5gs := &Tmsi5Gs{
		AmfId: id.AmfId,
		Tmsi:  id.Tmsi,
	}
	return tmsi5gs.String()
}

func (id *Guti) Parse(guti string) error {
	var plmnId, remain string
	var err error
	if len(guti) == 19 {
		plmnId = guti[0:5]
		remain = guti[5:]
		remain = guti[5:]
	} else if len(guti) == 20 {
		plmnId = guti[0:6]
		remain = guti[6:]
	} else {
		return fmt.Errorf("GUTI string must be 19 or 20 in length")
	}

	if err = id.PlmnId.Parse(plmnId); err != nil {
		return nasError("decode plmnId in guti", err)
	}

	var buf []byte
	if buf, err = hex.DecodeString(remain); err != nil {
		return nasError("Hex decoding guti", err)
	}
	//buf must be 7 byte lengths
	copy(id.AmfId.bytes[:], buf[0:3])
	id.Tmsi = binary.BigEndian.Uint32(buf[3:7])
	return nil
}
