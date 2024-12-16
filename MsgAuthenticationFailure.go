/** this file was generated at 2024-12-16 17:55:27.327771 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * AUTHENTICATION FAILURE
 ******************************************************/
type AuthenticationFailure struct {
	MmHeader
	GmmCause                       uint8  //M: V [1]
	AuthenticationFailureParameter []byte //O: TLV [30][16]
}

func (msg *AuthenticationFailure) encode() (wire []byte, err error) {
	var buf []byte
	// M: V[1]
	wire = append(wire, uint8(msg.GmmCause))

	// O: TLV[16]
	if len(msg.AuthenticationFailureParameter) > 0 {
		tmp := newBytesEncoder(msg.AuthenticationFailureParameter)
		if buf, err = encodeLV(false, uint16(14), uint16(14), tmp); err != nil {
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
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// M V[1]
	if offset+1 > wireLen {
		err = nasError("decoding GmmCause [M V 1]", ErrIncomplete)
		return
	}
	msg.GmmCause = wire[offset]
	offset++

	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x30: //O: TLV[16]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(14), uint16(14), v); err != nil {
				err = nasError("decoding AuthenticationFailureParameter [O TLV 16]", err)
				return
			}
			offset += consumed
			msg.AuthenticationFailureParameter = []byte(*v)
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
