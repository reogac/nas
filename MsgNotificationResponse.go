/**generated time: 2024-07-17 15:11:00.945046**/

package nas

/*******************************************************
 * NOTIFICATION RESPONSE
 ******************************************************/
type NotificationResponse struct {
	MmHeader
	PduSessionStatus *PduSessionStatus //TLV [50][4-34]
}

func (msg *NotificationResponse) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding NotificationResponse", err)
		}
	}()
	var buf []byte
	if msg.PduSessionStatus != nil {
		// TLV[4-34]
		if buf, err = encodeLV(false, uint16(2), uint16(32), msg.PduSessionStatus); err != nil {
			err = nasError("encoding PduSessionStatus [O TLV 4-34]", err)
			return
		}
		wire = append(append(wire, 0x50), buf...)
	}

	msg.msgType = NotificationResponseMsgType //set message type to NOTIFICATION RESPONSE
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *NotificationResponse) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding NotificationResponse", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x50: //TLV[4-34]
			offset++ //consume IEI
			v := &PduSessionStatus{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(32), v); err != nil {
				err = nasError("decoding PduSessionStatus [O TLV 4-34]", err)
				return
			}
			offset += consumed
			msg.PduSessionStatus = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
