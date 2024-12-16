/** this file was generated at 2024-12-16 17:55:27.330176 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * PDU SESSION AUTHENTICATION COMPLETE
 ******************************************************/
type PduSessionAuthenticationComplete struct {
	SmHeader
	EapMessage                           []byte                                //M: LV-E [6-1502]
	ExtendedProtocolConfigurationOptions *ExtendedProtocolConfigurationOptions //O: TLV-E [7B][4-65538]
}

func (msg *PduSessionAuthenticationComplete) encode() (wire []byte, err error) {
	var buf []byte
	// M: LV-E[6-1502]
	tmp := newBytesEncoder(msg.EapMessage)
	if buf, err = encodeLV(true, uint16(4), uint16(1500), tmp); err != nil {
		err = nasError("encoding EapMessage [M LV-E 6-1502]", err)
		return
	}
	wire = append(wire, buf...)

	// O: TLV-E[4-65538]
	if msg.ExtendedProtocolConfigurationOptions != nil {
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.ExtendedProtocolConfigurationOptions); err != nil {
			err = nasError("encoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x7B), buf...)
	}

	msg.msgType = PduSessionAuthenticationCompleteMsgType //set message type to PDU SESSION AUTHENTICATION COMPLETE
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *PduSessionAuthenticationComplete) decodeBody(wire []byte) (err error) {
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// M LV-E[6-1502]
	v := new(bytesDecoder)
	if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), v); err != nil {
		err = nasError("decoding EapMessage [M LV-E 6-1502]", err)
		return
	}
	offset += consumed
	msg.EapMessage = []byte(*v)
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
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
