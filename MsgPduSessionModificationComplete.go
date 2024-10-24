/**generated time: 2024-07-17 15:11:00.948783**/

package nas

/*******************************************************
 * PDU SESSION MODIFICATION COMPLETE
 ******************************************************/
type PduSessionModificationComplete struct {
	SmHeader
	ExtendedProtocolConfigurationOptions *ExtendedProtocolConfigurationOptions //TLV-E [7B][4-65538]
	PortManagementInformationContainer   *Bytes                                //TLV-E [74][4-65538]
}

func (msg *PduSessionModificationComplete) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding PduSessionModificationComplete", err)
		}
	}()
	var buf []byte
	if msg.ExtendedProtocolConfigurationOptions != nil {
		// TLV-E[4-65538]
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.ExtendedProtocolConfigurationOptions); err != nil {
			err = nasError("encoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x7B), buf...)
	}

	if msg.PortManagementInformationContainer != nil {
		// TLV-E[4-65538]
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.PortManagementInformationContainer); err != nil {
			err = nasError("encoding PortManagementInformationContainer [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x74), buf...)
	}

	msg.msgType = PduSessionModificationCompleteMsgType //set message type to PDU SESSION MODIFICATION COMPLETE
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *PduSessionModificationComplete) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding PduSessionModificationComplete", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
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
		case 0x74: //TLV-E[4-65538]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding PortManagementInformationContainer [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.PortManagementInformationContainer = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
