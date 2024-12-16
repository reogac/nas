/**generated time: 2024-12-16 16:36:18.695663**/

package nas

/*******************************************************
 * 5GMM STATUS
 ******************************************************/
type GmmStatus struct {
	MmHeader
	GmmCause uint8 //M: V [1]
}

func (msg *GmmStatus) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding GmmStatus", err)
		}
	}()
	// M: V[1]
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
	// M V[1]
	if offset+1 > wireLen {
		err = nasError("decoding GmmCause [M V 1]", ErrIncomplete)
		return
	}
	msg.GmmCause = wire[offset]
	offset++

	return
}
