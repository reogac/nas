/**generated time: 2024-07-17 15:11:00.944922**/

package nas

/*******************************************************
 * 5GMM STATUS
 ******************************************************/
type GmmStatus struct {
	MmHeader
	GmmCause Uint8 //V [1]
}

func (msg *GmmStatus) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding GmmStatus", err)
		}
	}()
	// V[1]
	wire = append(wire, uint8(msg.GmmCause))

	msg.msgType = GmmStatusMsgType //set message type to 5GMM STATUS
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *GmmStatus) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding GmmStatus", err)
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
