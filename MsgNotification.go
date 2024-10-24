/**generated time: 2024-07-17 15:11:00.944984**/

package nas

/*******************************************************
 * NOTIFICATION
 ******************************************************/
type Notification struct {
	MmHeader
	AccessType Uint8 //V [1/2]
}

func (msg *Notification) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding Notification", err)
		}
	}()
	//V[1/2]
	v := (uint8(msg.AccessType) & 0x0f) //fill righthalf
	wire = append(wire, v)

	msg.msgType = NotificationMsgType //set message type to NOTIFICATION
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *Notification) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding Notification", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	// V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding AccessType [M V 1/2]", ErrIncomplete)
		return
	}
	msg.AccessType = Uint8(0x0f & wire[offset]) //righthalf
	offset++

	return
}
