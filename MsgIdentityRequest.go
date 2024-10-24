/**generated time: 2024-07-17 15:11:00.944284**/

package nas

/*******************************************************
 * IDENTITY REQUEST
 ******************************************************/
type IdentityRequest struct {
	MmHeader
	IdentityType Uint8 //V [1/2]
}

func (msg *IdentityRequest) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding IdentityRequest", err)
		}
	}()
	//V[1/2]
	v := (uint8(msg.IdentityType) & 0x0f) //fill righthalf
	wire = append(wire, v)

	msg.msgType = IdentityRequestMsgType //set message type to IDENTITY REQUEST
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *IdentityRequest) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding IdentityRequest", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	// V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding IdentityType [M V 1/2]", ErrIncomplete)
		return
	}
	msg.IdentityType = Uint8(0x0f & wire[offset]) //righthalf
	offset++

	return
}
