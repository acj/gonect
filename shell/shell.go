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
//	"freenect"
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
			//freenect.SetTiltDegs(arg, 0)
			fmt.Printf("Tilting to %d degrees\n", arg)

			cmd = ""
			arg = OutOfRange
		case cmd == "led":
			fmt.Println("Cycling LEDs")
			break
		case cmd == "quit":
			return
		}

		fmt.Printf("gonect> ")
		s.Scan()
	} 
}
