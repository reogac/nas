/** this file was generated at 2024-12-16 17:55:27.328487 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * NOTIFICATION RESPONSE
 ******************************************************/
type NotificationResponse struct {
	MmHeader
	PduSessionStatus *PduSessionStatus //O: TLV [50][4-34]
}

func (msg *NotificationResponse) encode() (wire []byte, err error) {
	var buf []byte
	// O: TLV[4-34]
	if msg.PduSessionStatus != nil {
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
	offset := 0
	wireLen := len(wire)
	consumed := 0
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x50: //O: TLV[4-34]
			offset++ //consume IEI
			v := new(PduSessionStatus)
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
