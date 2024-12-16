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
/** this file was generated at 2024-12-16 17:55:27.326539 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * SERVICE ACCEPT
 ******************************************************/
type ServiceAccept struct {
	MmHeader
	PduSessionStatus                                                            *PduSessionStatus                       //O: TLV [50][4-34]
	PduSessionReactivationResult                                                *PduSessionReactivationResult           //O: TLV [26][4-34]
	PduSessionReactivationResultErrorCause                                      *PduSessionReactivationResultErrorCause //O: TLV-E [72][5-515]
	EapMessage                                                                  []byte                                  //O: TLV-E [78][7-1503]
	T3448Value                                                                  *GprsTimer2                             //O: TLV [6B][3]
	AdditionalRequestResult                                                     *uint8                                  //O: TLV [34][3]
	ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming                    *TrackingAreaIdentityList               //O: TLV [1D][9-114]
	ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService *TrackingAreaIdentityList               //O: TLV [1E][9-114]
}

func (msg *ServiceAccept) encode() (wire []byte, err error) {
	var buf []byte
	// O: TLV[4-34]
	if msg.PduSessionStatus != nil {
		if buf, err = encodeLV(false, uint16(2), uint16(32), msg.PduSessionStatus); err != nil {
			err = nasError("encoding PduSessionStatus [O TLV 4-34]", err)
			return
		}
		wire = append(append(wire, 0x50), buf...)
	}

	// O: TLV[4-34]
	if msg.PduSessionReactivationResult != nil {
		if buf, err = encodeLV(false, uint16(2), uint16(32), msg.PduSessionReactivationResult); err != nil {
			err = nasError("encoding PduSessionReactivationResult [O TLV 4-34]", err)
			return
		}
		wire = append(append(wire, 0x26), buf...)
	}

	// O: TLV-E[5-515]
	if msg.PduSessionReactivationResultErrorCause != nil {
		if buf, err = encodeLV(true, uint16(2), uint16(512), msg.PduSessionReactivationResultErrorCause); err != nil {
			err = nasError("encoding PduSessionReactivationResultErrorCause [O TLV-E 5-515]", err)
			return
		}
		wire = append(append(wire, 0x72), buf...)
	}

	// O: TLV-E[7-1503]
	if len(msg.EapMessage) > 0 {
		tmp := newBytesEncoder(msg.EapMessage)
		if buf, err = encodeLV(true, uint16(4), uint16(1500), tmp); err != nil {
			err = nasError("encoding EapMessage [O TLV-E 7-1503]", err)
			return
		}
		wire = append(append(wire, 0x78), buf...)
	}

	// O: TLV[3]
	if msg.T3448Value != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3448Value); err != nil {
			err = nasError("encoding T3448Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x6B), buf...)
	}

	// O: TLV[3]
	if msg.AdditionalRequestResult != nil {
		tmp := newUint8Encoder(*msg.AdditionalRequestResult)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding AdditionalRequestResult [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x34), buf...)
	}

	// O: TLV[9-114]
	if msg.ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming != nil {
		if buf, err = encodeLV(false, uint16(7), uint16(112), msg.ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming); err != nil {
			err = nasError("encoding ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming [O TLV 9-114]", err)
			return
		}
		wire = append(append(wire, 0x1D), buf...)
	}

	// O: TLV[9-114]
	if msg.ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService != nil {
		if buf, err = encodeLV(false, uint16(7), uint16(112), msg.ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService); err != nil {
			err = nasError("encoding ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService [O TLV 9-114]", err)
			return
		}
		wire = append(append(wire, 0x1E), buf...)
	}

	msg.msgType = ServiceAcceptMsgType //set message type to SERVICE ACCEPT
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *ServiceAccept) decodeBody(wire []byte) (err error) {
	offset := 0
	wireLen := len(wire)
	consumed := 0
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x50: //O: TLV[4-34]
			offset++ //consume IEI
			v := new(PduSessionStatus)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(32), v); err != nil {
				err = nasError("decoding PduSessionStatus [O TLV 4-34]", err)
				return
			}
			offset += consumed
			msg.PduSessionStatus = v
		case 0x26: //O: TLV[4-34]
			offset++ //consume IEI
			v := new(PduSessionReactivationResult)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(32), v); err != nil {
				err = nasError("decoding PduSessionReactivationResult [O TLV 4-34]", err)
				return
			}
			offset += consumed
			msg.PduSessionReactivationResult = v
		case 0x72: //O: TLV-E[5-515]
			offset++ //consume IEI
			v := new(PduSessionReactivationResultErrorCause)
			if consumed, err = decodeLV(wire[offset:], true, uint16(2), uint16(512), v); err != nil {
				err = nasError("decoding PduSessionReactivationResultErrorCause [O TLV-E 5-515]", err)
				return
			}
			offset += consumed
			msg.PduSessionReactivationResultErrorCause = v
		case 0x78: //O: TLV-E[7-1503]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), v); err != nil {
				err = nasError("decoding EapMessage [O TLV-E 7-1503]", err)
				return
			}
			offset += consumed
			msg.EapMessage = []byte(*v)
		case 0x6B: //O: TLV[3]
			offset++ //consume IEI
			v := new(GprsTimer2)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3448Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3448Value = v
		case 0x34: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding AdditionalRequestResult [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.AdditionalRequestResult = (*uint8)(v)
		case 0x1D: //O: TLV[9-114]
			offset++ //consume IEI
			v := new(TrackingAreaIdentityList)
			if consumed, err = decodeLV(wire[offset:], false, uint16(7), uint16(112), v); err != nil {
				err = nasError("decoding ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming [O TLV 9-114]", err)
				return
			}
			offset += consumed
			msg.ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming = v
		case 0x1E: //O: TLV[9-114]
			offset++ //consume IEI
			v := new(TrackingAreaIdentityList)
			if consumed, err = decodeLV(wire[offset:], false, uint16(7), uint16(112), v); err != nil {
				err = nasError("decoding ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService [O TLV 9-114]", err)
				return
			}
			offset += consumed
			msg.ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
