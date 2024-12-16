/**generated time: 2024-12-16 16:36:18.696232**/

package nas

/*******************************************************
 * PDU SESSION ESTABLISHMENT REQUEST
 ******************************************************/
type PduSessionEstablishmentRequest struct {
	SmHeader
	IntegrityProtectionMaximumDataRate     IntegrityProtectionMaximumDataRate    //M: V [2]
	PduSessionType                         *uint8                                //O: TV [9-][1]
	SscMode                                *uint8                                //O: TV [A-][1]
	GsmCapability                          []byte                                //O: TLV [28][3-15]
	MaximumNumberOfSupportedPacketFilters  *uint16                               //O: TV [55][3]
	AlwaysOnPduSessionRequested            *uint8                                //O: TV [B-][1]
	SmPduDnRequestContainer                []byte                                //O: TLV [39][3-255]
	ExtendedProtocolConfigurationOptions   *ExtendedProtocolConfigurationOptions //O: TLV-E [7B][4-65538]
	IpHeaderCompressionConfiguration       []byte                                //O: TLV [66][5-257]
	DsTtEthernetPortMacAddress             []byte                                //O: TLV [6E][8]
	UeDsTtResidenceTime                    []byte                                //O: TLV [6F][10]
	PortManagementInformationContainer     []byte                                //O: TLV-E [74][8-65538]
	EthernetHeaderCompressionConfiguration *uint8                                //O: TLV [1F][3]
	SuggestedInterfaceIdentifier           *PduAddress                           //O: TLV [29][11]
	ServiceLevelAaContainer                []byte                                //O: TLV-E [72][6-n]
	RequestedMbsContainer                  []byte                                //O: TLV-E [70][8-65538]
	PduSessionPairId                       *uint8                                //O: TLV [34][3]
	Rsn                                    *uint8                                //O: TLV [35][3]
}

func (msg *PduSessionEstablishmentRequest) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding PduSessionEstablishmentRequest", err)
		}
	}()
	var buf []byte
	// M: V[2]
	if buf, err = msg.IntegrityProtectionMaximumDataRate.encode(); err != nil {
		err = nasError("encoding IntegrityProtectionMaximumDataRate [M V 2]", err)
		return
	}
	if len(buf) != 2 {
		err = nasError("encoding IntegrityProtectionMaximumDataRate [M V 2]", ErrInvalidSize)
		return
	}
	wire = append(wire, buf...)

	// O: TV[1]
	if msg.PduSessionType != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x09<<4)|(uint8(*msg.PduSessionType)&0x0f))
	}

	// O: TV[1]
	if msg.SscMode != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0A<<4)|(uint8(*msg.SscMode)&0x0f))
	}

	// O: TLV[3-15]
	if len(msg.GsmCapability) > 0 {
		tmp := newBytesEncoder(msg.GsmCapability)
		if buf, err = encodeLV(false, uint16(1), uint16(13), tmp); err != nil {
			err = nasError("encoding GsmCapability [O TLV 3-15]", err)
			return
		}
		wire = append(append(wire, 0x28), buf...)
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

	// O: TLV[3-255]
	if len(msg.SmPduDnRequestContainer) > 0 {
		tmp := newBytesEncoder(msg.SmPduDnRequestContainer)
		if buf, err = encodeLV(false, uint16(1), uint16(253), tmp); err != nil {
			err = nasError("encoding SmPduDnRequestContainer [O TLV 3-255]", err)
			return
		}
		wire = append(append(wire, 0x39), buf...)
	}

	// O: TLV-E[4-65538]
	if msg.ExtendedProtocolConfigurationOptions != nil {
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.ExtendedProtocolConfigurationOptions); err != nil {
			err = nasError("encoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x7B), buf...)
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

	// O: TLV[8]
	if len(msg.DsTtEthernetPortMacAddress) > 0 {
		tmp := newBytesEncoder(msg.DsTtEthernetPortMacAddress)
		if buf, err = encodeLV(false, uint16(6), uint16(6), tmp); err != nil {
			err = nasError("encoding DsTtEthernetPortMacAddress [O TLV 8]", err)
			return
		}
		wire = append(append(wire, 0x6E), buf...)
	}

	// O: TLV[10]
	if len(msg.UeDsTtResidenceTime) > 0 {
		tmp := newBytesEncoder(msg.UeDsTtResidenceTime)
		if buf, err = encodeLV(false, uint16(8), uint16(8), tmp); err != nil {
			err = nasError("encoding UeDsTtResidenceTime [O TLV 10]", err)
			return
		}
		wire = append(append(wire, 0x6F), buf...)
	}

	// O: TLV-E[8-65538]
	if len(msg.PortManagementInformationContainer) > 0 {
		tmp := newBytesEncoder(msg.PortManagementInformationContainer)
		if buf, err = encodeLV(true, uint16(5), uint16(0), tmp); err != nil {
			err = nasError("encoding PortManagementInformationContainer [O TLV-E 8-65538]", err)
			return
		}
		wire = append(append(wire, 0x74), buf...)
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

	// O: TLV[11]
	if msg.SuggestedInterfaceIdentifier != nil {
		if buf, err = encodeLV(false, uint16(9), uint16(9), msg.SuggestedInterfaceIdentifier); err != nil {
			err = nasError("encoding SuggestedInterfaceIdentifier [O TLV 11]", err)
			return
		}
		wire = append(append(wire, 0x29), buf...)
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

	// O: TLV-E[8-65538]
	if len(msg.RequestedMbsContainer) > 0 {
		tmp := newBytesEncoder(msg.RequestedMbsContainer)
		if buf, err = encodeLV(true, uint16(5), uint16(0), tmp); err != nil {
			err = nasError("encoding RequestedMbsContainer [O TLV-E 8-65538]", err)
			return
		}
		wire = append(append(wire, 0x70), buf...)
	}

	// O: TLV[3]
	if msg.PduSessionPairId != nil {
		tmp := newUint8Encoder(*msg.PduSessionPairId)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding PduSessionPairId [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x34), buf...)
	}

	// O: TLV[3]
	if msg.Rsn != nil {
		tmp := newUint8Encoder(*msg.Rsn)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding Rsn [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x35), buf...)
	}

	msg.msgType = PduSessionEstablishmentRequestMsgType //set message type to PDU SESSION ESTABLISHMENT REQUEST
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *PduSessionEstablishmentRequest) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding PduSessionEstablishmentRequest", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// M V[2]
	if offset+2 > wireLen {
		err = nasError("decoding IntegrityProtectionMaximumDataRate [M V 2]", ErrIncomplete)
		return
	}
	if err = msg.IntegrityProtectionMaximumDataRate.decode(wire[offset : offset+2]); err != nil {
		err = nasError("decoding IntegrityProtectionMaximumDataRate [M V 2]", err)
		return
	}
	offset += 2

	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x09: //O: TV[1]
			msg.PduSessionType = new(uint8)
			*msg.PduSessionType = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x0A: //O: TV[1]
			msg.SscMode = new(uint8)
			*msg.SscMode = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x28: //O: TLV[3-15]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(13), v); err != nil {
				err = nasError("decoding GsmCapability [O TLV 3-15]", err)
				return
			}
			offset += consumed
			msg.GsmCapability = []byte(*v)
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
		case 0x39: //O: TLV[3-255]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(253), v); err != nil {
				err = nasError("decoding SmPduDnRequestContainer [O TLV 3-255]", err)
				return
			}
			offset += consumed
			msg.SmPduDnRequestContainer = []byte(*v)
		case 0x7B: //O: TLV-E[4-65538]
			offset++ //consume IEI
			v := new(ExtendedProtocolConfigurationOptions)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.ExtendedProtocolConfigurationOptions = v
		case 0x66: //O: TLV[5-257]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(255), v); err != nil {
				err = nasError("decoding IpHeaderCompressionConfiguration [O TLV 5-257]", err)
				return
			}
			offset += consumed
			msg.IpHeaderCompressionConfiguration = []byte(*v)
		case 0x6E: //O: TLV[8]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(6), uint16(6), v); err != nil {
				err = nasError("decoding DsTtEthernetPortMacAddress [O TLV 8]", err)
				return
			}
			offset += consumed
			msg.DsTtEthernetPortMacAddress = []byte(*v)
		case 0x6F: //O: TLV[10]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(8), uint16(8), v); err != nil {
				err = nasError("decoding UeDsTtResidenceTime [O TLV 10]", err)
				return
			}
			offset += consumed
			msg.UeDsTtResidenceTime = []byte(*v)
		case 0x74: //O: TLV-E[8-65538]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(5), uint16(0), v); err != nil {
				err = nasError("decoding PortManagementInformationContainer [O TLV-E 8-65538]", err)
				return
			}
			offset += consumed
			msg.PortManagementInformationContainer = []byte(*v)
		case 0x1F: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding EthernetHeaderCompressionConfiguration [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.EthernetHeaderCompressionConfiguration = (*uint8)(v)
		case 0x29: //O: TLV[11]
			offset++ //consume IEI
			v := new(PduAddress)
			if consumed, err = decodeLV(wire[offset:], false, uint16(9), uint16(9), v); err != nil {
				err = nasError("decoding SuggestedInterfaceIdentifier [O TLV 11]", err)
				return
			}
			offset += consumed
			msg.SuggestedInterfaceIdentifier = v
		case 0x72: //O: TLV-E[6-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(3), uint16(0), v); err != nil {
				err = nasError("decoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
				return
			}
			offset += consumed
			msg.ServiceLevelAaContainer = []byte(*v)
		case 0x70: //O: TLV-E[8-65538]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(5), uint16(0), v); err != nil {
				err = nasError("decoding RequestedMbsContainer [O TLV-E 8-65538]", err)
				return
			}
			offset += consumed
			msg.RequestedMbsContainer = []byte(*v)
		case 0x34: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding PduSessionPairId [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.PduSessionPairId = (*uint8)(v)
		case 0x35: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding Rsn [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.Rsn = (*uint8)(v)
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
