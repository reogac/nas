/**generated time: 2024-12-16 16:36:18.698069**/

package nas

/*******************************************************
 * PDU SESSION MODIFICATION COMMAND
 ******************************************************/
type PduSessionModificationCommand struct {
	SmHeader
	GsmCause                               *uint8                                //O: TV [59][2]
	SessionAmbr                            *SessionAmbr                          //O: TLV [2A][8]
	RqTimerValue                           *uint8                                //O: TV [56][2]
	AlwaysOnPduSessionIndication           *uint8                                //O: TV [8-][1]
	AuthorizedQosRules                     *QosRules                             //O: TLV-E [7A][7-65538]
	MappedEpsBearerContexts                []byte                                //O: TLV-E [75][7-65538]
	AuthorizedQosFlowDescriptions          *QosFlowDescriptions                  //O: TLV-E [79][6-65538]
	ExtendedProtocolConfigurationOptions   *ExtendedProtocolConfigurationOptions //O: TLV-E [7B][4-65538]
	AtsssContainer                         []byte                                //O: TLV-E [77][3-65538]
	IpHeaderCompressionConfiguration       []byte                                //O: TLV [66][5-257]
	PortManagementInformationContainer     []byte                                //O: TLV-E [74][4-65538]
	ServingPlmnRateControl                 *uint16                               //O: TLV [1E][4]
	EthernetHeaderCompressionConfiguration *uint8                                //O: TLV [1F][3]
	ReceivedMbsContainer                   []byte                                //O: TLV-E [71][9-65538]
	ServiceLevelAaContainer                []byte                                //O: TLV-E [72][6-n]
}

func (msg *PduSessionModificationCommand) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding PduSessionModificationCommand", err)
		}
	}()
	var buf []byte
	// O: TV[2]
	if msg.GsmCause != nil {
		wire = append(wire, []byte{0x59, uint8(*msg.GsmCause)}...)
	}

	// O: TLV[8]
	if msg.SessionAmbr != nil {
		if buf, err = encodeLV(false, uint16(6), uint16(6), msg.SessionAmbr); err != nil {
			err = nasError("encoding SessionAmbr [O TLV 8]", err)
			return
		}
		wire = append(append(wire, 0x2A), buf...)
	}

	// O: TV[2]
	if msg.RqTimerValue != nil {
		wire = append(wire, []byte{0x56, uint8(*msg.RqTimerValue)}...)
	}

	// O: TV[1]
	if msg.AlwaysOnPduSessionIndication != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x08<<4)|(uint8(*msg.AlwaysOnPduSessionIndication)&0x0f))
	}

	// O: TLV-E[7-65538]
	if msg.AuthorizedQosRules != nil {
		if buf, err = encodeLV(true, uint16(4), uint16(0), msg.AuthorizedQosRules); err != nil {
			err = nasError("encoding AuthorizedQosRules [O TLV-E 7-65538]", err)
			return
		}
		wire = append(append(wire, 0x7A), buf...)
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

	// O: TLV-E[3-65538]
	if len(msg.AtsssContainer) > 0 {
		tmp := newBytesEncoder(msg.AtsssContainer)
		if buf, err = encodeLV(true, uint16(0), uint16(0), tmp); err != nil {
			err = nasError("encoding AtsssContainer [O TLV-E 3-65538]", err)
			return
		}
		wire = append(append(wire, 0x77), buf...)
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

	// O: TLV-E[4-65538]
	if len(msg.PortManagementInformationContainer) > 0 {
		tmp := newBytesEncoder(msg.PortManagementInformationContainer)
		if buf, err = encodeLV(true, uint16(1), uint16(0), tmp); err != nil {
			err = nasError("encoding PortManagementInformationContainer [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x74), buf...)
	}

	// O: TLV[4]
	if msg.ServingPlmnRateControl != nil {
		tmp := newUint16Encoder(*msg.ServingPlmnRateControl)
		if buf, err = encodeLV(false, uint16(2), uint16(2), tmp); err != nil {
			err = nasError("encoding ServingPlmnRateControl [O TLV 4]", err)
			return
		}
		wire = append(append(wire, 0x1E), buf...)
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

	// O: TLV-E[9-65538]
	if len(msg.ReceivedMbsContainer) > 0 {
		tmp := newBytesEncoder(msg.ReceivedMbsContainer)
		if buf, err = encodeLV(true, uint16(6), uint16(0), tmp); err != nil {
			err = nasError("encoding ReceivedMbsContainer [O TLV-E 9-65538]", err)
			return
		}
		wire = append(append(wire, 0x71), buf...)
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
		case 0x59: //O: TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding GsmCause [O TV 2]", ErrIncomplete)
				return
			}
			msg.GsmCause = new(uint8)
			offset++ //consume IEI
			*msg.GsmCause = wire[offset]
			offset++
		case 0x2A: //O: TLV[8]
			offset++ //consume IEI
			v := new(SessionAmbr)
			if consumed, err = decodeLV(wire[offset:], false, uint16(6), uint16(6), v); err != nil {
				err = nasError("decoding SessionAmbr [O TLV 8]", err)
				return
			}
			offset += consumed
			msg.SessionAmbr = v
		case 0x56: //O: TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding RqTimerValue [O TV 2]", ErrIncomplete)
				return
			}
			msg.RqTimerValue = new(uint8)
			offset++ //consume IEI
			*msg.RqTimerValue = wire[offset]
			offset++
		case 0x08: //O: TV[1]
			msg.AlwaysOnPduSessionIndication = new(uint8)
			*msg.AlwaysOnPduSessionIndication = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x7A: //O: TLV-E[7-65538]
			offset++ //consume IEI
			v := new(QosRules)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), v); err != nil {
				err = nasError("decoding AuthorizedQosRules [O TLV-E 7-65538]", err)
				return
			}
			offset += consumed
			msg.AuthorizedQosRules = v
		case 0x75: //O: TLV-E[7-65538]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), v); err != nil {
				err = nasError("decoding MappedEpsBearerContexts [O TLV-E 7-65538]", err)
				return
			}
			offset += consumed
			msg.MappedEpsBearerContexts = []byte(*v)
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
		case 0x77: //O: TLV-E[3-65538]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding AtsssContainer [O TLV-E 3-65538]", err)
				return
			}
			offset += consumed
			msg.AtsssContainer = []byte(*v)
		case 0x66: //O: TLV[5-257]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(255), v); err != nil {
				err = nasError("decoding IpHeaderCompressionConfiguration [O TLV 5-257]", err)
				return
			}
			offset += consumed
			msg.IpHeaderCompressionConfiguration = []byte(*v)
		case 0x74: //O: TLV-E[4-65538]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding PortManagementInformationContainer [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.PortManagementInformationContainer = []byte(*v)
		case 0x1E: //O: TLV[4]
			offset++ //consume IEI
			v := new(uint16Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(2), v); err != nil {
				err = nasError("decoding ServingPlmnRateControl [O TLV 4]", err)
				return
			}
			offset += consumed
			msg.ServingPlmnRateControl = (*uint16)(v)
		case 0x1F: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding EthernetHeaderCompressionConfiguration [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.EthernetHeaderCompressionConfiguration = (*uint8)(v)
		case 0x71: //O: TLV-E[9-65538]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(6), uint16(0), v); err != nil {
				err = nasError("decoding ReceivedMbsContainer [O TLV-E 9-65538]", err)
				return
			}
			offset += consumed
			msg.ReceivedMbsContainer = []byte(*v)
		case 0x72: //O: TLV-E[6-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(3), uint16(0), v); err != nil {
				err = nasError("decoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
				return
			}
			offset += consumed
			msg.ServiceLevelAaContainer = []byte(*v)
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
