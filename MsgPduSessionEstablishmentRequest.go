/**generated time: 2024-07-17 15:11:00.945693**/

package nas

/*******************************************************
 * PDU SESSION ESTABLISHMENT REQUEST
 ******************************************************/
type PduSessionEstablishmentRequest struct {
	SmHeader
	IntegrityProtectionMaximumDataRate     IntegrityProtectionMaximumDataRate    //V [2]
	PduSessionType                         *Uint8                                //TV [9-][1]
	SscMode                                *Uint8                                //TV [A-][1]
	GsmCapability                          *Bytes                                //TLV [28][3-15]
	MaximumNumberOfSupportedPacketFilters  *Uint16                               //TV [55][3]
	AlwaysOnPduSessionRequested            *Uint8                                //TV [B-][1]
	SmPduDnRequestContainer                *Bytes                                //TLV [39][3-255]
	ExtendedProtocolConfigurationOptions   *ExtendedProtocolConfigurationOptions //TLV-E [7B][4-65538]
	IpHeaderCompressionConfiguration       *Bytes                                //TLV [66][5-257]
	DsTtEthernetPortMacAddress             *Bytes                                //TLV [6E][8]
	UeDsTtResidenceTime                    *Bytes                                //TLV [6F][10]
	PortManagementInformationContainer     *Bytes                                //TLV-E [74][8-65538]
	EthernetHeaderCompressionConfiguration *Uint8                                //TLV [1F][3]
	SuggestedInterfaceIdentifier           *PduAddress                           //TLV [29][11]
	ServiceLevelAaContainer                *Bytes                                //TLV-E [72][6-n]
	RequestedMbsContainer                  *Bytes                                //TLV-E [70][8-65538]
	PduSessionPairId                       *Uint8                                //TLV [34][3]
	Rsn                                    *Uint8                                //TLV [35][3]
}

func (msg *PduSessionEstablishmentRequest) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding PduSessionEstablishmentRequest", err)
		}
	}()
	var buf []byte
	// V[2]
	if buf, err = msg.IntegrityProtectionMaximumDataRate.encode(); err != nil {
		err = nasError("encoding IntegrityProtectionMaximumDataRate [M V 2]", err)
		return
	}
	if len(buf) != 2 {
		err = nasError("encoding IntegrityProtectionMaximumDataRate [M V 2]", ErrInvalidSize)
		return
	}
	wire = append(wire, buf...)

	if msg.PduSessionType != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x09<<4)|(uint8(*msg.PduSessionType)&0x0f))
	}

	if msg.SscMode != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0A<<4)|(uint8(*msg.SscMode)&0x0f))
	}

	if msg.GsmCapability != nil {
		// TLV[3-15]
		if buf, err = encodeLV(false, uint16(1), uint16(13), msg.GsmCapability); err != nil {
			err = nasError("encoding GsmCapability [O TLV 3-15]", err)
			return
		}
		wire = append(append(wire, 0x28), buf...)
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

	if msg.SmPduDnRequestContainer != nil {
		// TLV[3-255]
		if buf, err = encodeLV(false, uint16(1), uint16(253), msg.SmPduDnRequestContainer); err != nil {
			err = nasError("encoding SmPduDnRequestContainer [O TLV 3-255]", err)
			return
		}
		wire = append(append(wire, 0x39), buf...)
	}

	if msg.ExtendedProtocolConfigurationOptions != nil {
		// TLV-E[4-65538]
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.ExtendedProtocolConfigurationOptions); err != nil {
			err = nasError("encoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x7B), buf...)
	}

	if msg.IpHeaderCompressionConfiguration != nil {
		// TLV[5-257]
		if buf, err = encodeLV(false, uint16(3), uint16(255), msg.IpHeaderCompressionConfiguration); err != nil {
			err = nasError("encoding IpHeaderCompressionConfiguration [O TLV 5-257]", err)
			return
		}
		wire = append(append(wire, 0x66), buf...)
	}

	if msg.DsTtEthernetPortMacAddress != nil {
		// TLV[8]
		if buf, err = encodeLV(false, uint16(6), uint16(6), msg.DsTtEthernetPortMacAddress); err != nil {
			err = nasError("encoding DsTtEthernetPortMacAddress [O TLV 8]", err)
			return
		}
		wire = append(append(wire, 0x6E), buf...)
	}

	if msg.UeDsTtResidenceTime != nil {
		// TLV[10]
		if buf, err = encodeLV(false, uint16(8), uint16(8), msg.UeDsTtResidenceTime); err != nil {
			err = nasError("encoding UeDsTtResidenceTime [O TLV 10]", err)
			return
		}
		wire = append(append(wire, 0x6F), buf...)
	}

	if msg.PortManagementInformationContainer != nil {
		// TLV-E[8-65538]
		if buf, err = encodeLV(true, uint16(5), uint16(0), msg.PortManagementInformationContainer); err != nil {
			err = nasError("encoding PortManagementInformationContainer [O TLV-E 8-65538]", err)
			return
		}
		wire = append(append(wire, 0x74), buf...)
	}

	if msg.EthernetHeaderCompressionConfiguration != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.EthernetHeaderCompressionConfiguration); err != nil {
			err = nasError("encoding EthernetHeaderCompressionConfiguration [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x1F), buf...)
	}

	if msg.SuggestedInterfaceIdentifier != nil {
		// TLV[11]
		if buf, err = encodeLV(false, uint16(9), uint16(9), msg.SuggestedInterfaceIdentifier); err != nil {
			err = nasError("encoding SuggestedInterfaceIdentifier [O TLV 11]", err)
			return
		}
		wire = append(append(wire, 0x29), buf...)
	}

	if msg.ServiceLevelAaContainer != nil {
		// TLV-E[6-n]
		if buf, err = encodeLV(true, uint16(3), uint16(0), msg.ServiceLevelAaContainer); err != nil {
			err = nasError("encoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
			return
		}
		wire = append(append(wire, 0x72), buf...)
	}

	if msg.RequestedMbsContainer != nil {
		// TLV-E[8-65538]
		if buf, err = encodeLV(true, uint16(5), uint16(0), msg.RequestedMbsContainer); err != nil {
			err = nasError("encoding RequestedMbsContainer [O TLV-E 8-65538]", err)
			return
		}
		wire = append(append(wire, 0x70), buf...)
	}

	if msg.PduSessionPairId != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.PduSessionPairId); err != nil {
			err = nasError("encoding PduSessionPairId [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x34), buf...)
	}

	if msg.Rsn != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.Rsn); err != nil {
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
	// V[2]
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
		case 0x09: //TV[1]
			msg.PduSessionType = new(Uint8)
			*msg.PduSessionType = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x0A: //TV[1]
			msg.SscMode = new(Uint8)
			*msg.SscMode = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x28: //TLV[3-15]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(13), v); err != nil {
				err = nasError("decoding GsmCapability [O TLV 3-15]", err)
				return
			}
			offset += consumed
			msg.GsmCapability = v
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
		case 0x39: //TLV[3-255]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(253), v); err != nil {
				err = nasError("decoding SmPduDnRequestContainer [O TLV 3-255]", err)
				return
			}
			offset += consumed
			msg.SmPduDnRequestContainer = v
		case 0x7B: //TLV-E[4-65538]
			offset++ //consume IEI
			v := &ExtendedProtocolConfigurationOptions{}
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.ExtendedProtocolConfigurationOptions = v
		case 0x66: //TLV[5-257]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(255), v); err != nil {
				err = nasError("decoding IpHeaderCompressionConfiguration [O TLV 5-257]", err)
				return
			}
			offset += consumed
			msg.IpHeaderCompressionConfiguration = v
		case 0x6E: //TLV[8]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(6), uint16(6), v); err != nil {
				err = nasError("decoding DsTtEthernetPortMacAddress [O TLV 8]", err)
				return
			}
			offset += consumed
			msg.DsTtEthernetPortMacAddress = v
		case 0x6F: //TLV[10]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(8), uint16(8), v); err != nil {
				err = nasError("decoding UeDsTtResidenceTime [O TLV 10]", err)
				return
			}
			offset += consumed
			msg.UeDsTtResidenceTime = v
		case 0x74: //TLV-E[8-65538]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(5), uint16(0), v); err != nil {
				err = nasError("decoding PortManagementInformationContainer [O TLV-E 8-65538]", err)
				return
			}
			offset += consumed
			msg.PortManagementInformationContainer = v
		case 0x1F: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding EthernetHeaderCompressionConfiguration [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.EthernetHeaderCompressionConfiguration = v
		case 0x29: //TLV[11]
			offset++ //consume IEI
			v := &PduAddress{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(9), uint16(9), v); err != nil {
				err = nasError("decoding SuggestedInterfaceIdentifier [O TLV 11]", err)
				return
			}
			offset += consumed
			msg.SuggestedInterfaceIdentifier = v
		case 0x72: //TLV-E[6-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(3), uint16(0), v); err != nil {
				err = nasError("decoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
				return
			}
			offset += consumed
			msg.ServiceLevelAaContainer = v
		case 0x70: //TLV-E[8-65538]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(5), uint16(0), v); err != nil {
				err = nasError("decoding RequestedMbsContainer [O TLV-E 8-65538]", err)
				return
			}
			offset += consumed
			msg.RequestedMbsContainer = v
		case 0x34: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding PduSessionPairId [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.PduSessionPairId = v
		case 0x35: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding Rsn [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.Rsn = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
