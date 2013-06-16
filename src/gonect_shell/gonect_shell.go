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
	"fmt"
	"github.com/acj/gonect/src/freenect"
	"os"
	"text/scanner"
	"strconv"
)

const OutOfRange = 1024
const StartWithDevice = 0

func main() {
	d := freenect.NewFreenectDevice(StartWithDevice)

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
		case cmd == "help":
			showHelpMessage()
			break
		case cmd == "tilt":
			s.Scan()
			arg, _ = strconv.Atoi(s.TokenText())
			if arg == OutOfRange { break }

			d.SetTiltDegs(arg)

			cmd = ""
			arg = OutOfRange
		case cmd == "led":
			s.Scan()
			led_string := s.TokenText()
			flash_led(d, led_string)
			break
		case cmd == "rgb_frame":
			s.Scan()
			filename := s.TokenText()
			freenect.SaveRGBAFrame(d, filename)
			break
		case cmd == "ir_frame":
			s.Scan()
			filename := s.TokenText()
			freenect.SaveIRFrame(d, filename)
			break
		case cmd == "depth_frame":
			s.Scan()
			filename := s.TokenText()
			freenect.SaveDepthFrame(d, filename)
			break
		case cmd == "quit":
			return
		}

		fmt.Printf("gonect> ")
		s.Scan()
	} 
}

func showHelpMessage() {
	fmt.Println(
		"Available commands:\n\n",
		"tilt <degrees [-30,30]>\tTilt the Kinect\n",
		"led <off,green,red,yellow,blink_yellow,blink_green,blink_red_yellow>\n",
		"\t\t\t\tChange the color of the LED indicator\n",
		"rgb_frame <filename.jpg>\tSave an RGBA frame from the color camera\n",
		"ir_frame <filename.jpg>\tSave an infrared frame from the IR camera\n",
		"depth_frame <filename.jpg>\tSave a B&W frame from the depth camera\n",
		"quit\t\t\t\tExit this shell interpreter")
}

func flash_led(d *freenect.FreenectDevice, led string) {
	switch {
	case led == "off":
		d.SetLed(freenect.LED_OFF)
		break
	case led == "green":
		d.SetLed(freenect.LED_GREEN)
		break
	case led == "red":
		d.SetLed(freenect.LED_RED)
		break
	case led == "yellow":
		d.SetLed(freenect.LED_YELLOW)
		break
	case led == "blink_yellow":
		d.SetLed(freenect.LED_BLINK_YELLOW)
		break
	case led == "blink_green":
		d.SetLed(freenect.LED_BLINK_GREEN)
		break
	case led == "blink_red_yellow":
		d.SetLed(freenect.LED_BLINK_RED_YELLOW)
		break
	}
}