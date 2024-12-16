/** this file was generated at 2024-12-16 17:55:27.330329 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * PDU SESSION MODIFICATION REQUEST
 ******************************************************/
type PduSessionModificationRequest struct {
	SmHeader
	GsmCapability                          []byte                                //O: TLV [28][3-15]
	GsmCause                               *uint8                                //O: TV [59][2]
	MaximumNumberOfSupportedPacketFilters  *uint16                               //O: TV [55][3]
	AlwaysOnPduSessionRequested            *uint8                                //O: TV [B-][1]
	IntegrityProtectionMaximumDataRate     *IntegrityProtectionMaximumDataRate   //O: TV [13][3]
	RequestedQosRules                      *QosRules                             //O: TLV-E [7A][7-65538]
	RequestedQosFlowDescriptions           *QosFlowDescriptions                  //O: TLV-E [79][6-65538]
	MappedEpsBearerContexts                []byte                                //O: TLV-E [75][7-65538]
	ExtendedProtocolConfigurationOptions   *ExtendedProtocolConfigurationOptions //O: TLV-E [7B][4-65538]
	PortManagementInformationContainer     []byte                                //O: TLV-E [74][4-65538]
	IpHeaderCompressionConfiguration       []byte                                //O: TLV [66][5-257]
	EthernetHeaderCompressionConfiguration *uint8                                //O: TLV [1F][3]
	RequestedMbsContainer                  []byte                                //O: TLV-E [70][8-65538]
	ServiceLevelAaContainer                []byte                                //O: TLV-E [72][6-n]
}

func (msg *PduSessionModificationRequest) encode() (wire []byte, err error) {
	var buf []byte
	// O: TLV[3-15]
	if len(msg.GsmCapability) > 0 {
		tmp := newBytesEncoder(msg.GsmCapability)
		if buf, err = encodeLV(false, uint16(1), uint16(13), tmp); err != nil {
			err = nasError("encoding GsmCapability [O TLV 3-15]", err)
			return
		}
		wire = append(append(wire, 0x28), buf...)
	}

	// O: TV[2]
	if msg.GsmCause != nil {
		wire = append(wire, []byte{0x59, uint8(*msg.GsmCause)}...)
	}

	// O: TV[3]
	if msg.MaximumNumberOfSupportedPacketFilters != nil {
		tmp := newUint16Encoder(*msg.MaximumNumberOfSupportedPacketFilters)
		if buf, err = tmp.encode(); err != nil {
			err = nasError("encoding MaximumNumberOfSupportedPacketFilters [O TV 3]", err)
			return
		}
		if len(buf) != 2 {
			err = nasError("encoding MaximumNumberOfSupportedPacketFilters [O TV 3]", ErrInvalidSize)
			return
		}
		wire = append(append(wire, 0x55), buf...)
	}

	// O: TV[1]
	if msg.AlwaysOnPduSessionRequested != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0B<<4)|(uint8(*msg.AlwaysOnPduSessionRequested)&0x0f))
	}

	// O: TV[3]
	if msg.IntegrityProtectionMaximumDataRate != nil {
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

	// O: TLV-E[7-65538]
	if msg.RequestedQosRules != nil {
		if buf, err = encodeLV(true, uint16(4), uint16(0), msg.RequestedQosRules); err != nil {
			err = nasError("encoding RequestedQosRules [O TLV-E 7-65538]", err)
			return
		}
		wire = append(append(wire, 0x7A), buf...)
	}

	// O: TLV-E[6-65538]
	if msg.RequestedQosFlowDescriptions != nil {
		if buf, err = encodeLV(true, uint16(3), uint16(0), msg.RequestedQosFlowDescriptions); err != nil {
			err = nasError("encoding RequestedQosFlowDescriptions [O TLV-E 6-65538]", err)
			return
		}
		wire = append(append(wire, 0x79), buf...)
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

	// O: TLV-E[4-65538]
	if msg.ExtendedProtocolConfigurationOptions != nil {
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.ExtendedProtocolConfigurationOptions); err != nil {
			err = nasError("encoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x7B), buf...)
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

	// O: TLV-E[8-65538]
	if len(msg.RequestedMbsContainer) > 0 {
		tmp := newBytesEncoder(msg.RequestedMbsContainer)
		if buf, err = encodeLV(true, uint16(5), uint16(0), tmp); err != nil {
			err = nasError("encoding RequestedMbsContainer [O TLV-E 8-65538]", err)
			return
		}
		wire = append(append(wire, 0x70), buf...)
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

	msg.msgType = PduSessionModificationRequestMsgType //set message type to PDU SESSION MODIFICATION REQUEST
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *PduSessionModificationRequest) decodeBody(wire []byte) (err error) {
	offset := 0
	wireLen := len(wire)
	consumed := 0
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x28: //O: TLV[3-15]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(13), v); err != nil {
				err = nasError("decoding GsmCapability [O TLV 3-15]", err)
				return
			}
			offset += consumed
			msg.GsmCapability = []byte(*v)
		case 0x59: //O: TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding GsmCause [O TV 2]", ErrIncomplete)
				return
			}
			msg.GsmCause = new(uint8)
			offset++ //consume IEI
			*msg.GsmCause = wire[offset]
			offset++
		case 0x55: //O: TV[3]
			if offset+3 > wireLen {
				err = nasError("decoding MaximumNumberOfSupportedPacketFilters [O TV 3]", ErrIncomplete)
				return
			}
			offset++ //consume IEI
			v := new(uint16Decoder)
			if err = v.decode(wire[offset : offset+2]); err != nil {
				err = nasError("decoding MaximumNumberOfSupportedPacketFilters [O TV 3]", err)
				return
			}
			msg.MaximumNumberOfSupportedPacketFilters = (*uint16)(v)
			offset += 2

		case 0x0B: //O: TV[1]
			msg.AlwaysOnPduSessionRequested = new(uint8)
			*msg.AlwaysOnPduSessionRequested = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x13: //O: TV[3]
			if offset+3 > wireLen {
				err = nasError("decoding IntegrityProtectionMaximumDataRate [O TV 3]", ErrIncomplete)
				return
			}
			offset++ //consume IEI
			v := new(IntegrityProtectionMaximumDataRate)
			if err = v.decode(wire[offset : offset+2]); err != nil {
				err = nasError("decoding IntegrityProtectionMaximumDataRate [O TV 3]", err)
				return
			}
			msg.IntegrityProtectionMaximumDataRate = v
			offset += 2

		case 0x7A: //O: TLV-E[7-65538]
			offset++ //consume IEI
			v := new(QosRules)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), v); err != nil {
				err = nasError("decoding RequestedQosRules [O TLV-E 7-65538]", err)
				return
			}
			offset += consumed
			msg.RequestedQosRules = v
		case 0x79: //O: TLV-E[6-65538]
			offset++ //consume IEI
			v := new(QosFlowDescriptions)
			if consumed, err = decodeLV(wire[offset:], true, uint16(3), uint16(0), v); err != nil {
				err = nasError("decoding RequestedQosFlowDescriptions [O TLV-E 6-65538]", err)
				return
			}
			offset += consumed
			msg.RequestedQosFlowDescriptions = v
		case 0x75: //O: TLV-E[7-65538]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), v); err != nil {
				err = nasError("decoding MappedEpsBearerContexts [O TLV-E 7-65538]", err)
				return
			}
			offset += consumed
			msg.MappedEpsBearerContexts = []byte(*v)
		case 0x7B: //O: TLV-E[4-65538]
			offset++ //consume IEI
			v := new(ExtendedProtocolConfigurationOptions)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.ExtendedProtocolConfigurationOptions = v
		case 0x74: //O: TLV-E[4-65538]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding PortManagementInformationContainer [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.PortManagementInformationContainer = []byte(*v)
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
		case 0x70: //O: TLV-E[8-65538]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(5), uint16(0), v); err != nil {
				err = nasError("decoding RequestedMbsContainer [O TLV-E 8-65538]", err)
				return
			}
			offset += consumed
			msg.RequestedMbsContainer = []byte(*v)
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
