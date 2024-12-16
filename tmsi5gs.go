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
