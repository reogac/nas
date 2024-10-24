/**generated time: 2024-07-17 15:11:00.944856**/

package nas

/*******************************************************
 * SECURITY MODE REJECT
 ******************************************************/
type SecurityModeReject struct {
	MmHeader
	GmmCause Uint8 //V [1]
}

func (msg *SecurityModeReject) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding SecurityModeReject", err)
		}
	}()
	// V[1]
	wire = append(wire, uint8(msg.GmmCause))

	msg.msgType = SecurityModeRejectMsgType //set message type to SECURITY MODE REJECT
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *SecurityModeReject) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding SecurityModeReject", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	// V[1]
	if offset+1 > wireLen {
		err = nasError("decoding GmmCause [M V 1]", ErrIncomplete)
		return
	}
	msg.GmmCause = Uint8(wire[offset])
	offset++

	return
}
