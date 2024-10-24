/**generated time: 2024-07-17 15:11:00.944068**/

package nas

/*******************************************************
 * AUTHENTICATION FAILURE
 ******************************************************/
type AuthenticationFailure struct {
	MmHeader
	GmmCause                       Uint8  //V [1]
	AuthenticationFailureParameter *Bytes //TLV [30][16]
}

func (msg *AuthenticationFailure) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding AuthenticationFailure", err)
		}
	}()
	var buf []byte
	// V[1]
	wire = append(wire, uint8(msg.GmmCause))

	if msg.AuthenticationFailureParameter != nil {
		// TLV[16]
		if buf, err = encodeLV(false, uint16(14), uint16(14), msg.AuthenticationFailureParameter); err != nil {
			err = nasError("encoding AuthenticationFailureParameter [O TLV 16]", err)
			return
		}
		wire = append(append(wire, 0x30), buf...)
	}

	msg.msgType = AuthenticationFailureMsgType //set message type to AUTHENTICATION FAILURE
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *AuthenticationFailure) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding AuthenticationFailure", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// V[1]
	if offset+1 > wireLen {
		err = nasError("decoding GmmCause [M V 1]", ErrIncomplete)
		return
	}
	msg.GmmCause = Uint8(wire[offset])
	offset++

	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x30: //TLV[16]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(14), uint16(14), v); err != nil {
				err = nasError("decoding AuthenticationFailureParameter [O TLV 16]", err)
				return
			}
			offset += consumed
			msg.AuthenticationFailureParameter = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
