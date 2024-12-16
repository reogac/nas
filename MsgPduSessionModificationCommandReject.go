/**generated time: 2024-12-16 16:36:18.698499**/

package nas

/*******************************************************
 * PDU SESSION MODIFICATION COMMAND REJECT
 ******************************************************/
type PduSessionModificationCommandReject struct {
	SmHeader
	GsmCause                             uint8                                 //M: V [1]
	ExtendedProtocolConfigurationOptions *ExtendedProtocolConfigurationOptions //O: TLV-E [7B][4-65538]
}

func (msg *PduSessionModificationCommandReject) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding PduSessionModificationCommandReject", err)
		}
	}()
	var buf []byte
	// M: V[1]
	wire = append(wire, uint8(msg.GsmCause))

	// O: TLV-E[4-65538]
	if msg.ExtendedProtocolConfigurationOptions != nil {
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
	// M V[1]
	if offset+1 > wireLen {
		err = nasError("decoding GsmCause [M V 1]", ErrIncomplete)
		return
	}
	msg.GsmCause = wire[offset]
	offset++

	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x7B: //O: TLV-E[4-65538]
			offset++ //consume IEI
			v := new(ExtendedProtocolConfigurationOptions)
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
