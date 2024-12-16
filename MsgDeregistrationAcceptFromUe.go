/**generated time: 2024-12-16 16:36:18.690381**/

package nas

/*******************************************************
 * DEREGISTRATION ACCEPT FROM UE
 ******************************************************/
type DeregistrationAcceptFromUe struct {
	MmHeader
}

func (msg *DeregistrationAcceptFromUe) encode() (wire []byte, err error) {
	msg.msgType = DeregistrationAcceptFromUeMsgType //set message type to DEREGISTRATION ACCEPT FROM UE
	wire = msg.headerBytes()
	return
}
func (msg *DeregistrationAcceptFromUe) decodeBody(wire []byte) (err error) {
	if len(wire) > 0 {
		err = ErrTail
	}
	return
}
