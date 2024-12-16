/**generated time: 2024-12-16 16:36:18.697511**/

package nas

/*******************************************************
 * PDU SESSION AUTHENTICATION RESULT
 ******************************************************/
type PduSessionAuthenticationResult struct {
	SmHeader
	EapMessage                           []byte                                //O: TLV-E [78][7-1503]
	ExtendedProtocolConfigurationOptions *ExtendedProtocolConfigurationOptions //O: TLV-E [7B][4-65538]
}

func (msg *PduSessionAuthenticationResult) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding PduSessionAuthenticationResult", err)
		}
	}()
	var buf []byte
	// O: TLV-E[7-1503]
	if len(msg.EapMessage) > 0 {
		tmp := newBytesEncoder(msg.EapMessage)
		if buf, err = encodeLV(true, uint16(4), uint16(1500), tmp); err != nil {
			err = nasError("encoding EapMessage [O TLV-E 7-1503]", err)
			return
		}
		wire = append(append(wire, 0x78), buf...)
	}

	// O: TLV-E[4-65538]
	if msg.ExtendedProtocolConfigurationOptions != nil {
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.ExtendedProtocolConfigurationOptions); err != nil {
			err = nasError("encoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x7B), buf...)
	}

	msg.msgType = PduSessionAuthenticationResultMsgType //set message type to PDU SESSION AUTHENTICATION RESULT
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *PduSessionAuthenticationResult) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding PduSessionAuthenticationResult", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x78: //O: TLV-E[7-1503]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), v); err != nil {
				err = nasError("decoding EapMessage [O TLV-E 7-1503]", err)
				return
			}
			offset += consumed
			msg.EapMessage = []byte(*v)
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
