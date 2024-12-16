/** this file was generated at 2024-12-16 17:55:27.327503 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * AUTHENTICATION REQUEST
 ******************************************************/
type AuthenticationRequest struct {
	MmHeader
	Ngksi                       KeySetIdentifier //M: V [1/2]
	Abba                        []byte           //M: LV [3-n]
	AuthenticationParameterRand []byte           //O: TV [21][17]
	AuthenticationParameterAutn []byte           //O: TLV [20][18]
	EapMessage                  []byte           //O: TLV-E [78][7-1503]
}

func (msg *AuthenticationRequest) encode() (wire []byte, err error) {
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
	// M: LV[3-n]
	wire = append(wire, v)

	tmp := newBytesEncoder(msg.Abba)
	if buf, err = encodeLV(false, uint16(2), uint16(0), tmp); err != nil {
		err = nasError("encoding Abba [M LV 3-n]", err)
		return
	}
	wire = append(wire, buf...)

	// O: TV[17]
	if len(msg.AuthenticationParameterRand) > 0 {
		tmp := newBytesEncoder(msg.AuthenticationParameterRand)
		if buf, err = tmp.encode(); err != nil {
			err = nasError("encoding AuthenticationParameterRand [O TV 17]", err)
			return
		}
		if len(buf) != 16 {
			err = nasError("encoding AuthenticationParameterRand [O TV 17]", ErrInvalidSize)
			return
		}
		wire = append(append(wire, 0x21), buf...)
	}

	// O: TLV[18]
	if len(msg.AuthenticationParameterAutn) > 0 {
		tmp := newBytesEncoder(msg.AuthenticationParameterAutn)
		if buf, err = encodeLV(false, uint16(16), uint16(16), tmp); err != nil {
			err = nasError("encoding AuthenticationParameterAutn [O TLV 18]", err)
			return
		}
		wire = append(append(wire, 0x20), buf...)
	}

	// O: TLV-E[7-1503]
	if len(msg.EapMessage) > 0 {
		tmp := newBytesEncoder(msg.EapMessage)
		if buf, err = encodeLV(true, uint16(4), uint16(1500), tmp); err != nil {
			err = nasError("encoding EapMessage [O TLV-E 7-1503]", err)
			return
		}
		wire = append(append(wire, 0x78), buf...)
	}

	msg.msgType = AuthenticationRequestMsgType //set message type to AUTHENTICATION REQUEST
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *AuthenticationRequest) decodeBody(wire []byte) (err error) {
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
	// M LV[3-n]
	offset++

	v := new(bytesDecoder)
	if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(0), v); err != nil {
		err = nasError("decoding Abba [M LV 3-n]", err)
		return
	}
	offset += consumed
	msg.Abba = []byte(*v)
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x21: //O: TV[17]
			if offset+17 > wireLen {
				err = nasError("decoding AuthenticationParameterRand [O TV 17]", ErrIncomplete)
				return
			}
			offset++ //consume IEI
			v := new(bytesDecoder)
			if err = v.decode(wire[offset : offset+16]); err != nil {
				err = nasError("decoding AuthenticationParameterRand [O TV 17]", err)
				return
			}
			msg.AuthenticationParameterRand = []byte(*v)
			offset += 16

		case 0x20: //O: TLV[18]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(16), uint16(16), v); err != nil {
				err = nasError("decoding AuthenticationParameterAutn [O TLV 18]", err)
				return
			}
			offset += consumed
			msg.AuthenticationParameterAutn = []byte(*v)
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
