/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Basic shell for using the Go freenect package.
//
// Author: Adam Jensen <acj@linuxguy.org>

package main

import (
	"demo"
	"fmt"
	"freenect"
	"os"
	"text/scanner"
	"strconv"
)

const OutOfRange = 1024

func main() {
	var s scanner.Scanner
	s.Init(os.Stdin)

	var cmd string
	var arg int = OutOfRange

	fmt.Print("gonect> ")

	tok := s.Scan()
	for tok != scanner.EOF {
		if scanner.TokenString(tok) == "Ident" {
			cmd = s.TokenText()
		}

		switch {
		case cmd == "tilt":
			s.Scan()
			arg, _ = strconv.Atoi(s.TokenText())
			if arg == OutOfRange { break }

			freenect.SetTiltDegs(arg, 0)

			cmd = ""
			arg = OutOfRange
		case cmd == "led":
			s.Scan()
			led_string := s.TokenText()
			flash_led(led_string)
			break
		case cmd == "rgb_frame":
			driver.TestRGBAFrame(0)
			break
		case cmd == "rgb_video":
			driver.TestVideo(0)
			break
		case cmd == "ir_video":
			driver.TestIR(0)
			break
		case cmd == "depth_video":
			driver.TestDepth(0)
			break
		case cmd == "quit":
			return
		}

		fmt.Printf("gonect> ")
		s.Scan()
	} 
}

func flash_led(led string) {
	switch {
	case led == "off":
		freenect.SetLed(freenect.LED_OFF, 0)
		break
	case led == "green":
		freenect.SetLed(freenect.LED_GREEN, 0)
		break
	case led == "red":
		freenect.SetLed(freenect.LED_RED, 0)
		break
	case led == "yellow":
		freenect.SetLed(freenect.LED_YELLOW, 0)
		break
	case led == "blink_yellow":
		freenect.SetLed(freenect.LED_BLINK_YELLOW, 0)
		break
	case led == "blink_green":
		freenect.SetLed(freenect.LED_BLINK_GREEN, 0)
		break
	case led == "blink_red_yellow":
		freenect.SetLed(freenect.LED_BLINK_RED_YELLOW, 0)
		break
	}
}