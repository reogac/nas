/** this file was generated at 2024-12-16 17:55:27.328393 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * 5GMM STATUS
 ******************************************************/
type GmmStatus struct {
	MmHeader
	GmmCause uint8 //M: V [1]
}

func (msg *GmmStatus) encode() (wire []byte, err error) {
	// M: V[1]
	wire = append(wire, uint8(msg.GmmCause))

	msg.msgType = GmmStatusMsgType //set message type to 5GMM STATUS
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *GmmStatus) decodeBody(wire []byte) (err error) {
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
