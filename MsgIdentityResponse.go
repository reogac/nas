/**generated time: 2024-07-17 15:11:00.944350**/

package nas

/*******************************************************
 * IDENTITY RESPONSE
 ******************************************************/
type IdentityResponse struct {
	MmHeader
	MobileIdentity MobileIdentity //LV-E [3-n]
}

func (msg *IdentityResponse) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding IdentityResponse", err)
		}
	}()
	var buf []byte
	// LV-E[3-n]
	if buf, err = encodeLV(true, uint16(1), uint16(0), &msg.MobileIdentity); err != nil {
		err = nasError("encoding MobileIdentity [M LV-E 3-n]", err)
		return
	}
	wire = append(wire, buf...)

	msg.msgType = IdentityResponseMsgType //set message type to IDENTITY RESPONSE
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *IdentityResponse) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding IdentityResponse", err)
		}
	}()
	offset := 0
	consumed := 0
	// LV-E[3-n]
	if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), &msg.MobileIdentity); err != nil {
		err = nasError("decoding MobileIdentity [M LV-E 3-n]", err)
		return
	}
	offset += consumed
	return
}
