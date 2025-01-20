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

// 9.11.3.5 5GS network feature support [O TLV 3-5]
type NetworkFeatureSupport struct {
	bytes     [3]byte
	hasOctet1 bool
	hasOctet2 bool
}

func (ie *NetworkFeatureSupport) encode() (wire []byte, err error) {
	encodeLen := 1
	if ie.hasOctet1 {
		encodeLen = 2
	}
	if ie.hasOctet2 {
		encodeLen = 3
	}
	wire = ie.bytes[:encodeLen]
	return
}

func (ie *NetworkFeatureSupport) decode(wire []byte) (err error) {
	switch len(wire) {
	case 0:
		return ErrIncomplete
	case 1:
	case 2:
		ie.hasOctet1 = true
	case 3:
		ie.hasOctet1 = true
		ie.hasOctet2 = true
	default:
		return ErrTail
	}
	copy(ie.bytes[:], wire)
	return
}

// octet 0, bit 7
func (ie *NetworkFeatureSupport) GetMPSI() bool {
	return getBit(ie.bytes[0], 7) == 1
}

func (ie *NetworkFeatureSupport) SetMPSI(f bool) {
	if f {
		ie.bytes[0] = setBit(ie.bytes[0], 7)
	} else {
		ie.bytes[0] = clearBit(ie.bytes[0], 7)
	}
}

// octet 0, bit 6
func (ie *NetworkFeatureSupport) GetIWKN26() bool {
	return getBit(ie.bytes[0], 6) == 1
}

func (ie *NetworkFeatureSupport) SetIWKN26(f bool) {
	if f {
		ie.bytes[0] = setBit(ie.bytes[0], 6)
	} else {
		ie.bytes[0] = clearBit(ie.bytes[0], 6)
	}
}

// octet 0, bit 5,4
func (ie *NetworkFeatureSupport) GetEMF() uint8 {
	return (ie.bytes[0] & bitMask(4, 2)) >> 4
}
func (ie *NetworkFeatureSupport) SetEMF(v uint8) {
	mask := uint8(0x03) //2 bits
	ie.bytes[0] = (ie.bytes[0] & ^(mask << 4)) | ((v & mask) << 4)
}

// octet 0, bit 3,2
func (ie *NetworkFeatureSupport) GetEMC() uint8 {
	return (ie.bytes[0] & bitMask(2, 2)) >> 2
}

func (ie *NetworkFeatureSupport) SetEMC(v uint8) {
	mask := uint8(0x03)
	ie.bytes[0] = (ie.bytes[0] & ^(mask << 2)) | ((v & mask) << 2)
}

// octet 0, bit 1
func (ie *NetworkFeatureSupport) GetIMSVoPSN3GPP() bool {
	if len(ie.bytes) > 1 {
		return getBit(ie.bytes[0], 1) == 1
	}
	return false
}

func (ie *NetworkFeatureSupport) SetIMSVoPSN3GPP(f bool) {
	if f {
		ie.bytes[0] = setBit(ie.bytes[0], 1)
	} else {
		ie.bytes[0] = clearBit(ie.bytes[0], 1)
	}
}

// octet 0, bit 0
func (ie *NetworkFeatureSupport) GetIMSVoPS3GPP() bool {
	return getBit(ie.bytes[0], 0) == 0
}
func (ie *NetworkFeatureSupport) SetIMSVoPS3GPP(f bool) {
	if f {
		ie.bytes[0] = setBit(ie.bytes[0], 0)
	} else {
		ie.bytes[0] = clearBit(ie.bytes[0], 0)
	}
}

// octet 1, bit 1
func (ie *NetworkFeatureSupport) GetMCSI() bool {
	if ie.hasOctet1 {
		return getBit(ie.bytes[1], 1) == 1
	}
	return false
}

func (ie *NetworkFeatureSupport) SetMCSI(f bool) {
	ie.hasOctet1 = true
	if f {
		ie.bytes[1] = setBit(ie.bytes[1], 1)
	} else {
		ie.bytes[1] = clearBit(ie.bytes[1], 1)
	}
}

// octet 1, bit 0
func (ie *NetworkFeatureSupport) GetEMCN() bool {
	if ie.hasOctet1 {
		return getBit(ie.bytes[1], 0) == 0
	}
	return false
}

func (ie *NetworkFeatureSupport) SetEMCN(f bool) {
	ie.hasOctet1 = true
	if f {
		ie.bytes[1] = setBit(ie.bytes[1], 0)
	} else {
		ie.bytes[1] = clearBit(ie.bytes[1], 0)
	}
}
