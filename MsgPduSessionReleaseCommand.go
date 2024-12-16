/**generated time: 2024-12-16 16:36:18.698721**/

package nas

/*******************************************************
 * PDU SESSION RELEASE COMMAND
 ******************************************************/
type PduSessionReleaseCommand struct {
	SmHeader
	GsmCause                             uint8                                 //M: V [1]
	BackOffTimerValue                    *GprsTimer3                           //O: TLV [37][3]
	EapMessage                           []byte                                //O: TLV-E [78][7-1503]
	GsmCongestionReAttemptIndicator      *uint8                                //O: TLV [61][3]
	ExtendedProtocolConfigurationOptions *ExtendedProtocolConfigurationOptions //O: TLV-E [7B][4-65538]
	AccessType                           *uint8                                //O: TV [D-][1]
	ServiceLevelAaContainer              []byte                                //O: TLV-E [72][6-n]
}

func (msg *PduSessionReleaseCommand) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding PduSessionReleaseCommand", err)
		}
	}()
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

	// O: TV[1]
	if msg.AccessType != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0D<<4)|(uint8(*msg.AccessType)&0x0f))
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

	msg.msgType = PduSessionReleaseCommandMsgType //set message type to PDU SESSION RELEASE COMMAND
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *PduSessionReleaseCommand) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding PduSessionReleaseCommand", err)
		}
	}()
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
		case 0x78: //O: TLV-E[7-1503]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), v); err != nil {
				err = nasError("decoding EapMessage [O TLV-E 7-1503]", err)
				return
			}
			offset += consumed
			msg.EapMessage = []byte(*v)
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
		case 0x0D: //O: TV[1]
			msg.AccessType = new(uint8)
			*msg.AccessType = wire[offset] & 0x0f //take right 1/2
			offset++
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
