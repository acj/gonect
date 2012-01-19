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

// Example driver program for using the gonect package.
//
// Author: Adam Jensen <acj@linuxguy.org>

package driver

/*
#cgo CFLAGS: -I/usr/local/include/libfreenect 
#cgo LDFLAGS: -lfreenect -lfreenect_sync -lcv -lcxcore -lhighgui
#include <stdlib.h>
#include <stdio.h>
#include <libfreenect.h>
#include <libfreenect_sync.h>
#include <opencv/cv.h>
#include <opencv/highgui.h>
*/
import "C"

import (
    "gonect"
    "fmt"
    "time"
	"os"
	"opencv"
	"unsafe"
)

func Run() {
	gonect.Init()
	fmt.Println("Number of devices: ", gonect.GetNumDevices())
    //TestTilting(0)
    //TestLed(0)
	TestVideo(0)
	//TestDepth(0)
	gonect.Stop()
	gonect.Shutdown()
}

func TestVideo(device_index int) {
	win := opencv.NewWindow("image")
	for {
		data, _ := gonect.GetVideo(0)
		//out := C.freenect_sync_get_video(&data, &timestamp, C.int(device_index), C.FREENECT_VIDEO_RGB)
		cedge := opencv.CreateImage(640, 480, opencv.IPL_DEPTH_8U, 3)
		defer cedge.Release()
		C.cvSetData(unsafe.Pointer(cedge), data, C.int(1*3*640))
		win.ShowImage(cedge)
		if opencv.WaitKey(10) == 27 {
			os.Exit(0)
        }
	}
}

func TestDepth(device_index int) {
	win := opencv.NewWindow("image")
	for {
		data, _ := gonect.GetDepth(0)
		//out := C.freenect_sync_get_depth(&data, &timestamp, C.int(device_index), C.FREENECT_DEPTH_11BIT)
		cedge := opencv.CreateImage(640, 480, opencv.IPL_DEPTH_8U, 3)
		defer cedge.Release()
		C.cvSetData(unsafe.Pointer(cedge), data, C.int(1*3*640))
		win.ShowImage(cedge)
		if opencv.WaitKey(10) == 27 {
			os.Exit(0)
        }
	}
}

func TestTilting() {
    fmt.Println("Tilting down")
    gonect.SetTiltDegs(-30, 0)
	ts := gonect.TiltStatusCode(0)
	for i := 0; i < 3; i++ {
		time.Sleep(1000000000)
		ts = gonect.GetTiltStatus(gonect.GetTiltState(0), 0)
		fmt.Println("\tTilt status: ", ts, "(", gonect.GetTiltDegs(gonect.GetTiltState(0)), " degrees)")
	}
    fmt.Println("Tilting up")
    gonect.SetTiltDegs(30, 0)
	for i := 0; i < 3; i++ {
		time.Sleep(1000000000)
		ts = gonect.GetTiltStatus(gonect.GetTiltState(0), 0)
		fmt.Println("\tTilt status: ", ts, "(", gonect.GetTiltDegs(gonect.GetTiltState(0)), " degrees)")
	}
    fmt.Println("Tilting level")
    gonect.SetTiltDegs(0, 0)
	for i := 0; i < 3; i++ {
		time.Sleep(1000000000)
		ts = gonect.GetTiltStatus(gonect.GetTiltState(0), 0)
		fmt.Println("\tTilt status: ", ts, "(", gonect.GetTiltDegs(gonect.GetTiltState(0)), " degrees)")
	}
}

func TestLed() {
	fmt.Println("Changing LED status")
	fmt.Println("\tOFF")
    gonect.SetLed(gonect.LED_OFF, 0)
    time.Sleep(1000000000)
	fmt.Println("\tGREEN")
    gonect.SetLed(gonect.LED_GREEN, 0)
    time.Sleep(1000000000)
	fmt.Println("\tRED")
    gonect.SetLed(gonect.LED_RED, 0)
    time.Sleep(1000000000)
	fmt.Println("\tYELLOW")
    gonect.SetLed(gonect.LED_YELLOW, 0)
    time.Sleep(1000000000)
	fmt.Println("\tBLINK YELLOW")
    gonect.SetLed(gonect.LED_BLINK_YELLOW, 0)
    time.Sleep(3000000000)
	fmt.Println("\tBLINK_GREEN")
    gonect.SetLed(gonect.LED_BLINK_GREEN, 0)
    time.Sleep(3000000000)
	fmt.Println("\tBLINK RED/YELLOW")
    gonect.SetLed(gonect.LED_BLINK_RED_YELLOW, 0)
    time.Sleep(3000000000)
	fmt.Println("\tOFF")
    gonect.SetLed(gonect.LED_OFF, 0)
}