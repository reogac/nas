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
/** this file was generated at 2024-12-16 17:55:27.330669 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * PDU SESSION MODIFICATION REJECT
 ******************************************************/
type PduSessionModificationReject struct {
	SmHeader
	GsmCause                             uint8                                 //M: V [1]
	BackOffTimerValue                    *GprsTimer3                           //O: TLV [37][3]
	GsmCongestionReAttemptIndicator      *uint8                                //O: TLV [61][3]
	ExtendedProtocolConfigurationOptions *ExtendedProtocolConfigurationOptions //O: TLV-E [7B][4-65538]
	ReAttemptIndicator                   *uint8                                //O: TLV [1D][3]
}

func (msg *PduSessionModificationReject) encode() (wire []byte, err error) {
	var buf []byte
	// M: V[1]
	wire = append(wire, uint8(msg.GsmCause))

	// O: TLV[3]
	if msg.BackOffTimerValue != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.BackOffTimerValue); err != nil {
			err = nasError("encoding BackOffTimerValue [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x37), buf...)
	}

	// O: TLV[3]
	if msg.GsmCongestionReAttemptIndicator != nil {
		tmp := newUint8Encoder(*msg.GsmCongestionReAttemptIndicator)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding GsmCongestionReAttemptIndicator [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x61), buf...)
	}

	// O: TLV-E[4-65538]
	if msg.ExtendedProtocolConfigurationOptions != nil {
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.ExtendedProtocolConfigurationOptions); err != nil {
			err = nasError("encoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x7B), buf...)
	}

	// O: TLV[3]
	if msg.ReAttemptIndicator != nil {
		tmp := newUint8Encoder(*msg.ReAttemptIndicator)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding ReAttemptIndicator [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x1D), buf...)
	}

	msg.msgType = PduSessionModificationRejectMsgType //set message type to PDU SESSION MODIFICATION REJECT
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *PduSessionModificationReject) decodeBody(wire []byte) (err error) {
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// M V[1]
	if offset+1 > wireLen {
		err = nasError("decoding GsmCause [M V 1]", ErrIncomplete)
		return
	}
	msg.GsmCause = wire[offset]
	offset++

	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x37: //O: TLV[3]
			offset++ //consume IEI
			v := new(GprsTimer3)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding BackOffTimerValue [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.BackOffTimerValue = v
		case 0x61: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding GsmCongestionReAttemptIndicator [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.GsmCongestionReAttemptIndicator = (*uint8)(v)
		case 0x7B: //O: TLV-E[4-65538]
			offset++ //consume IEI
			v := new(ExtendedProtocolConfigurationOptions)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.ExtendedProtocolConfigurationOptions = v
		case 0x1D: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding ReAttemptIndicator [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.ReAttemptIndicator = (*uint8)(v)
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
