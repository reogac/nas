/** this file was generated at 2024-12-16 17:55:27.331700 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * 5GSM STATUS
 ******************************************************/
type GsmStatus struct {
	SmHeader
	GsmCause uint8 //M: V [1]
}

func (msg *GsmStatus) encode() (wire []byte, err error) {
	// M: V[1]
	wire = append(wire, uint8(msg.GsmCause))

	msg.msgType = GsmStatusMsgType //set message type to 5GSM STATUS
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *GsmStatus) decodeBody(wire []byte) (err error) {
	offset := 0
	wireLen := len(wire)
	// M V[1]
	if offset+1 > wireLen {
		err = nasError("decoding GsmCause [M V 1]", ErrIncomplete)
		return
	}
	msg.GsmCause = wire[offset]
	offset++

	return
}
