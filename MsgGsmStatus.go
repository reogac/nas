/**generated time: 2024-12-16 16:36:18.698975**/

package nas

/*******************************************************
 * 5GSM STATUS
 ******************************************************/
type GsmStatus struct {
	SmHeader
	GsmCause uint8 //M: V [1]
}

func (msg *GsmStatus) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding GsmStatus", err)
		}
	}()
	// M: V[1]
	wire = append(wire, uint8(msg.GsmCause))

	msg.msgType = GsmStatusMsgType //set message type to 5GSM STATUS
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *GsmStatus) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding GsmStatus", err)
		}
	}()
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
