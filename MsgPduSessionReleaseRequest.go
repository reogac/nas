/** this file was generated at 2024-12-16 17:55:27.331308 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * PDU SESSION RELEASE REQUEST
 ******************************************************/
type PduSessionReleaseRequest struct {
	SmHeader
	GsmCause                             *uint8                                //O: TV [59][2]
	ExtendedProtocolConfigurationOptions *ExtendedProtocolConfigurationOptions //O: TLV-E [7B][4-65538]
}

func (msg *PduSessionReleaseRequest) encode() (wire []byte, err error) {
	var buf []byte
	// O: TV[2]
	if msg.GsmCause != nil {
		wire = append(wire, []byte{0x59, uint8(*msg.GsmCause)}...)
	}

	// O: TLV-E[4-65538]
	if msg.ExtendedProtocolConfigurationOptions != nil {
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.ExtendedProtocolConfigurationOptions); err != nil {
			err = nasError("encoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x7B), buf...)
	}

	msg.msgType = PduSessionReleaseRequestMsgType //set message type to PDU SESSION RELEASE REQUEST
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *PduSessionReleaseRequest) decodeBody(wire []byte) (err error) {
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
		case 0x7B: //O: TLV-E[4-65538]
			offset++ //consume IEI
			v := new(ExtendedProtocolConfigurationOptions)
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
