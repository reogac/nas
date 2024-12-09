/**generated time: 2024-07-17 15:11:00.937494**/

package nas

/*******************************************************
 * DEREGISTRATION ACCEPT TO UE
 ******************************************************/
type DeregistrationAcceptToUe struct {
	MmHeader
}

func (msg *DeregistrationAcceptToUe) encode() (wire []byte, err error) {
	msg.msgType = DeregistrationAcceptToUeMsgType //set message type to DEREGISTRATION ACCEPT TO UE
	wire = msg.headerBytes()
	return
}
func (msg *DeregistrationAcceptToUe) decodeBody(wire []byte) (err error) {
	if len(wire) > 0 {
		err = ErrTail
	}
	return
}
