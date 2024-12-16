/** this file was generated at 2024-12-16 17:55:27.327978 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * IDENTITY RESPONSE
 ******************************************************/
type IdentityResponse struct {
	MmHeader
	MobileIdentity MobileIdentity //M: LV-E [3-n]
}

func (msg *IdentityResponse) encode() (wire []byte, err error) {
	var buf []byte
	// M: LV-E[3-n]
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
	offset := 0
	consumed := 0
	// M LV-E[3-n]
	if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), &msg.MobileIdentity); err != nil {
		err = nasError("decoding MobileIdentity [M LV-E 3-n]", err)
		return
	}
	offset += consumed
	return
}
