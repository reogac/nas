/**generated time: 2024-12-16 16:36:18.695109**/

package nas

/*******************************************************
 * AUTHENTICATION RESULT
 ******************************************************/
type AuthenticationResult struct {
	MmHeader
	Ngksi      KeySetIdentifier //M: V [1/2]
	EapMessage []byte           //M: LV-E [6-1502]
	Abba       []byte           //O: TLV [38][4-n]
}

func (msg *AuthenticationResult) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding AuthenticationResult", err)
		}
	}()
	var buf []byte
	// M: V[1/2]
	if buf, err = msg.Ngksi.encode(); err != nil {
		err = nasError("encoding Ngksi [M V 1/2]", err)
		return
	}
	if len(buf) != 1 {
		err = nasError("encoding Ngksi [M V 1/2]", ErrInvalidSize)
		return
	}
	v := (buf[0] & 0x0f) //fill righthalf
	// M: LV-E[6-1502]
	wire = append(wire, v)

	tmp := newBytesEncoder(msg.EapMessage)
	if buf, err = encodeLV(true, uint16(4), uint16(1500), tmp); err != nil {
		err = nasError("encoding EapMessage [M LV-E 6-1502]", err)
		return
	}
	wire = append(wire, buf...)

	// O: TLV[4-n]
	if len(msg.Abba) > 0 {
		tmp := newBytesEncoder(msg.Abba)
		if buf, err = encodeLV(false, uint16(2), uint16(0), tmp); err != nil {
			err = nasError("encoding Abba [O TLV 4-n]", err)
			return
		}
		wire = append(append(wire, 0x38), buf...)
	}

	msg.msgType = AuthenticationResultMsgType //set message type to AUTHENTICATION RESULT
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *AuthenticationResult) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding AuthenticationResult", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// M V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding Ngksi [M V 1/2]", ErrIncomplete)
		return
	}
	if err = msg.Ngksi.decode([]byte{0x0f & wire[offset] /*righthalf*/}); err != nil {
		err = nasError("decoding Ngksi [M V 1/2]", err)
		return
	}
	// M LV-E[6-1502]
	offset++

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
		case 0x38: //O: TLV[4-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(0), v); err != nil {
				err = nasError("decoding Abba [O TLV 4-n]", err)
				return
			}
			offset += consumed
			msg.Abba = []byte(*v)
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
