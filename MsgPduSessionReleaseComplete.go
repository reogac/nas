/**generated time: 2024-07-17 15:11:00.949436**/

package nas

/*******************************************************
 * PDU SESSION RELEASE COMPLETE
 ******************************************************/
type PduSessionReleaseComplete struct {
	SmHeader
	GsmCause                             *Uint8                                //TV [59][2]
	ExtendedProtocolConfigurationOptions *ExtendedProtocolConfigurationOptions //TLV-E [7B][4-65538]
}

func (msg *PduSessionReleaseComplete) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding PduSessionReleaseComplete", err)
		}
	}()
	var buf []byte
	if msg.GsmCause != nil {
		//TV[2]
		wire = append(wire, []byte{0x59, uint8(*msg.GsmCause)}...)
	}

	if msg.ExtendedProtocolConfigurationOptions != nil {
		// TLV-E[4-65538]
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.ExtendedProtocolConfigurationOptions); err != nil {
			err = nasError("encoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x7B), buf...)
	}

	msg.msgType = PduSessionReleaseCompleteMsgType //set message type to PDU SESSION RELEASE COMPLETE
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *PduSessionReleaseComplete) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding PduSessionReleaseComplete", err)
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
		case 0x7B: //TLV-E[4-65538]
			offset++ //consume IEI
			v := &ExtendedProtocolConfigurationOptions{}
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.ExtendedProtocolConfigurationOptions = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
