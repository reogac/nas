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
	"fmt"
)

type TaiListInf interface {
	GetType() uint8
	readBytes([]byte)
	bytes() []byte
}

// 9.11.3.9 5GS tracking area identity list [O TLV 9-114]
type TrackingAreaIdentityList struct {
	Lists []TaiListInf
}

func (l *TrackingAreaIdentityList) decode(wire []byte) (err error) {
	wireLen := len(wire)
	offset := 0
	var listType, numElems uint8
	var listSize int
	var listInf TaiListInf
	for offset < wireLen {
		listType = wire[offset] & 0x60 >> 5 //bit 6,5
		numElems = wire[offset]&0x1f + 1
		if numElems > 16 {
			err = fmt.Errorf("TAI list must have no more than 16 elements")
			return
		}
		//log.Infof("num elems = %d", numElems)
		switch listType {
		case 0:
			listSize = 4 + 3*int(numElems)
			listInf = new(TaiListType0)
		case 1:
			listSize = 7
			listInf = new(TaiListType1)
		case 2:
			listSize = 1 + 6*int(numElems)
			listInf = new(TaiListType2)
		default:
			err = fmt.Errorf("Unknown list type")
		}
		if err != nil {
			return
		}
		if offset+listSize > wireLen {
			err = ErrIncomplete
			return
		} else {
			listInf.readBytes(wire[offset : offset+listSize])
			l.Lists = append(l.Lists, listInf)
			offset += listSize
		}
	}
	return
}

func (l *TrackingAreaIdentityList) encode() (wire []byte, err error) {
	for _, inf := range l.Lists {
		wire = append(wire, inf.bytes()...)
	}
	return
}

// partial tracking area identity list - type 00
type TaiListType0 struct {
	PlmnId  PlmnId
	TacList []uint32 //length can't be more than 16
	Allowed bool     //spare bit (zero) or service area allowance bit
}

func (l *TaiListType0) GetType() uint8 {
	return 0
}

func (l *TaiListType0) bytes() []byte {
	size := uint8(16) //max size
	if len(l.TacList) < 16 {
		size = uint8(len(l.TacList))
	}

	header := size - 1

	if l.Allowed {
		header |= uint8(1) << 7
	}

	wire := []byte{header, 0, 0, 0}
	copy(wire[1:4], l.PlmnId.bytes[:])
	for i := 0; i < int(size); i++ {
		wire = append(wire, tacBytes(l.TacList[i])...)
	}
	//log.Infof("encode type 0 %d bytes", len(wire))
	return wire
}

// wire size must be exact
func (l *TaiListType0) readBytes(wire []byte) {
	l.Allowed = wire[0]>>7 == 1

	size := wire[0]&0x1f + 1
	copy(l.PlmnId.bytes[:], wire[1:4])
	l.TacList = make([]uint32, size)
	for i := 0; i < int(size); i++ {
		l.TacList[i] = tacFromBytes(wire[4+3*i : 7+3*i])
	}
}

// partial tracking area identity list - type 01
type TaiListType1 struct {
	PlmnId  PlmnId
	Size    uint8  //num of TACs
	Start   uint32 //start TAC
	Allowed bool   //spare bit (zero) or service area allowance bit
}

func (l *TaiListType1) GetType() uint8 {
	return 1
}

func (l *TaiListType1) bytes() []byte {
	var wire [7]byte
	wire[0] = 15
	if l.Size < 16 {
		wire[0] = l.Size - 1
	}

	wire[0] |= uint8(1) << 5 //set to type 1

	if l.Allowed {
		wire[0] |= uint8(1) << 7
	}

	copy(wire[1:4], l.PlmnId.bytes[:])
	copy(wire[4:7], tacBytes(l.Start))

	//log.Infof("encode type 1 %d bytes", len(wire))
	return wire[:]
}

// wire must be 7 bytes
func (l *TaiListType1) readBytes(wire []byte) {
	l.Allowed = wire[0]>>7 == 1

	l.Size = wire[0]&0x1f + 1
	copy(l.PlmnId.bytes[:], wire[1:4])
	l.Start = tacFromBytes(wire[4:7])
}

// partial tracking area identity list - type 10
type TaiListType2 struct {
	List    []TrackingAreaIdentity
	Allowed bool //spare bit (zero) or service area allowance bit
}

func (l *TaiListType2) GetType() uint8 {
	return 2
}

func (l *TaiListType2) bytes() []byte {
	header := uint8(2) << 5 //set to type 2

	size := uint8(16)
	if len(l.List) < 16 {
		size = uint8(len(l.List))
	}

	header |= size - 1

	if l.Allowed {
		header |= uint8(1) << 7
	}

	wire := []byte{header}

	for i := 0; i < int(size); i++ {
		buf, _ := l.List[i].encode()
		wire = append(wire, buf...)
	}
	return wire
}

// wire size must be exact
func (l *TaiListType2) readBytes(wire []byte) {
	l.Allowed = wire[0]>>7 == 1

	size := wire[0]&0x1f + 1

	var tai TrackingAreaIdentity
	l.List = make([]TrackingAreaIdentity, size)
	for i := 0; i < int(size); i++ {
		tai.readBytes(wire[1+6*i : 7+6*i])
		l.List[i] = tai
	}
}

// partial tracking area identity list - type 11
type TaiListType3 struct {
	PlmnId  PlmnId
	Allowed bool //service area allowance bit
}

func (l *TaiListType3) GetType() uint8 {
	return 3
}

func (l *TaiListType3) bytes() []byte {
	var wire [4]byte
	wire[0] = uint8(3) << 5 //set type 3
	if l.Allowed {
		wire[0] |= uint8(1) << 7
	}
	copy(wire[1:], l.PlmnId.bytes[:])
	//log.Infof("encode type 3 %d bytes", len(wire))
	return wire[:]
}

// wire size must be 4
func (l *TaiListType3) readBytes(wire []byte) {
	l.Allowed = wire[0]>>7 == 1

	copy(l.PlmnId.bytes[:], wire[1:4])
}
