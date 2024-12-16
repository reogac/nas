/** this file was generated at 2024-12-16 17:55:27.327717 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * AUTHENTICATION REJECT
 ******************************************************/
type AuthenticationReject struct {
	MmHeader
	EapMessage []byte //O: TLV-E [78][7-1503]
}

func (msg *AuthenticationReject) encode() (wire []byte, err error) {
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

	msg.msgType = AuthenticationRejectMsgType //set message type to AUTHENTICATION REJECT
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *AuthenticationReject) decodeBody(wire []byte) (err error) {
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
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
