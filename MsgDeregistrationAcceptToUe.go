/** this file was generated at 2024-12-16 17:55:27.322699 by tqtung@etri.re.kr **/

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
