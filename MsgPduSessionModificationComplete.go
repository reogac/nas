/**generated time: 2024-12-16 16:36:18.698420**/

package nas

/*******************************************************
 * PDU SESSION MODIFICATION COMPLETE
 ******************************************************/
type PduSessionModificationComplete struct {
	SmHeader
	ExtendedProtocolConfigurationOptions *ExtendedProtocolConfigurationOptions //O: TLV-E [7B][4-65538]
	PortManagementInformationContainer   []byte                                //O: TLV-E [74][4-65538]
}

func (msg *PduSessionModificationComplete) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding PduSessionModificationComplete", err)
		}
	}()
	var buf []byte
	// O: TLV-E[4-65538]
	if msg.ExtendedProtocolConfigurationOptions != nil {
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.ExtendedProtocolConfigurationOptions); err != nil {
			err = nasError("encoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x7B), buf...)
	}

	// O: TLV-E[4-65538]
	if len(msg.PortManagementInformationContainer) > 0 {
		tmp := newBytesEncoder(msg.PortManagementInformationContainer)
		if buf, err = encodeLV(true, uint16(1), uint16(0), tmp); err != nil {
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
		case 0x7B: //O: TLV-E[4-65538]
			offset++ //consume IEI
			v := new(ExtendedProtocolConfigurationOptions)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedProtocolConfigurationOptions [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.ExtendedProtocolConfigurationOptions = v
		case 0x74: //O: TLV-E[4-65538]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding PortManagementInformationContainer [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.PortManagementInformationContainer = []byte(*v)
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
