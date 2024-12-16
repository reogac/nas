/**generated time: 2024-12-16 16:36:18.690427**/

package nas

/*******************************************************
 * CONFIGURATION UPDATE COMPLETE
 ******************************************************/
type ConfigurationUpdateComplete struct {
	MmHeader
}

func (msg *ConfigurationUpdateComplete) encode() (wire []byte, err error) {
	msg.msgType = ConfigurationUpdateCompleteMsgType //set message type to CONFIGURATION UPDATE COMPLETE
	wire = msg.headerBytes()
	return
}
func (msg *ConfigurationUpdateComplete) decodeBody(wire []byte) (err error) {
	if len(wire) > 0 {
		err = ErrTail
	}
	return
}
