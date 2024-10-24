/**generated time: 2024-07-17 15:11:00.947611**/

package nas

/*******************************************************
 * PDU SESSION MODIFICATION REQUEST
 ******************************************************/
type PduSessionModificationRequest struct {
	SmHeader
	GsmCapability                          *Bytes                                //TLV [28][3-15]
	GsmCause                               *Uint8                                //TV [59][2]
	MaximumNumberOfSupportedPacketFilters  *Uint16                               //TV [55][3]
	AlwaysOnPduSessionRequested            *Uint8                                //TV [B-][1]
	IntegrityProtectionMaximumDataRate     *IntegrityProtectionMaximumDataRate   //TV [13][3]
	RequestedQosRules                      *QosRules                             //TLV-E [7A][7-65538]
	RequestedQosFlowDescriptions           *QosFlowDescriptions                  //TLV-E [79][6-65538]
	MappedEpsBearerContexts                *Bytes                                //TLV-E [75][7-65538]
	ExtendedProtocolConfigurationOptions   *ExtendedProtocolConfigurationOptions //TLV-E [7B][4-65538]
	PortManagementInformationContainer     *Bytes                                //TLV-E [74][4-65538]
	IpHeaderCompressionConfiguration       *Bytes                                //TLV [66][5-257]
	EthernetHeaderCompressionConfiguration *Uint8                                //TLV [1F][3]
	RequestedMbsContainer                  *Bytes                                //TLV-E [70][8-65538]
	ServiceLevelAaContainer                *Bytes                                //TLV-E [72][6-n]
}

func (msg *PduSessionModificationRequest) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding PduSessionModificationRequest", err)
		}
	}()
	var buf []byte
	if msg.GsmCapability != nil {
		// TLV[3-15]
		if buf, err = encodeLV(false, uint16(1), uint16(13), msg.GsmCapability); err != nil {
			err = nasError("encoding GsmCapability [O TLV 3-15]", err)
			return
		}
		wire = append(append(wire, 0x28), buf...)
	}

	if msg.GsmCause != nil {
		//TV[2]
		wire = append(wire, []byte{0x59, uint8(*msg.GsmCause)}...)
	}

	if msg.MaximumNumberOfSupportedPacketFilters != nil {
		// TV[3]
		if buf, err = msg.MaximumNumberOfSupportedPacketFilters.encode(); err != nil {
			err = nasError("encoding MaximumNumberOfSupportedPacketFilters [O TV 3]", err)
			return
		}
		if len(buf) != 2 {
			err = nasError("encoding MaximumNumberOfSupportedPacketFilters [O TV 3]", ErrInvalidSize)
			return
		}
		wire = append(append(wire, 0x55), buf...)
	}

	if msg.AlwaysOnPduSessionRequested != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0B<<4)|(uint8(*msg.AlwaysOnPduSessionRequested)&0x0f))
	}

	if msg.IntegrityProtectionMaximumDataRate != nil {
		// TV[3]
		if buf, err = msg.IntegrityProtectionMaximumDataRate.encode(); err != nil {
			err = nasError("encoding IntegrityProtectionMaximumDataRate [O TV 3]", err)
			return
		}
		if len(buf) != 2 {
			err = nasError("encoding IntegrityProtectionMaximumDataRate [O TV 3]", ErrInvalidSize)
			return
		}
		wire = append(append(wire, 0x13), buf...)
	}

	if msg.RequestedQosRules != nil {
		// TLV-E[7-65538]
		if buf, err = encodeLV(true, uint16(4), uint16(0), msg.RequestedQosRules); err != nil {
			err = nasError("encoding RequestedQosRules [O TLV-E 7-65538]", err)
			return
		}
		wire = append(append(wire, 0x7A), buf...)
	}

	if msg.RequestedQosFlowDescriptions != nil {
		// TLV-E[6-65538]
		if buf, err = encodeLV(true, uint16(3), uint16(0), msg.RequestedQosFlowDescriptions); err != nil {
			err = nasError("encoding RequestedQosFlowDescriptions [O TLV-E 6-65538]", err)
			return
		}
		wire = append(append(wire, 0x79), buf...)
	}

	if msg.MappedEpsBearerContexts != nil {
		// TLV-E[7-65538]
		if buf, err = encodeLV(true, uint16(4), uint16(0), msg.MappedEpsBearerContexts); err != nil {
			err = nasError("encoding MappedEpsBearerContexts [O TLV-E 7-65538]", err)
			return
		}
		wire = append(append(wire, 0x75), buf...)
	}

	if msg.ExtendedProtocolConfigurationOptions != nil {
		// TLV-E[4-65538]
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.ExtendedProtocolConfigurationOptions); err != nil {
			err = nasError("encoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x7B), buf...)
	}

	if msg.PortManagementInformationContainer != nil {
		// TLV-E[4-65538]
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.PortManagementInformationContainer); err != nil {
			err = nasError("encoding PortManagementInformationContainer [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x74), buf...)
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

	if msg.RequestedMbsContainer != nil {
		// TLV-E[8-65538]
		if buf, err = encodeLV(true, uint16(5), uint16(0), msg.RequestedMbsContainer); err != nil {
			err = nasError("encoding RequestedMbsContainer [O TLV-E 8-65538]", err)
			return
		}
		wire = append(append(wire, 0x70), buf...)
	}

	if msg.ServiceLevelAaContainer != nil {
		// TLV-E[6-n]
		if buf, err = encodeLV(true, uint16(3), uint16(0), msg.ServiceLevelAaContainer); err != nil {
			err = nasError("encoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
			return
		}
		wire = append(append(wire, 0x72), buf...)
	}

	msg.msgType = PduSessionModificationRequestMsgType //set message type to PDU SESSION MODIFICATION REQUEST
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *PduSessionModificationRequest) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding PduSessionModificationRequest", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x28: //TLV[3-15]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(13), v); err != nil {
				err = nasError("decoding GsmCapability [O TLV 3-15]", err)
				return
			}
			offset += consumed
			msg.GsmCapability = v
		case 0x59: //TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding GsmCause [O TV 2]", ErrIncomplete)
				return
			}
			msg.GsmCause = new(Uint8)
			offset++ //consume IEI
			*msg.GsmCause = Uint8(wire[offset])
			offset++
		case 0x55: //TV[3]
			if offset+3 > wireLen {
				err = nasError("decoding MaximumNumberOfSupportedPacketFilters [O TV 3]", ErrIncomplete)
				return
			}
			offset++ //consume IEI
			v := new(Uint16)
			if err = v.decode(wire[offset : offset+2]); err != nil {
				err = nasError("decoding MaximumNumberOfSupportedPacketFilters [O TV 3]", err)
				return
			}
			msg.MaximumNumberOfSupportedPacketFilters = v
			offset += 2

		case 0x0B: //TV[1]
			msg.AlwaysOnPduSessionRequested = new(Uint8)
			*msg.AlwaysOnPduSessionRequested = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x13: //TV[3]
			if offset+3 > wireLen {
				err = nasError("decoding IntegrityProtectionMaximumDataRate [O TV 3]", ErrIncomplete)
				return
			}
			offset++ //consume IEI
			v := &IntegrityProtectionMaximumDataRate{}
			if err = v.decode(wire[offset : offset+2]); err != nil {
				err = nasError("decoding IntegrityProtectionMaximumDataRate [O TV 3]", err)
				return
			}
			msg.IntegrityProtectionMaximumDataRate = v
			offset += 2

		case 0x7A: //TLV-E[7-65538]
			offset++ //consume IEI
			v := &QosRules{}
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), v); err != nil {
				err = nasError("decoding RequestedQosRules [O TLV-E 7-65538]", err)
				return
			}
			offset += consumed
			msg.RequestedQosRules = v
		case 0x79: //TLV-E[6-65538]
			offset++ //consume IEI
			v := &QosFlowDescriptions{}
			if consumed, err = decodeLV(wire[offset:], true, uint16(3), uint16(0), v); err != nil {
				err = nasError("decoding RequestedQosFlowDescriptions [O TLV-E 6-65538]", err)
				return
			}
			offset += consumed
			msg.RequestedQosFlowDescriptions = v
		case 0x75: //TLV-E[7-65538]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), v); err != nil {
				err = nasError("decoding MappedEpsBearerContexts [O TLV-E 7-65538]", err)
				return
			}
			offset += consumed
			msg.MappedEpsBearerContexts = v
		case 0x7B: //TLV-E[4-65538]
			offset++ //consume IEI
			v := &ExtendedProtocolConfigurationOptions{}
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.ExtendedProtocolConfigurationOptions = v
		case 0x74: //TLV-E[4-65538]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding PortManagementInformationContainer [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.PortManagementInformationContainer = v
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
		case 0x70: //TLV-E[8-65538]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(5), uint16(0), v); err != nil {
				err = nasError("decoding RequestedMbsContainer [O TLV-E 8-65538]", err)
				return
			}
			offset += consumed
			msg.RequestedMbsContainer = v
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
