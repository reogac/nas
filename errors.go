package nas

import "fmt"

var (
	ErrIncomplete     error = fmt.Errorf("Incomplete data")
	ErrTail           error = fmt.Errorf("Junk tail")
	ErrInvalidSize    error = fmt.Errorf("Size violation")
	ErrUnknownIei     error = fmt.Errorf("Unknown IEI")
	ErrUnknownEpd     error = fmt.Errorf("Unknown EPD")
	ErrUnknownSec     error = fmt.Errorf("Unknown security type")
	ErrUnknownMsgType error = fmt.Errorf("Unknown message type")
)

type NasError struct {
	msg  string
	prev error
}

func nasError(msg string, prev error) error {
	return &NasError{
		msg:  msg,
		prev: prev,
	}
}

func (nasErr *NasError) Error() (errStr string) {
	errStr = nasErr.msg

	prev := nasErr.prev
	for prev != nil {
		if cur, ok := prev.(*NasError); ok {
			errStr = fmt.Sprintf("%s <<= %s", errStr, cur.msg)
			prev = cur.prev
		} else {
			errStr = fmt.Sprintf("%s <<= %s", errStr, prev.Error())
			prev = nil
		}
	}
	return
}
