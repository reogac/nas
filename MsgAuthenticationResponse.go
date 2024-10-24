/**generated time: 2024-07-17 15:11:00.943897**/

package nas

/*******************************************************
 * AUTHENTICATION RESPONSE
 ******************************************************/
type AuthenticationResponse struct {
	MmHeader
	AuthenticationResponseParameter *Bytes //TLV [2D][18]
	EapMessage                      *Bytes //TLV-E [78][7-1503]
}

func (msg *AuthenticationResponse) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding AuthenticationResponse", err)
		}
	}()
	var buf []byte
	if msg.AuthenticationResponseParameter != nil {
		// TLV[18]
		if buf, err = encodeLV(false, uint16(16), uint16(16), msg.AuthenticationResponseParameter); err != nil {
			err = nasError("encoding AuthenticationResponseParameter [O TLV 18]", err)
			return
		}
		wire = append(append(wire, 0x2D), buf...)
	}

	if msg.EapMessage != nil {
		// TLV-E[7-1503]
		if buf, err = encodeLV(true, uint16(4), uint16(1500), msg.EapMessage); err != nil {
			err = nasError("encoding EapMessage [O TLV-E 7-1503]", err)
			return
		}
		wire = append(append(wire, 0x78), buf...)
	}

	msg.msgType = AuthenticationResponseMsgType //set message type to AUTHENTICATION RESPONSE
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *AuthenticationResponse) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding AuthenticationResponse", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x2D: //TLV[18]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(16), uint16(16), v); err != nil {
				err = nasError("decoding AuthenticationResponseParameter [O TLV 18]", err)
				return
			}
			offset += consumed
			msg.AuthenticationResponseParameter = v
		case 0x78: //TLV-E[7-1503]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), v); err != nil {
				err = nasError("decoding EapMessage [O TLV-E 7-1503]", err)
				return
			}
			offset += consumed
			msg.EapMessage = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
