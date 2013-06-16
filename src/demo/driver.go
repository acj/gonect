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
#cgo LDFLAGS: -lfreenect -lfreenect_sync -lopencv_highgui -lopencv_core -lopencv_imgproc
#include <stdlib.h>
#include <stdio.h>
#include <libfreenect.h>
#include <libfreenect_sync.h>
#include <opencv/cv.h>
#include <opencv/highgui.h>
*/
import "C"

import (
	"freenect"
	"fmt"
	"image/jpeg"
	"os"
	"time"
	"unsafe"
)

func Run() {
	freenect.Init()
	fmt.Println("Number of devices: ", freenect.GetNumDevices())
	TestTilting(0)
	TestLed(0)
	TestVideo(0)
	TestIR(0)
	TestDepth(0)
	freenect.Stop()
	freenect.Shutdown()
}

func TestRGBAFrame(device_index int) {
	img := freenect.GetRGBAFrame(device_index)

 	toimg, _ := os.Create("test.jpg")
 	defer toimg.Close()

 	jpeg.Encode(toimg, img, &jpeg.Options{jpeg.DefaultQuality})
}

func TestVideo(device_index int) {
	fmt.Println("Testing RGB video. Press ESC to stop.")

	// for {
	// 	data, _ := freenect.GetVideo(device_index, freenect.FREENECT_VIDEO_RGB)
	// 	cedge := C.cvCreateImageHeader(C.cvSize(640, 480), 8, 3)
	// 	C.cvSetData(unsafe.Pointer(cedge), data, C.int(1*3*640))
	// 	C.cvCvtColor(unsafe.Pointer(cedge), unsafe.Pointer(cedge), C.CV_RGB2BGR)
	// 	C.cvShowImage(C.CString("RGB"), unsafe.Pointer(cedge))
	// 	if C.cvWaitKey(10) == 27 {
	// 		return
	// 	}
	// }
}

func TestIR(device_index int) {
	fmt.Println("Testing infrared. Press ESC to stop.")
	tempImage := C.cvLoadImage(C.CString("/Users/acj/scratch/2012-04-24_1338.jpg"), C.CV_LOAD_IMAGE_COLOR)
	C.cvNamedWindow( C.CString("IR"), C.CV_WINDOW_NORMAL )
	C.cvShowImage(C.CString("IR"), unsafe.Pointer(tempImage))
	cedge := C.cvCreateImageHeader(C.cvSize(640, 480), 8, 1)
	for {
		data, _ := freenect.GetVideo(device_index, freenect.FREENECT_VIDEO_IR_8BIT)
		C.cvSetData(unsafe.Pointer(cedge), data, C.int(1*1*640))
		C.cvShowImage(C.CString("IR"), unsafe.Pointer(cedge))
		if C.cvWaitKey(10) == 27 {
			return
		}
	}
}

func TestDepth(device_index int) {
	fmt.Println("Testing depth. Press ESC to stop.")
	for {
		data, _ := freenect.GetDepth(0, freenect.FREENECT_DEPTH_11BIT)
		cedge := C.cvCreateImageHeader(C.cvSize(640, 480), 16, 1)
		C.cvSetData(unsafe.Pointer(cedge), data, C.int(1*2*640))
		C.cvShowImage(C.CString("Depth"), unsafe.Pointer(cedge))
		if C.cvWaitKey(10) == 27 {
			return
		}
	}
}

func TestTilting(device_index int) {
	fmt.Println("Tilting down")
	freenect.SetTiltDegs(-30, 0)
	ts := freenect.TiltStatusCode(0)
	for i := 0; i < 3; i++ {
		time.Sleep(1000000000)
		ts = freenect.GetTiltStatus(freenect.GetTiltState(0), device_index)
		fmt.Println("\tTilt status: ", ts, "(", freenect.GetTiltDegs(freenect.GetTiltState(0)), " degrees)")
	}
	fmt.Println("Tilting up")
	freenect.SetTiltDegs(30, 0)
	for i := 0; i < 3; i++ {
		time.Sleep(1000000000)
		ts = freenect.GetTiltStatus(freenect.GetTiltState(0), device_index)
		fmt.Println("\tTilt status: ", ts, "(", freenect.GetTiltDegs(freenect.GetTiltState(0)), " degrees)")
	}
	fmt.Println("Tilting level")
	freenect.SetTiltDegs(0, 0)
	for i := 0; i < 3; i++ {
		time.Sleep(1000000000)
		ts = freenect.GetTiltStatus(freenect.GetTiltState(0), device_index)
		fmt.Println("\tTilt status: ", ts, "(", freenect.GetTiltDegs(freenect.GetTiltState(0)), " degrees)")
	}
}

func TestLed(device_index int) {
	fmt.Println("Changing LED status")
	fmt.Println("\tOFF")
	freenect.SetLed(freenect.LED_OFF, device_index)
	time.Sleep(1000000000)
	fmt.Println("\tGREEN")
	freenect.SetLed(freenect.LED_GREEN, device_index)
	time.Sleep(1000000000)
	fmt.Println("\tRED")
	freenect.SetLed(freenect.LED_RED, device_index)
	time.Sleep(1000000000)
	fmt.Println("\tYELLOW")
	freenect.SetLed(freenect.LED_YELLOW, device_index)
	time.Sleep(1000000000)
	fmt.Println("\tBLINK YELLOW")
	freenect.SetLed(freenect.LED_BLINK_YELLOW, device_index)
	time.Sleep(3000000000)
	fmt.Println("\tBLINK_GREEN")
	freenect.SetLed(freenect.LED_BLINK_GREEN, device_index)
	time.Sleep(3000000000)
	fmt.Println("\tBLINK RED/YELLOW")
	freenect.SetLed(freenect.LED_BLINK_RED_YELLOW, device_index)
	time.Sleep(3000000000)
	fmt.Println("\tOFF")
	freenect.SetLed(freenect.LED_OFF, device_index)
}
