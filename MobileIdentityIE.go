/*
* Copyright [2024] [Quang Tung Thai <tqtung@etri.re.kr>]
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */
package nas

import (
	"encoding/hex"
	// "encoding/binary"
	"fmt"
)

const (
	MobileIdentity5GSTypeNoIdentity uint8 = 0x00
	MobileIdentity5GSTypeSuci       uint8 = 0x01
	MobileIdentity5GSType5gGuti     uint8 = 0x02
	MobileIdentity5GSTypeImei       uint8 = 0x03
	MobileIdentity5GSType5gSTmsi    uint8 = 0x04
	MobileIdentity5GSTypeImeisv     uint8 = 0x05
	MobileIdentity5GSTypeMac        uint8 = 0x06
	MobileIdentity5GSTypeEui64      uint8 = 0x07
)

type MobileIdentityInf interface {
	getIdentityType() uint8
	encode() ([]byte, error)
	decode([]byte) error
	String() string
}

// 9.11.3.4 5GS mobile identity [M LV-E 6-n]
type MobileIdentity struct {
	Id MobileIdentityInf
}

func (id *MobileIdentity) GetType() uint8 {
	return id.Id.getIdentityType()
}

func (id *MobileIdentity) String() string {
	return id.Id.String()
}

func (ie *MobileIdentity) encode() (wire []byte, err error) {
	wire, err = ie.Id.encode()
	return
}

func (ie *MobileIdentity) decode(wire []byte) (err error) {
	if len(wire) < 1 {
		err = ErrIncomplete
		return
	}
	idType := wire[0] & 0x07
	switch idType {
	case MobileIdentity5GSTypeSuci:
		id := new(Suci)
		if err = id.decode(wire); err == nil {
			ie.Id = id
		}

	case MobileIdentity5GSType5gGuti:
		id := new(Guti)
		if err = id.decode(wire); err == nil {
			ie.Id = id
		}

	case MobileIdentity5GSType5gSTmsi:
		id := new(Tmsi5Gs)
		if err = id.decode(wire); err == nil {
			ie.Id = id
		}

	case MobileIdentity5GSTypeImei:
		id := &Imei{
			IsSv: false,
		}
		if err = id.decode(wire); err == nil {
			ie.Id = id
		}

	case MobileIdentity5GSTypeImeisv:
		id := &Imei{
			IsSv: true,
		}
		if err = id.decode(wire); err == nil {
			ie.Id = id
		}

	case MobileIdentity5GSTypeNoIdentity:
		id := new(IdentityNone)
		if err = id.decode(wire); err == nil {
			ie.Id = id
		}

	case MobileIdentity5GSTypeMac:
		id := new(MacIdentity)
		if err = id.decode(wire); err == nil {
			ie.Id = id
		}

	case MobileIdentity5GSTypeEui64:
		id := new(Eui64Identity)
		if err = id.decode(wire); err == nil {
			ie.Id = id
		}

	default:
		err = fmt.Errorf("Unknown identity type: %d", idType)
	}

	return
}

type IdentityNone struct{}

// getIdentityType is type 000, "No identity" (TS 24.501 9.11.3.4). It was
// the 5G-GUTI's, so a caller that switched on the type took a UE without an
// identity for one with a GUTI, and panicked asserting it was a *Guti.
func (id *IdentityNone) getIdentityType() uint8 {
	return MobileIdentity5GSTypeNoIdentity
}

func (id *IdentityNone) encode() (wire []byte, err error) {
	wire = []byte{id.getIdentityType()}
	return
}

func (id *IdentityNone) decode(wire []byte) error {
	if len(wire) > 1 {
		return ErrTail
	}
	return nil
}

func (id *IdentityNone) String() string {
	return ""
}

type MacIdentity struct {
	//Mauri is the MAC address usage restriction indication: set, the MAC
	//address is not to be used as an equipment identifier (TS 24.501
	//9.11.3.4). Exported so a caller can read it and set it - it was decoded
	//into a field nothing outside this package could see
	Mauri bool
	Bytes [6]byte
}

func (id *MacIdentity) String() string {
	return hex.EncodeToString(id.Bytes[:])
}

func (id *MacIdentity) encode() (wire []byte, err error) {
	firstByte := id.getIdentityType() & 0x07 //last 3 bits
	if id.Mauri {                            //set MAURI bit
		firstByte |= macMauri
	}
	wire = append([]byte{firstByte}, id.Bytes[:]...)
	return
}

func (id *MacIdentity) decode(wire []byte) error {
	if len(wire) < 7 {
		return ErrIncomplete
	} else if len(wire) > 7 {
		return ErrTail
	}

	copy(id.Bytes[:], wire[1:])
	id.Mauri = wire[0]&macMauri != 0
	return nil
}

// macMauri is the MAURI bit of a MAC address identity's first octet: bit 4,
// counting from 1 (TS 24.501 9.11.3.4). It was read from bit 5, a spare one,
// so a MAURI the UE set decoded as clear.
const macMauri uint8 = 0x08

func (id *MacIdentity) getIdentityType() uint8 {
	return MobileIdentity5GSTypeMac //Mac type
}

type Eui64Identity struct {
	Bytes [8]byte
}

func (id *Eui64Identity) encode() (wire []byte, err error) {
	wire = append([]byte{id.getIdentityType()}, id.Bytes[:]...)
	return
}

// wire must be at least 1 byte
func (id *Eui64Identity) decode(wire []byte) (err error) {
	if len(wire) < 9 {
		return ErrIncomplete
	} else if len(wire) > 9 {
		return ErrTail
	}

	copy(id.Bytes[:], wire[1:])
	return
}

func (id *Eui64Identity) getIdentityType() uint8 {
	return MobileIdentity5GSTypeEui64 //Eui64 type
}

func (id *Eui64Identity) String() string {
	return hex.EncodeToString(id.Bytes[:])
}
