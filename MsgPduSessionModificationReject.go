/**generated time: 2024-07-17 15:11:00.948090**/

package nas

/*******************************************************
 * PDU SESSION MODIFICATION REJECT
 ******************************************************/
type PduSessionModificationReject struct {
	SmHeader
	GsmCause                             Uint8                                 //V [1]
	BackOffTimerValue                    *GprsTimer3                           //TLV [37][3]
	GsmCongestionReAttemptIndicator      *Uint8                                //TLV [61][3]
	ExtendedProtocolConfigurationOptions *ExtendedProtocolConfigurationOptions //TLV-E [7B][4-65538]
	ReAttemptIndicator                   *Uint8                                //TLV [1D][3]
}

func (msg *PduSessionModificationReject) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding PduSessionModificationReject", err)
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

	if msg.ReAttemptIndicator != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.ReAttemptIndicator); err != nil {
			err = nasError("encoding ReAttemptIndicator [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x1D), buf...)
	}

	msg.msgType = PduSessionModificationRejectMsgType //set message type to PDU SESSION MODIFICATION REJECT
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *PduSessionModificationReject) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding PduSessionModificationReject", err)
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
		case 0x1D: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding ReAttemptIndicator [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.ReAttemptIndicator = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
