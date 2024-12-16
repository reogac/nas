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
/** this file was generated at 2024-12-16 17:55:27.329404 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * PDU SESSION ESTABLISHMENT ACCEPT
 ******************************************************/
type PduSessionEstablishmentAccept struct {
	SmHeader
	SelectedPduSessionType                 uint8                                 //M: V [1/2]
	SelectedSscMode                        uint8                                 //M: V [1/2]
	AuthorizedQosRules                     QosRules                              //M: LV-E [6-65538]
	SessionAmbr                            SessionAmbr                           //M: LV [7]
	GsmCause                               *uint8                                //O: TV [59][2]
	PduAddress                             *PduAddress                           //O: TLV [29][7-31]
	RqTimerValue                           *uint8                                //O: TV [56][2]
	SNssai                                 *SNssai                               //O: TLV [22][3-10]
	AlwaysOnPduSessionIndication           *uint8                                //O: TV [8-][1]
	MappedEpsBearerContexts                []byte                                //O: TLV-E [75][7-65538]
	EapMessage                             []byte                                //O: TLV-E [78][7-1503]
	AuthorizedQosFlowDescriptions          *QosFlowDescriptions                  //O: TLV-E [79][6-65538]
	ExtendedProtocolConfigurationOptions   *ExtendedProtocolConfigurationOptions //O: TLV-E [7B][4-65538]
	Dnn                                    *Dnn                                  //O: TLV [25][3-102]
	GsmNetworkFeatureSupport               []byte                                //O: TLV [17][3-15]
	ServingPlmnRateControl                 *uint16                               //O: TLV [18][4]
	AtsssContainer                         []byte                                //O: TLV-E [77][3-65538]
	ControlPlaneOnlyIndication             *uint8                                //O: TV [C-][1]
	IpHeaderCompressionConfiguration       []byte                                //O: TLV [66][5-257]
	EthernetHeaderCompressionConfiguration *uint8                                //O: TLV [1F][3]
	ServiceLevelAaContainer                []byte                                //O: TLV-E [72][6-n]
	ReceivedMbsContainer                   []byte                                //O: TLV-E [71][9-65538]
}

func (msg *PduSessionEstablishmentAccept) encode() (wire []byte, err error) {
	var buf []byte
	//M: V[1/2]
	v := (uint8(msg.SelectedPduSessionType) & 0x0f) //fill righthalf
	// M: V[1/2]
	v |= (uint8(msg.SelectedSscMode) & 0x0f) << 4 //fill lefthalf
	wire = append(wire, v)

	// M: LV-E[6-65538]
	if buf, err = encodeLV(true, uint16(4), uint16(0), &msg.AuthorizedQosRules); err != nil {
		err = nasError("encoding AuthorizedQosRules [M LV-E 6-65538]", err)
		return
	}
	wire = append(wire, buf...)

	// M: LV[7]
	if buf, err = encodeLV(false, uint16(6), uint16(6), &msg.SessionAmbr); err != nil {
		err = nasError("encoding SessionAmbr [M LV 7]", err)
		return
	}
	wire = append(wire, buf...)

	// O: TV[2]
	if msg.GsmCause != nil {
		wire = append(wire, []byte{0x59, uint8(*msg.GsmCause)}...)
	}

	// O: TLV[7-31]
	if msg.PduAddress != nil {
		if buf, err = encodeLV(false, uint16(5), uint16(29), msg.PduAddress); err != nil {
			err = nasError("encoding PduAddress [O TLV 7-31]", err)
			return
		}
		wire = append(append(wire, 0x29), buf...)
	}

	// O: TV[2]
	if msg.RqTimerValue != nil {
		wire = append(wire, []byte{0x56, uint8(*msg.RqTimerValue)}...)
	}

	// O: TLV[3-10]
	if msg.SNssai != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(8), msg.SNssai); err != nil {
			err = nasError("encoding SNssai [O TLV 3-10]", err)
			return
		}
		wire = append(append(wire, 0x22), buf...)
	}

	// O: TV[1]
	if msg.AlwaysOnPduSessionIndication != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x08<<4)|(uint8(*msg.AlwaysOnPduSessionIndication)&0x0f))
	}

	// O: TLV-E[7-65538]
	if len(msg.MappedEpsBearerContexts) > 0 {
		tmp := newBytesEncoder(msg.MappedEpsBearerContexts)
		if buf, err = encodeLV(true, uint16(4), uint16(0), tmp); err != nil {
			err = nasError("encoding MappedEpsBearerContexts [O TLV-E 7-65538]", err)
			return
		}
		wire = append(append(wire, 0x75), buf...)
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

	// O: TLV-E[6-65538]
	if msg.AuthorizedQosFlowDescriptions != nil {
		if buf, err = encodeLV(true, uint16(3), uint16(0), msg.AuthorizedQosFlowDescriptions); err != nil {
			err = nasError("encoding AuthorizedQosFlowDescriptions [O TLV-E 6-65538]", err)
			return
		}
		wire = append(append(wire, 0x79), buf...)
	}

	// O: TLV-E[4-65538]
	if msg.ExtendedProtocolConfigurationOptions != nil {
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.ExtendedProtocolConfigurationOptions); err != nil {
			err = nasError("encoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x7B), buf...)
	}

	// O: TLV[3-102]
	if msg.Dnn != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(100), msg.Dnn); err != nil {
			err = nasError("encoding Dnn [O TLV 3-102]", err)
			return
		}
		wire = append(append(wire, 0x25), buf...)
	}

	// O: TLV[3-15]
	if len(msg.GsmNetworkFeatureSupport) > 0 {
		tmp := newBytesEncoder(msg.GsmNetworkFeatureSupport)
		if buf, err = encodeLV(false, uint16(1), uint16(13), tmp); err != nil {
			err = nasError("encoding GsmNetworkFeatureSupport [O TLV 3-15]", err)
			return
		}
		wire = append(append(wire, 0x17), buf...)
	}

	// O: TLV[4]
	if msg.ServingPlmnRateControl != nil {
		tmp := newUint16Encoder(*msg.ServingPlmnRateControl)
		if buf, err = encodeLV(false, uint16(2), uint16(2), tmp); err != nil {
			err = nasError("encoding ServingPlmnRateControl [O TLV 4]", err)
			return
		}
		wire = append(append(wire, 0x18), buf...)
	}

	// O: TLV-E[3-65538]
	if len(msg.AtsssContainer) > 0 {
		tmp := newBytesEncoder(msg.AtsssContainer)
		if buf, err = encodeLV(true, uint16(0), uint16(0), tmp); err != nil {
			err = nasError("encoding AtsssContainer [O TLV-E 3-65538]", err)
			return
		}
		wire = append(append(wire, 0x77), buf...)
	}

	// O: TV[1]
	if msg.ControlPlaneOnlyIndication != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0C<<4)|(uint8(*msg.ControlPlaneOnlyIndication)&0x0f))
	}

	// O: TLV[5-257]
	if len(msg.IpHeaderCompressionConfiguration) > 0 {
		tmp := newBytesEncoder(msg.IpHeaderCompressionConfiguration)
		if buf, err = encodeLV(false, uint16(3), uint16(255), tmp); err != nil {
			err = nasError("encoding IpHeaderCompressionConfiguration [O TLV 5-257]", err)
			return
		}
		wire = append(append(wire, 0x66), buf...)
	}

	// O: TLV[3]
	if msg.EthernetHeaderCompressionConfiguration != nil {
		tmp := newUint8Encoder(*msg.EthernetHeaderCompressionConfiguration)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding EthernetHeaderCompressionConfiguration [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x1F), buf...)
	}

	// O: TLV-E[6-n]
	if len(msg.ServiceLevelAaContainer) > 0 {
		tmp := newBytesEncoder(msg.ServiceLevelAaContainer)
		if buf, err = encodeLV(true, uint16(3), uint16(0), tmp); err != nil {
			err = nasError("encoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
			return
		}
		wire = append(append(wire, 0x72), buf...)
	}

	// O: TLV-E[9-65538]
	if len(msg.ReceivedMbsContainer) > 0 {
		tmp := newBytesEncoder(msg.ReceivedMbsContainer)
		if buf, err = encodeLV(true, uint16(6), uint16(0), tmp); err != nil {
			err = nasError("encoding ReceivedMbsContainer [O TLV-E 9-65538]", err)
			return
		}
		wire = append(append(wire, 0x71), buf...)
	}

	msg.msgType = PduSessionEstablishmentAcceptMsgType //set message type to PDU SESSION ESTABLISHMENT ACCEPT
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *PduSessionEstablishmentAccept) decodeBody(wire []byte) (err error) {
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// M V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding SelectedPduSessionType [M V 1/2]", ErrIncomplete)
		return
	}
	msg.SelectedPduSessionType = 0x0f & wire[offset] //righthalf
	// M V[1/2]
	msg.SelectedSscMode = (0xf0 & wire[offset]) >> 4 //lefthalf
	offset++

	// M LV-E[6-65538]
	if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), &msg.AuthorizedQosRules); err != nil {
		err = nasError("decoding AuthorizedQosRules [M LV-E 6-65538]", err)
		return
	}
	offset += consumed
	// M LV[7]
	if consumed, err = decodeLV(wire[offset:], false, uint16(6), uint16(6), &msg.SessionAmbr); err != nil {
		err = nasError("decoding SessionAmbr [M LV 7]", err)
		return
	}
	offset += consumed
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x59: //O: TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding GsmCause [O TV 2]", ErrIncomplete)
				return
			}
			msg.GsmCause = new(uint8)
			offset++ //consume IEI
			*msg.GsmCause = wire[offset]
			offset++
		case 0x29: //O: TLV[7-31]
			offset++ //consume IEI
			v := new(PduAddress)
			if consumed, err = decodeLV(wire[offset:], false, uint16(5), uint16(29), v); err != nil {
				err = nasError("decoding PduAddress [O TLV 7-31]", err)
				return
			}
			offset += consumed
			msg.PduAddress = v
		case 0x56: //O: TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding RqTimerValue [O TV 2]", ErrIncomplete)
				return
			}
			msg.RqTimerValue = new(uint8)
			offset++ //consume IEI
			*msg.RqTimerValue = wire[offset]
			offset++
		case 0x22: //O: TLV[3-10]
			offset++ //consume IEI
			v := new(SNssai)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(8), v); err != nil {
				err = nasError("decoding SNssai [O TLV 3-10]", err)
				return
			}
			offset += consumed
			msg.SNssai = v
		case 0x08: //O: TV[1]
			msg.AlwaysOnPduSessionIndication = new(uint8)
			*msg.AlwaysOnPduSessionIndication = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x75: //O: TLV-E[7-65538]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), v); err != nil {
				err = nasError("decoding MappedEpsBearerContexts [O TLV-E 7-65538]", err)
				return
			}
			offset += consumed
			msg.MappedEpsBearerContexts = []byte(*v)
		case 0x78: //O: TLV-E[7-1503]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), v); err != nil {
				err = nasError("decoding EapMessage [O TLV-E 7-1503]", err)
				return
			}
			offset += consumed
			msg.EapMessage = []byte(*v)
		case 0x79: //O: TLV-E[6-65538]
			offset++ //consume IEI
			v := new(QosFlowDescriptions)
			if consumed, err = decodeLV(wire[offset:], true, uint16(3), uint16(0), v); err != nil {
				err = nasError("decoding AuthorizedQosFlowDescriptions [O TLV-E 6-65538]", err)
				return
			}
			offset += consumed
			msg.AuthorizedQosFlowDescriptions = v
		case 0x7B: //O: TLV-E[4-65538]
			offset++ //consume IEI
			v := new(ExtendedProtocolConfigurationOptions)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.ExtendedProtocolConfigurationOptions = v
		case 0x25: //O: TLV[3-102]
			offset++ //consume IEI
			v := new(Dnn)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(100), v); err != nil {
				err = nasError("decoding Dnn [O TLV 3-102]", err)
				return
			}
			offset += consumed
			msg.Dnn = v
		case 0x17: //O: TLV[3-15]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(13), v); err != nil {
				err = nasError("decoding GsmNetworkFeatureSupport [O TLV 3-15]", err)
				return
			}
			offset += consumed
			msg.GsmNetworkFeatureSupport = []byte(*v)
		case 0x18: //O: TLV[4]
			offset++ //consume IEI
			v := new(uint16Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(2), v); err != nil {
				err = nasError("decoding ServingPlmnRateControl [O TLV 4]", err)
				return
			}
			offset += consumed
			msg.ServingPlmnRateControl = (*uint16)(v)
		case 0x77: //O: TLV-E[3-65538]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding AtsssContainer [O TLV-E 3-65538]", err)
				return
			}
			offset += consumed
			msg.AtsssContainer = []byte(*v)
		case 0x0C: //O: TV[1]
			msg.ControlPlaneOnlyIndication = new(uint8)
			*msg.ControlPlaneOnlyIndication = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x66: //O: TLV[5-257]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(255), v); err != nil {
				err = nasError("decoding IpHeaderCompressionConfiguration [O TLV 5-257]", err)
				return
			}
			offset += consumed
			msg.IpHeaderCompressionConfiguration = []byte(*v)
		case 0x1F: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding EthernetHeaderCompressionConfiguration [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.EthernetHeaderCompressionConfiguration = (*uint8)(v)
		case 0x72: //O: TLV-E[6-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(3), uint16(0), v); err != nil {
				err = nasError("decoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
				return
			}
			offset += consumed
			msg.ServiceLevelAaContainer = []byte(*v)
		case 0x71: //O: TLV-E[9-65538]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(6), uint16(0), v); err != nil {
				err = nasError("decoding ReceivedMbsContainer [O TLV-E 9-65538]", err)
				return
			}
			offset += consumed
			msg.ReceivedMbsContainer = []byte(*v)
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
