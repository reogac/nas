/*
* Copyright [2024] [Quang Tung Thai <tqtung@etri.re.kr>]
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */
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
