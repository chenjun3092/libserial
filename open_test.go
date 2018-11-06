/*
 * Copyright Go-IIoT (https://github.com/goiiot)
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

package libserial

import (
	"flag"
	"fmt"
	"time"
)

var (
	inputPty  string
	outputPty string

	testOptions = []Option{
		WithBaudRate(1200),
		WithDataBits(8),
		WithParity(ParityNone),
		WithReadTimeout(time.Second),
		WithHardwareFlowControl(false),
		WithSoftwareFlowControl(false),
	}
)

func init() {
	flag.StringVar(&inputPty, "i", "", "input pty file path")
	flag.StringVar(&outputPty, "o", "", "input pty file path")
	flag.Parse()

	if inputPty == "" {
		panic("input pty is nil")
	}

	if outputPty == "" {
		panic("output pty is nil")
	}
}

func getSerialPort() (reader, writer *SerialPort) {
	w, err := Open(inputPty, testOptions...)
	if err != nil {
		panic(fmt.Sprintf("fatal err: %v", err))
	}

	r, err := Open(outputPty, testOptions...)
	if err != nil {
		panic(fmt.Sprintf("fatal err: %v", err))
	}

	return r, w
}