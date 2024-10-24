/**generated time: 2024-07-17 15:11:00.947304**/

package nas

/*******************************************************
 * PDU SESSION AUTHENTICATION COMMAND
 ******************************************************/
type PduSessionAuthenticationCommand struct {
	SmHeader
	EapMessage                           Bytes                                 //LV-E [6-1502]
	ExtendedProtocolConfigurationOptions *ExtendedProtocolConfigurationOptions //TLV-E [7B][4-65538]
}

func (msg *PduSessionAuthenticationCommand) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding PduSessionAuthenticationCommand", err)
		}
	}()
	var buf []byte
	// LV-E[6-1502]
	if buf, err = encodeLV(true, uint16(4), uint16(1500), &msg.EapMessage); err != nil {
		err = nasError("encoding EapMessage [M LV-E 6-1502]", err)
		return
	}
	wire = append(wire, buf...)

	if msg.ExtendedProtocolConfigurationOptions != nil {
		// TLV-E[4-65538]
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.ExtendedProtocolConfigurationOptions); err != nil {
			err = nasError("encoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x7B), buf...)
	}

	msg.msgType = PduSessionAuthenticationCommandMsgType //set message type to PDU SESSION AUTHENTICATION COMMAND
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *PduSessionAuthenticationCommand) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding PduSessionAuthenticationCommand", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// LV-E[6-1502]
	if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), &msg.EapMessage); err != nil {
		err = nasError("decoding EapMessage [M LV-E 6-1502]", err)
		return
	}
	offset += consumed
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
