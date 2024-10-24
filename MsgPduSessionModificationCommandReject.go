/**generated time: 2024-07-17 15:11:00.948889**/

package nas

/*******************************************************
 * PDU SESSION MODIFICATION COMMAND REJECT
 ******************************************************/
type PduSessionModificationCommandReject struct {
	SmHeader
	GsmCause                             Uint8                                 //V [1]
	ExtendedProtocolConfigurationOptions *ExtendedProtocolConfigurationOptions //TLV-E [7B][4-65538]
}

func (msg *PduSessionModificationCommandReject) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding PduSessionModificationCommandReject", err)
		}
	}()
	var buf []byte
	// V[1]
	wire = append(wire, uint8(msg.GsmCause))

	if msg.ExtendedProtocolConfigurationOptions != nil {
		// TLV-E[4-65538]
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.ExtendedProtocolConfigurationOptions); err != nil {
			err = nasError("encoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x7B), buf...)
	}

	msg.msgType = PduSessionModificationCommandRejectMsgType //set message type to PDU SESSION MODIFICATION COMMAND REJECT
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *PduSessionModificationCommandReject) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding PduSessionModificationCommandReject", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// V[1]
	if offset+1 > wireLen {
		err = nasError("decoding GsmCause [M V 1]", ErrIncomplete)
		return
	}
	msg.GsmCause = Uint8(wire[offset])
	offset++

	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x7B: //TLV-E[4-65538]
			offset++ //consume IEI
			v := &ExtendedProtocolConfigurationOptions{}
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.ExtendedProtocolConfigurationOptions = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
