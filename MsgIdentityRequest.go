/**generated time: 2024-12-16 16:36:18.695198**/

package nas

/*******************************************************
 * IDENTITY REQUEST
 ******************************************************/
type IdentityRequest struct {
	MmHeader
	IdentityType uint8 //M: V [1/2]
}

func (msg *IdentityRequest) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding IdentityRequest", err)
		}
	}()
	//M: V[1/2]
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
	// M V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding IdentityType [M V 1/2]", ErrIncomplete)
		return
	}
	msg.IdentityType = 0x0f & wire[offset] //righthalf
	offset++

	return
}
