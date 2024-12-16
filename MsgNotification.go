/** this file was generated at 2024-12-16 17:55:27.328439 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * NOTIFICATION
 ******************************************************/
type Notification struct {
	MmHeader
	AccessType uint8 //M: V [1/2]
}

func (msg *Notification) encode() (wire []byte, err error) {
	//M: V[1/2]
	v := (uint8(msg.AccessType) & 0x0f) //fill righthalf
	wire = append(wire, v)

	msg.msgType = NotificationMsgType //set message type to NOTIFICATION
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *Notification) decodeBody(wire []byte) (err error) {
	offset := 0
	wireLen := len(wire)
	// M V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding AccessType [M V 1/2]", ErrIncomplete)
		return
	}
	msg.AccessType = 0x0f & wire[offset] //righthalf
	offset++

	return
}
