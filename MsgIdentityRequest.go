/** this file was generated at 2024-12-16 17:55:27.327928 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * IDENTITY REQUEST
 ******************************************************/
type IdentityRequest struct {
	MmHeader
	IdentityType uint8 //M: V [1/2]
}

func (msg *IdentityRequest) encode() (wire []byte, err error) {
	//M: V[1/2]
	v := (uint8(msg.IdentityType) & 0x0f) //fill righthalf
	wire = append(wire, v)

	msg.msgType = IdentityRequestMsgType //set message type to IDENTITY REQUEST
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *IdentityRequest) decodeBody(wire []byte) (err error) {
	offset := 0
	wireLen := len(wire)
	// M V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding IdentityType [M V 1/2]", ErrIncomplete)
		return
	}
	msg.IdentityType = 0x0f & wire[offset] //righthalf
	offset++

	return
}
