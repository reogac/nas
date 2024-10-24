package nas

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

type Tmsi5Gs struct {
	AmfId AmfId //ignore the AmfRegion (set to zero)
	Tmsi  uint32
}

func (id *Tmsi5Gs) getIdentityType() uint8 {
	return MobileIdentity5GSType5gSTmsi //5GsTmsi type
}

func (id *Tmsi5Gs) encode() (wire []byte, err error) {
	wire = make([]byte, 7)
	copy(wire[1:3], id.AmfId.bytes[1:]) //ignore AmfRegion
	binary.BigEndian.PutUint32(wire[3:7], id.Tmsi)
	wire[0] = id.getIdentityType()
	return
}

func (id *Tmsi5Gs) decode(wire []byte) (err error) {
	// identity type not included
	if len(wire) != 7 {
		err = ErrInvalidSize
		return
	}
	copy(id.AmfId.bytes[1:], wire[1:3])
	id.Tmsi = binary.BigEndian.Uint32(wire[3:7])
	return
}

func (id *Tmsi5Gs) String() string {
	buf, _ := id.encode()
	return hex.EncodeToString(buf[1:]) //strip of the identity type byte
}

func (id *Tmsi5Gs) Parse(s string) error {
	if len(s) != 12 {
		return fmt.Errorf("Tmsi5GS must be 12 digits")
	}
	if wire, err := hex.DecodeString(s); err != nil {
		return nasError("Hex decode tmsi5gs", err)
	} else {
		copy(id.AmfId.bytes[1:], wire[0:2])
		id.Tmsi = binary.BigEndian.Uint32(wire[2:6])
	}
	return nil
}
