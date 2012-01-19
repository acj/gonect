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

// gonect package, a wrapper for the libfreenect library.
//
// Author: Adam Jensen <acj@linuxguy.org>

package gonect

/*
#cgo CFLAGS: -I/usr/local/include/libfreenect 
#cgo LDFLAGS: -lfreenect -lfreenect_sync
#include <stdlib.h>
#include <stdio.h>
#include <libfreenect.h>
#include <libfreenect_sync.h>
freenect_context* f_ctx;

freenect_raw_tilt_state* create_tilt_state() {
    freenect_raw_tilt_state *ts = malloc(sizeof(freenect_raw_tilt_state));
    return ts;
}

int freenect_init_proxy() {
    return freenect_init(&f_ctx, NULL);
}
*/
import "C"
import "unsafe"

type TiltState struct {
    Accelerometer_x int16
    Accelerometer_y int16
    Accelerometer_z int16
    Tilt_angle	    int8
    Tilt_status     TiltStatusCode
}

type TiltStatusCode uint
const (
    STOPPED = 0			        /* 0x00 */
    MOVEMENT_LIMIT = 1		    /* 0x01 */
    MOVING_TO_NEW_POSITION = 4	/* 0x04 */
)

type LedOptions uint
const (
	LED_OFF              = 0
	LED_GREEN            = 1 
	LED_RED              = 2 
	LED_YELLOW           = 3 
	LED_BLINK_YELLOW     = 4 
	LED_BLINK_GREEN      = 5 
	LED_BLINK_RED_YELLOW = 6
)

func ConvertCTiltStructToGo(c_ts *C.freenect_raw_tilt_state) TiltState {
	ts := TiltState{
		int16(c_ts.accelerometer_x),
		int16(c_ts.accelerometer_y),
		int16(c_ts.accelerometer_z),
		int8(c_ts.tilt_angle),
		TiltStatusCode(int8(c_ts.tilt_status))}
	return ts
}

func ConvertGoTiltStructToC(ts TiltState) *C.freenect_raw_tilt_state {
	c_ts := C.create_tilt_state()
	c_ts.accelerometer_x = C.int16_t(ts.Accelerometer_x)
	c_ts.accelerometer_y = C.int16_t(ts.Accelerometer_y)
	c_ts.accelerometer_z = C.int16_t(ts.Accelerometer_z)
	c_ts.tilt_angle = C.int8_t(ts.Tilt_angle)
	c_ts.tilt_status = C.freenect_tilt_status_code(ts.Tilt_status)
	return c_ts
}

func Init() uint {
	return uint(C.freenect_init_proxy())
}

func GetVideo(device_index int) (unsafe.Pointer, uint32) {
	var data unsafe.Pointer
	var timestamp C.uint32_t
	out := C.freenect_sync_get_video(&data, &timestamp, C.int(device_index), C.FREENECT_VIDEO_RGB)
	if out > 0 {
		return nil, 0
	}
    return data, uint32(timestamp)
}

func GetDepth(device_index int) (unsafe.Pointer, uint32) {
	var data unsafe.Pointer
	var timestamp C.uint32_t
	out := C.freenect_sync_get_depth(&data, &timestamp, C.int(device_index), C.FREENECT_DEPTH_11BIT)
	if out > 0 {
		return nil, 0
	}
    return data, uint32(timestamp)
}

func GetTiltDegs(ts TiltState) float32 {
	c_ts := ConvertGoTiltStructToC(ts)
	return float32(C.freenect_get_tilt_degs(c_ts))
}

// Set the tilt angle (in degrees)
func SetTiltDegs(degs int, device_index int) uint {
    c_degs := C.int(degs)
    c_ndx := C.int(device_index)
    return uint(C.freenect_sync_set_tilt_degs(c_degs, c_ndx))
}

func GetTiltState(device_index int) TiltState {
    c_ts := (*C.freenect_raw_tilt_state)(C.create_tilt_state())
    C.freenect_sync_get_tilt_state(&c_ts, C.int(device_index))
	ts := ConvertCTiltStructToGo(c_ts)
    return ts
}

func GetTiltStatus(ts TiltState, device_index int) TiltStatusCode {
	tsc := TiltStatusCode(C.freenect_get_tilt_status(ConvertGoTiltStructToC(ts)))
	return tsc
}

func SetLed(color uint, device_index int) uint {
    C.freenect_sync_set_led(C.freenect_led_options(color), C.int(device_index))
    return 0
}

func GetNumDevices() uint {
	return uint(C.freenect_num_devices(C.f_ctx))
}

func Stop() {
	C.freenect_sync_stop()
}

func Shutdown() {
	C.freenect_shutdown(C.f_ctx)
}