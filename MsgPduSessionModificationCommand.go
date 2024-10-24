/**generated time: 2024-07-17 15:11:00.948290**/

package nas

/*******************************************************
 * PDU SESSION MODIFICATION COMMAND
 ******************************************************/
type PduSessionModificationCommand struct {
	SmHeader
	GsmCause                               *Uint8                                //TV [59][2]
	SessionAmbr                            *SessionAmbr                          //TLV [2A][8]
	RqTimerValue                           *Uint8                                //TV [56][2]
	AlwaysOnPduSessionIndication           *Uint8                                //TV [8-][1]
	AuthorizedQosRules                     *QosRules                             //TLV-E [7A][7-65538]
	MappedEpsBearerContexts                *Bytes                                //TLV-E [75][7-65538]
	AuthorizedQosFlowDescriptions          *QosFlowDescriptions                  //TLV-E [79][6-65538]
	ExtendedProtocolConfigurationOptions   *ExtendedProtocolConfigurationOptions //TLV-E [7B][4-65538]
	AtsssContainer                         *Bytes                                //TLV-E [77][3-65538]
	IpHeaderCompressionConfiguration       *Bytes                                //TLV [66][5-257]
	PortManagementInformationContainer     *Bytes                                //TLV-E [74][4-65538]
	ServingPlmnRateControl                 *Uint16                               //TLV [1E][4]
	EthernetHeaderCompressionConfiguration *Uint8                                //TLV [1F][3]
	ReceivedMbsContainer                   *Bytes                                //TLV-E [71][9-65538]
	ServiceLevelAaContainer                *Bytes                                //TLV-E [72][6-n]
}

func (msg *PduSessionModificationCommand) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding PduSessionModificationCommand", err)
		}
	}()
	var buf []byte
	if msg.GsmCause != nil {
		//TV[2]
		wire = append(wire, []byte{0x59, uint8(*msg.GsmCause)}...)
	}

	if msg.SessionAmbr != nil {
		// TLV[8]
		if buf, err = encodeLV(false, uint16(6), uint16(6), msg.SessionAmbr); err != nil {
			err = nasError("encoding SessionAmbr [O TLV 8]", err)
			return
		}
		wire = append(append(wire, 0x2A), buf...)
	}

	if msg.RqTimerValue != nil {
		//TV[2]
		wire = append(wire, []byte{0x56, uint8(*msg.RqTimerValue)}...)
	}

	if msg.AlwaysOnPduSessionIndication != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x08<<4)|(uint8(*msg.AlwaysOnPduSessionIndication)&0x0f))
	}

	if msg.AuthorizedQosRules != nil {
		// TLV-E[7-65538]
		if buf, err = encodeLV(true, uint16(4), uint16(0), msg.AuthorizedQosRules); err != nil {
			err = nasError("encoding AuthorizedQosRules [O TLV-E 7-65538]", err)
			return
		}
		wire = append(append(wire, 0x7A), buf...)
	}

	if msg.MappedEpsBearerContexts != nil {
		// TLV-E[7-65538]
		if buf, err = encodeLV(true, uint16(4), uint16(0), msg.MappedEpsBearerContexts); err != nil {
			err = nasError("encoding MappedEpsBearerContexts [O TLV-E 7-65538]", err)
			return
		}
		wire = append(append(wire, 0x75), buf...)
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

	if msg.AtsssContainer != nil {
		// TLV-E[3-65538]
		if buf, err = encodeLV(true, uint16(0), uint16(0), msg.AtsssContainer); err != nil {
			err = nasError("encoding AtsssContainer [O TLV-E 3-65538]", err)
			return
		}
		wire = append(append(wire, 0x77), buf...)
	}

	if msg.IpHeaderCompressionConfiguration != nil {
		// TLV[5-257]
		if buf, err = encodeLV(false, uint16(3), uint16(255), msg.IpHeaderCompressionConfiguration); err != nil {
			err = nasError("encoding IpHeaderCompressionConfiguration [O TLV 5-257]", err)
			return
		}
		wire = append(append(wire, 0x66), buf...)
	}

	if msg.PortManagementInformationContainer != nil {
		// TLV-E[4-65538]
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.PortManagementInformationContainer); err != nil {
			err = nasError("encoding PortManagementInformationContainer [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x74), buf...)
	}

	if msg.ServingPlmnRateControl != nil {
		// TLV[4]
		if buf, err = encodeLV(false, uint16(2), uint16(2), msg.ServingPlmnRateControl); err != nil {
			err = nasError("encoding ServingPlmnRateControl [O TLV 4]", err)
			return
		}
		wire = append(append(wire, 0x1E), buf...)
	}

	if msg.EthernetHeaderCompressionConfiguration != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.EthernetHeaderCompressionConfiguration); err != nil {
			err = nasError("encoding EthernetHeaderCompressionConfiguration [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x1F), buf...)
	}

	if msg.ReceivedMbsContainer != nil {
		// TLV-E[9-65538]
		if buf, err = encodeLV(true, uint16(6), uint16(0), msg.ReceivedMbsContainer); err != nil {
			err = nasError("encoding ReceivedMbsContainer [O TLV-E 9-65538]", err)
			return
		}
		wire = append(append(wire, 0x71), buf...)
	}

	if msg.ServiceLevelAaContainer != nil {
		// TLV-E[6-n]
		if buf, err = encodeLV(true, uint16(3), uint16(0), msg.ServiceLevelAaContainer); err != nil {
			err = nasError("encoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
			return
		}
		wire = append(append(wire, 0x72), buf...)
	}

	msg.msgType = PduSessionModificationCommandMsgType //set message type to PDU SESSION MODIFICATION COMMAND
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *PduSessionModificationCommand) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding PduSessionModificationCommand", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
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
		case 0x2A: //TLV[8]
			offset++ //consume IEI
			v := &SessionAmbr{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(6), uint16(6), v); err != nil {
				err = nasError("decoding SessionAmbr [O TLV 8]", err)
				return
			}
			offset += consumed
			msg.SessionAmbr = v
		case 0x56: //TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding RqTimerValue [O TV 2]", ErrIncomplete)
				return
			}
			msg.RqTimerValue = new(Uint8)
			offset++ //consume IEI
			*msg.RqTimerValue = Uint8(wire[offset])
			offset++
		case 0x08: //TV[1]
			msg.AlwaysOnPduSessionIndication = new(Uint8)
			*msg.AlwaysOnPduSessionIndication = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x7A: //TLV-E[7-65538]
			offset++ //consume IEI
			v := &QosRules{}
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), v); err != nil {
				err = nasError("decoding AuthorizedQosRules [O TLV-E 7-65538]", err)
				return
			}
			offset += consumed
			msg.AuthorizedQosRules = v
		case 0x75: //TLV-E[7-65538]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), v); err != nil {
				err = nasError("decoding MappedEpsBearerContexts [O TLV-E 7-65538]", err)
				return
			}
			offset += consumed
			msg.MappedEpsBearerContexts = v
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
		case 0x77: //TLV-E[3-65538]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding AtsssContainer [O TLV-E 3-65538]", err)
				return
			}
			offset += consumed
			msg.AtsssContainer = v
		case 0x66: //TLV[5-257]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(255), v); err != nil {
				err = nasError("decoding IpHeaderCompressionConfiguration [O TLV 5-257]", err)
				return
			}
			offset += consumed
			msg.IpHeaderCompressionConfiguration = v
		case 0x74: //TLV-E[4-65538]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding PortManagementInformationContainer [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.PortManagementInformationContainer = v
		case 0x1E: //TLV[4]
			offset++ //consume IEI
			v := new(Uint16)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(2), v); err != nil {
				err = nasError("decoding ServingPlmnRateControl [O TLV 4]", err)
				return
			}
			offset += consumed
			msg.ServingPlmnRateControl = v
		case 0x1F: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding EthernetHeaderCompressionConfiguration [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.EthernetHeaderCompressionConfiguration = v
		case 0x71: //TLV-E[9-65538]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(6), uint16(0), v); err != nil {
				err = nasError("decoding ReceivedMbsContainer [O TLV-E 9-65538]", err)
				return
			}
			offset += consumed
			msg.ReceivedMbsContainer = v
		case 0x72: //TLV-E[6-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(3), uint16(0), v); err != nil {
				err = nasError("decoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
				return
			}
			offset += consumed
			msg.ServiceLevelAaContainer = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
