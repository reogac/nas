/**generated time: 2024-12-16 16:36:18.697927**/

package nas

/*******************************************************
 * PDU SESSION MODIFICATION REJECT
 ******************************************************/
type PduSessionModificationReject struct {
	SmHeader
	GsmCause                             uint8                                 //M: V [1]
	BackOffTimerValue                    *GprsTimer3                           //O: TLV [37][3]
	GsmCongestionReAttemptIndicator      *uint8                                //O: TLV [61][3]
	ExtendedProtocolConfigurationOptions *ExtendedProtocolConfigurationOptions //O: TLV-E [7B][4-65538]
	ReAttemptIndicator                   *uint8                                //O: TLV [1D][3]
}

func (msg *PduSessionModificationReject) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding PduSessionModificationReject", err)
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

	// O: TLV[3]
	if msg.ReAttemptIndicator != nil {
		tmp := newUint8Encoder(*msg.ReAttemptIndicator)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
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
		case 0x1D: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding ReAttemptIndicator [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.ReAttemptIndicator = (*uint8)(v)
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
