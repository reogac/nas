/**generated time: 2024-07-17 15:11:00.940595**/

package nas

/*******************************************************
 * REGISTRATION COMPLETE
 ******************************************************/
type RegistrationComplete struct {
	MmHeader
	SorTransparentContainer *Bytes //TLV-E [73][20]
}

func (msg *RegistrationComplete) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding RegistrationComplete", err)
		}
	}()
	var buf []byte
	if msg.SorTransparentContainer != nil {
		// TLV-E[20]
		if buf, err = encodeLV(true, uint16(17), uint16(17), msg.SorTransparentContainer); err != nil {
			err = nasError("encoding SorTransparentContainer [O TLV-E 20]", err)
			return
		}
		wire = append(append(wire, 0x73), buf...)
	}

	msg.msgType = RegistrationCompleteMsgType //set message type to REGISTRATION COMPLETE
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *RegistrationComplete) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding RegistrationComplete", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x73: //TLV-E[20]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(17), uint16(17), v); err != nil {
				err = nasError("decoding SorTransparentContainer [O TLV-E 20]", err)
				return
			}
			offset += consumed
			msg.SorTransparentContainer = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
