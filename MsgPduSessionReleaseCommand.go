/**generated time: 2024-07-17 15:11:00.949182**/

package nas

/*******************************************************
 * PDU SESSION RELEASE COMMAND
 ******************************************************/
type PduSessionReleaseCommand struct {
	SmHeader
	GsmCause                             Uint8                                 //V [1]
	BackOffTimerValue                    *GprsTimer3                           //TLV [37][3]
	EapMessage                           *Bytes                                //TLV-E [78][7-1503]
	GsmCongestionReAttemptIndicator      *Uint8                                //TLV [61][3]
	ExtendedProtocolConfigurationOptions *ExtendedProtocolConfigurationOptions //TLV-E [7B][4-65538]
	AccessType                           *Uint8                                //TV [D-][1]
	ServiceLevelAaContainer              *Bytes                                //TLV-E [72][6-n]
}

func (msg *PduSessionReleaseCommand) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding PduSessionReleaseCommand", err)
		}
	}()
	var buf []byte
	// V[1]
	wire = append(wire, uint8(msg.GsmCause))

	if msg.BackOffTimerValue != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.BackOffTimerValue); err != nil {
			err = nasError("encoding BackOffTimerValue [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x37), buf...)
	}

	if msg.EapMessage != nil {
		// TLV-E[7-1503]
		if buf, err = encodeLV(true, uint16(4), uint16(1500), msg.EapMessage); err != nil {
			err = nasError("encoding EapMessage [O TLV-E 7-1503]", err)
			return
		}
		wire = append(append(wire, 0x78), buf...)
	}

	if msg.GsmCongestionReAttemptIndicator != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.GsmCongestionReAttemptIndicator); err != nil {
			err = nasError("encoding GsmCongestionReAttemptIndicator [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x61), buf...)
	}

	if msg.ExtendedProtocolConfigurationOptions != nil {
		// TLV-E[4-65538]
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.ExtendedProtocolConfigurationOptions); err != nil {
			err = nasError("encoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x7B), buf...)
	}

	if msg.AccessType != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0D<<4)|(uint8(*msg.AccessType)&0x0f))
	}

	if msg.ServiceLevelAaContainer != nil {
		// TLV-E[6-n]
		if buf, err = encodeLV(true, uint16(3), uint16(0), msg.ServiceLevelAaContainer); err != nil {
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
	// V[1]
	if offset+1 > wireLen {
		err = nasError("decoding GsmCause [M V 1]", ErrIncomplete)
		return
	}
	msg.GsmCause = Uint8(wire[offset])
	offset++

	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x37: //TLV[3]
			offset++ //consume IEI
			v := &GprsTimer3{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding BackOffTimerValue [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.BackOffTimerValue = v
		case 0x78: //TLV-E[7-1503]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), v); err != nil {
				err = nasError("decoding EapMessage [O TLV-E 7-1503]", err)
				return
			}
			offset += consumed
			msg.EapMessage = v
		case 0x61: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding GsmCongestionReAttemptIndicator [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.GsmCongestionReAttemptIndicator = v
		case 0x7B: //TLV-E[4-65538]
			offset++ //consume IEI
			v := &ExtendedProtocolConfigurationOptions{}
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.ExtendedProtocolConfigurationOptions = v
		case 0x0D: //TV[1]
			msg.AccessType = new(Uint8)
			*msg.AccessType = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
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
