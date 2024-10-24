/**generated time: 2024-07-17 15:11:00.946299**/

package nas

/*******************************************************
 * PDU SESSION ESTABLISHMENT ACCEPT
 ******************************************************/
type PduSessionEstablishmentAccept struct {
	SmHeader
	SelectedPduSessionType                 Uint8                                 //V [1/2]
	SelectedSscMode                        Uint8                                 //V [1/2]
	AuthorizedQosRules                     QosRules                              //LV-E [6-65538]
	SessionAmbr                            SessionAmbr                           //LV [7]
	GsmCause                               *Uint8                                //TV [59][2]
	PduAddress                             *PduAddress                           //TLV [29][7-31]
	RqTimerValue                           *Uint8                                //TV [56][2]
	SNssai                                 *SNssai                               //TLV [22][3-10]
	AlwaysOnPduSessionIndication           *Uint8                                //TV [8-][1]
	MappedEpsBearerContexts                *Bytes                                //TLV-E [75][7-65538]
	EapMessage                             *Bytes                                //TLV-E [78][7-1503]
	AuthorizedQosFlowDescriptions          *QosFlowDescriptions                  //TLV-E [79][6-65538]
	ExtendedProtocolConfigurationOptions   *ExtendedProtocolConfigurationOptions //TLV-E [7B][4-65538]
	Dnn                                    *Dnn                                  //TLV [25][3-102]
	GsmNetworkFeatureSupport               *Bytes                                //TLV [17][3-15]
	ServingPlmnRateControl                 *Uint16                               //TLV [18][4]
	AtsssContainer                         *Bytes                                //TLV-E [77][3-65538]
	ControlPlaneOnlyIndication             *Uint8                                //TV [C-][1]
	IpHeaderCompressionConfiguration       *Bytes                                //TLV [66][5-257]
	EthernetHeaderCompressionConfiguration *Uint8                                //TLV [1F][3]
	ServiceLevelAaContainer                *Bytes                                //TLV-E [72][6-n]
	ReceivedMbsContainer                   *Bytes                                //TLV-E [71][9-65538]
}

func (msg *PduSessionEstablishmentAccept) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding PduSessionEstablishmentAccept", err)
		}
	}()
	var buf []byte
	//V[1/2]
	v := (uint8(msg.SelectedPduSessionType) & 0x0f) //fill righthalf
	// V[1/2]
	v |= (uint8(msg.SelectedSscMode) & 0x0f) << 4 //fill lefthalf
	wire = append(wire, v)

	// LV-E[6-65538]
	if buf, err = encodeLV(true, uint16(4), uint16(0), &msg.AuthorizedQosRules); err != nil {
		err = nasError("encoding AuthorizedQosRules [M LV-E 6-65538]", err)
		return
	}
	wire = append(wire, buf...)

	// LV[7]
	if buf, err = encodeLV(false, uint16(6), uint16(6), &msg.SessionAmbr); err != nil {
		err = nasError("encoding SessionAmbr [M LV 7]", err)
		return
	}
	wire = append(wire, buf...)

	if msg.GsmCause != nil {
		//TV[2]
		wire = append(wire, []byte{0x59, uint8(*msg.GsmCause)}...)
	}

	if msg.PduAddress != nil {
		// TLV[7-31]
		if buf, err = encodeLV(false, uint16(5), uint16(29), msg.PduAddress); err != nil {
			err = nasError("encoding PduAddress [O TLV 7-31]", err)
			return
		}
		wire = append(append(wire, 0x29), buf...)
	}

	if msg.RqTimerValue != nil {
		//TV[2]
		wire = append(wire, []byte{0x56, uint8(*msg.RqTimerValue)}...)
	}

	if msg.SNssai != nil {
		// TLV[3-10]
		if buf, err = encodeLV(false, uint16(1), uint16(8), msg.SNssai); err != nil {
			err = nasError("encoding SNssai [O TLV 3-10]", err)
			return
		}
		wire = append(append(wire, 0x22), buf...)
	}

	if msg.AlwaysOnPduSessionIndication != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x08<<4)|(uint8(*msg.AlwaysOnPduSessionIndication)&0x0f))
	}

	if msg.MappedEpsBearerContexts != nil {
		// TLV-E[7-65538]
		if buf, err = encodeLV(true, uint16(4), uint16(0), msg.MappedEpsBearerContexts); err != nil {
			err = nasError("encoding MappedEpsBearerContexts [O TLV-E 7-65538]", err)
			return
		}
		wire = append(append(wire, 0x75), buf...)
	}

	if msg.EapMessage != nil {
		// TLV-E[7-1503]
		if buf, err = encodeLV(true, uint16(4), uint16(1500), msg.EapMessage); err != nil {
			err = nasError("encoding EapMessage [O TLV-E 7-1503]", err)
			return
		}
		wire = append(append(wire, 0x78), buf...)
	}

	if msg.AuthorizedQosFlowDescriptions != nil {
		// TLV-E[6-65538]
		if buf, err = encodeLV(true, uint16(3), uint16(0), msg.AuthorizedQosFlowDescriptions); err != nil {
			err = nasError("encoding AuthorizedQosFlowDescriptions [O TLV-E 6-65538]", err)
			return
		}
		wire = append(append(wire, 0x79), buf...)
	}

	if msg.ExtendedProtocolConfigurationOptions != nil {
		// TLV-E[4-65538]
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.ExtendedProtocolConfigurationOptions); err != nil {
			err = nasError("encoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x7B), buf...)
	}

	if msg.Dnn != nil {
		// TLV[3-102]
		if buf, err = encodeLV(false, uint16(1), uint16(100), msg.Dnn); err != nil {
			err = nasError("encoding Dnn [O TLV 3-102]", err)
			return
		}
		wire = append(append(wire, 0x25), buf...)
	}

	if msg.GsmNetworkFeatureSupport != nil {
		// TLV[3-15]
		if buf, err = encodeLV(false, uint16(1), uint16(13), msg.GsmNetworkFeatureSupport); err != nil {
			err = nasError("encoding GsmNetworkFeatureSupport [O TLV 3-15]", err)
			return
		}
		wire = append(append(wire, 0x17), buf...)
	}

	if msg.ServingPlmnRateControl != nil {
		// TLV[4]
		if buf, err = encodeLV(false, uint16(2), uint16(2), msg.ServingPlmnRateControl); err != nil {
			err = nasError("encoding ServingPlmnRateControl [O TLV 4]", err)
			return
		}
		wire = append(append(wire, 0x18), buf...)
	}

	if msg.AtsssContainer != nil {
		// TLV-E[3-65538]
		if buf, err = encodeLV(true, uint16(0), uint16(0), msg.AtsssContainer); err != nil {
			err = nasError("encoding AtsssContainer [O TLV-E 3-65538]", err)
			return
		}
		wire = append(append(wire, 0x77), buf...)
	}

	if msg.ControlPlaneOnlyIndication != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0C<<4)|(uint8(*msg.ControlPlaneOnlyIndication)&0x0f))
	}

	if msg.IpHeaderCompressionConfiguration != nil {
		// TLV[5-257]
		if buf, err = encodeLV(false, uint16(3), uint16(255), msg.IpHeaderCompressionConfiguration); err != nil {
			err = nasError("encoding IpHeaderCompressionConfiguration [O TLV 5-257]", err)
			return
		}
		wire = append(append(wire, 0x66), buf...)
	}

	if msg.EthernetHeaderCompressionConfiguration != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.EthernetHeaderCompressionConfiguration); err != nil {
			err = nasError("encoding EthernetHeaderCompressionConfiguration [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x1F), buf...)
	}

	if msg.ServiceLevelAaContainer != nil {
		// TLV-E[6-n]
		if buf, err = encodeLV(true, uint16(3), uint16(0), msg.ServiceLevelAaContainer); err != nil {
			err = nasError("encoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
			return
		}
		wire = append(append(wire, 0x72), buf...)
	}

	if msg.ReceivedMbsContainer != nil {
		// TLV-E[9-65538]
		if buf, err = encodeLV(true, uint16(6), uint16(0), msg.ReceivedMbsContainer); err != nil {
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
	defer func() {
		if err != nil {
			err = nasError("decoding PduSessionEstablishmentAccept", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding SelectedPduSessionType [M V 1/2]", ErrIncomplete)
		return
	}
	msg.SelectedPduSessionType = Uint8(0x0f & wire[offset]) //righthalf
	// V[1/2]
	msg.SelectedSscMode = Uint8((0xf0 & wire[offset]) >> 4) //lefthalf
	offset++

	// LV-E[6-65538]
	if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), &msg.AuthorizedQosRules); err != nil {
		err = nasError("decoding AuthorizedQosRules [M LV-E 6-65538]", err)
		return
	}
	offset += consumed
	// LV[7]
	if consumed, err = decodeLV(wire[offset:], false, uint16(6), uint16(6), &msg.SessionAmbr); err != nil {
		err = nasError("decoding SessionAmbr [M LV 7]", err)
		return
	}
	offset += consumed
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x59: //TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding GsmCause [O TV 2]", ErrIncomplete)
				return
			}
			msg.GsmCause = new(Uint8)
			offset++ //consume IEI
			*msg.GsmCause = Uint8(wire[offset])
			offset++
		case 0x29: //TLV[7-31]
			offset++ //consume IEI
			v := &PduAddress{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(5), uint16(29), v); err != nil {
				err = nasError("decoding PduAddress [O TLV 7-31]", err)
				return
			}
			offset += consumed
			msg.PduAddress = v
		case 0x56: //TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding RqTimerValue [O TV 2]", ErrIncomplete)
				return
			}
			msg.RqTimerValue = new(Uint8)
			offset++ //consume IEI
			*msg.RqTimerValue = Uint8(wire[offset])
			offset++
		case 0x22: //TLV[3-10]
			offset++ //consume IEI
			v := &SNssai{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(8), v); err != nil {
				err = nasError("decoding SNssai [O TLV 3-10]", err)
				return
			}
			offset += consumed
			msg.SNssai = v
		case 0x08: //TV[1]
			msg.AlwaysOnPduSessionIndication = new(Uint8)
			*msg.AlwaysOnPduSessionIndication = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x75: //TLV-E[7-65538]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), v); err != nil {
				err = nasError("decoding MappedEpsBearerContexts [O TLV-E 7-65538]", err)
				return
			}
			offset += consumed
			msg.MappedEpsBearerContexts = v
		case 0x78: //TLV-E[7-1503]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), v); err != nil {
				err = nasError("decoding EapMessage [O TLV-E 7-1503]", err)
				return
			}
			offset += consumed
			msg.EapMessage = v
		case 0x79: //TLV-E[6-65538]
			offset++ //consume IEI
			v := &QosFlowDescriptions{}
			if consumed, err = decodeLV(wire[offset:], true, uint16(3), uint16(0), v); err != nil {
				err = nasError("decoding AuthorizedQosFlowDescriptions [O TLV-E 6-65538]", err)
				return
			}
			offset += consumed
			msg.AuthorizedQosFlowDescriptions = v
		case 0x7B: //TLV-E[4-65538]
			offset++ //consume IEI
			v := &ExtendedProtocolConfigurationOptions{}
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.ExtendedProtocolConfigurationOptions = v
		case 0x25: //TLV[3-102]
			offset++ //consume IEI
			v := &Dnn{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(100), v); err != nil {
				err = nasError("decoding Dnn [O TLV 3-102]", err)
				return
			}
			offset += consumed
			msg.Dnn = v
		case 0x17: //TLV[3-15]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(13), v); err != nil {
				err = nasError("decoding GsmNetworkFeatureSupport [O TLV 3-15]", err)
				return
			}
			offset += consumed
			msg.GsmNetworkFeatureSupport = v
		case 0x18: //TLV[4]
			offset++ //consume IEI
			v := new(Uint16)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(2), v); err != nil {
				err = nasError("decoding ServingPlmnRateControl [O TLV 4]", err)
				return
			}
			offset += consumed
			msg.ServingPlmnRateControl = v
		case 0x77: //TLV-E[3-65538]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding AtsssContainer [O TLV-E 3-65538]", err)
				return
			}
			offset += consumed
			msg.AtsssContainer = v
		case 0x0C: //TV[1]
			msg.ControlPlaneOnlyIndication = new(Uint8)
			*msg.ControlPlaneOnlyIndication = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x66: //TLV[5-257]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(255), v); err != nil {
				err = nasError("decoding IpHeaderCompressionConfiguration [O TLV 5-257]", err)
				return
			}
			offset += consumed
			msg.IpHeaderCompressionConfiguration = v
		case 0x1F: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding EthernetHeaderCompressionConfiguration [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.EthernetHeaderCompressionConfiguration = v
		case 0x72: //TLV-E[6-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(3), uint16(0), v); err != nil {
				err = nasError("decoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
				return
			}
			offset += consumed
			msg.ServiceLevelAaContainer = v
		case 0x71: //TLV-E[9-65538]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(6), uint16(0), v); err != nil {
				err = nasError("decoding ReceivedMbsContainer [O TLV-E 9-65538]", err)
				return
			}
			offset += consumed
			msg.ReceivedMbsContainer = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
