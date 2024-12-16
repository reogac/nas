/** this file was generated at 2024-12-16 17:55:27.328343 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * SECURITY MODE REJECT
 ******************************************************/
type SecurityModeReject struct {
	MmHeader
	GmmCause uint8 //M: V [1]
}

func (msg *SecurityModeReject) encode() (wire []byte, err error) {
	// M: V[1]
	wire = append(wire, uint8(msg.GmmCause))

	msg.msgType = SecurityModeRejectMsgType //set message type to SECURITY MODE REJECT
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *SecurityModeReject) decodeBody(wire []byte) (err error) {
	offset := 0
	wireLen := len(wire)
	// M V[1]
	if offset+1 > wireLen {
		err = nasError("decoding GmmCause [M V 1]", ErrIncomplete)
		return
	}
	msg.GmmCause = wire[offset]
	offset++

	return
}
