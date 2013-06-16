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

// Sample code that shows how to use the freenect package.
//
// Author: Adam Jensen <acj@linuxguy.org>

package main

import (
	"fmt"
	"freenect"
	"time"
)

func main() {
	d := freenect.NewFreenectDevice(0)
	fmt.Println("Number of devices: ", d.GetNumDevices())
	DoTilting(d)
	DoLed(d)
	fmt.Println("Saving RGBA image to TestRGBA.jpg")
	freenect.SaveRGBAFrame(d, "TestRGBA.jpg")
	fmt.Println("Saving IR image to TestIR.jpg")
	freenect.SaveIRFrame(d, "TestIR.jpg")
	fmt.Println("Saving depth image to TestDepth.jpg")
	freenect.SaveDepthFrame(d, "TestDepth.jpg")
	d.Stop()
	d.Shutdown()
}

func DoTilting(d *freenect.FreenectDevice) {
	fmt.Println("Tilting down")
	d.SetTiltDegs(-30)
	ts := freenect.TiltStatusCode(0)
	for i := 0; i < 3; i++ {
		time.Sleep(1000000000)
		ts = d.GetTiltStatus(d.GetTiltState())
		fmt.Println("\tTilt status: ", ts, "(", d.GetTiltDegs(d.GetTiltState()), " degrees)")
	}
	fmt.Println("Tilting up")
	d.SetTiltDegs(30)
	for i := 0; i < 3; i++ {
		time.Sleep(1000000000)
		ts = d.GetTiltStatus(d.GetTiltState())
		fmt.Println("\tTilt status: ", ts, "(", d.GetTiltDegs(d.GetTiltState()), " degrees)")
	}
	fmt.Println("Tilting level")
	d.SetTiltDegs(0)
	for i := 0; i < 3; i++ {
		time.Sleep(1000000000)
		ts = d.GetTiltStatus(d.GetTiltState())
		fmt.Println("\tTilt status: ", ts, "(", d.GetTiltDegs(d.GetTiltState()), " degrees)")
	}
}

func DoLed(d *freenect.FreenectDevice) {
	fmt.Println("Changing LED status")
	fmt.Println("\tOFF")
	d.SetLed(freenect.LED_OFF)
	time.Sleep(1000000000)
	fmt.Println("\tGREEN")
	d.SetLed(freenect.LED_GREEN)
	time.Sleep(1000000000)
	fmt.Println("\tRED")
	d.SetLed(freenect.LED_RED)
	time.Sleep(1000000000)
	fmt.Println("\tYELLOW")
	d.SetLed(freenect.LED_YELLOW)
	time.Sleep(1000000000)
	fmt.Println("\tBLINK YELLOW")
	d.SetLed(freenect.LED_BLINK_YELLOW)
	time.Sleep(3000000000)
	fmt.Println("\tBLINK_GREEN")
	d.SetLed(freenect.LED_BLINK_GREEN)
	time.Sleep(3000000000)
	fmt.Println("\tBLINK RED/YELLOW")
	d.SetLed(freenect.LED_BLINK_RED_YELLOW)
	time.Sleep(3000000000)
	fmt.Println("\tOFF")
	d.SetLed(freenect.LED_OFF)
}