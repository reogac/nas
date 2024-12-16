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
/** this file was generated at 2024-12-16 17:55:27.328789 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * DL NAS TRANSPORT
 ******************************************************/
type DlNasTransport struct {
	MmHeader
	PayloadContainerType  uint8       //M: V [1/2]
	PayloadContainer      []byte      //M: LV-E [3-65537]
	PduSessionId          *uint8      //O: TV [12][2]
	AdditionalInformation []byte      //O: TLV [24][3-n]
	GmmCause              *uint8      //O: TV [58][2]
	BackOffTimerValue     *GprsTimer3 //O: TLV [37][3]
	LowerBoundTimerValue  *GprsTimer3 //O: TLV [3A][3]
}

func (msg *DlNasTransport) encode() (wire []byte, err error) {
	var buf []byte
	//M: V[1/2]
	v := (uint8(msg.PayloadContainerType) & 0x0f) //fill righthalf
	// M: LV-E[3-65537]
	wire = append(wire, v)

	tmp := newBytesEncoder(msg.PayloadContainer)
	if buf, err = encodeLV(true, uint16(1), uint16(0), tmp); err != nil {
		err = nasError("encoding PayloadContainer [M LV-E 3-65537]", err)
		return
	}
	wire = append(wire, buf...)

	// O: TV[2]
	if msg.PduSessionId != nil {
		wire = append(wire, []byte{0x12, uint8(*msg.PduSessionId)}...)
	}

	// O: TLV[3-n]
	if len(msg.AdditionalInformation) > 0 {
		tmp := newBytesEncoder(msg.AdditionalInformation)
		if buf, err = encodeLV(false, uint16(1), uint16(0), tmp); err != nil {
			err = nasError("encoding AdditionalInformation [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x24), buf...)
	}

	// O: TV[2]
	if msg.GmmCause != nil {
		wire = append(wire, []byte{0x58, uint8(*msg.GmmCause)}...)
	}

	// O: TLV[3]
	if msg.BackOffTimerValue != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.BackOffTimerValue); err != nil {
			err = nasError("encoding BackOffTimerValue [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x37), buf...)
	}

	// O: TLV[3]
	if msg.LowerBoundTimerValue != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.LowerBoundTimerValue); err != nil {
			err = nasError("encoding LowerBoundTimerValue [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x3A), buf...)
	}

	msg.msgType = DlNasTransportMsgType //set message type to DL NAS TRANSPORT
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *DlNasTransport) decodeBody(wire []byte) (err error) {
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// M V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding PayloadContainerType [M V 1/2]", ErrIncomplete)
		return
	}
	msg.PayloadContainerType = 0x0f & wire[offset] //righthalf
	// M LV-E[3-65537]
	offset++

	v := new(bytesDecoder)
	if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
		err = nasError("decoding PayloadContainer [M LV-E 3-65537]", err)
		return
	}
	offset += consumed
	msg.PayloadContainer = []byte(*v)
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x12: //O: TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding PduSessionId [C TV 2]", ErrIncomplete)
				return
			}
			msg.PduSessionId = new(uint8)
			offset++ //consume IEI
			*msg.PduSessionId = wire[offset]
			offset++
		case 0x24: //O: TLV[3-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding AdditionalInformation [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.AdditionalInformation = []byte(*v)
		case 0x58: //O: TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding GmmCause [O TV 2]", ErrIncomplete)
				return
			}
			msg.GmmCause = new(uint8)
			offset++ //consume IEI
			*msg.GmmCause = wire[offset]
			offset++
		case 0x37: //O: TLV[3]
			offset++ //consume IEI
			v := new(GprsTimer3)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding BackOffTimerValue [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.BackOffTimerValue = v
		case 0x3A: //O: TLV[3]
			offset++ //consume IEI
			v := new(GprsTimer3)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding LowerBoundTimerValue [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.LowerBoundTimerValue = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
