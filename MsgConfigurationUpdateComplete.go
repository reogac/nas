/**generated time: 2024-07-17 15:11:00.937516**/

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
