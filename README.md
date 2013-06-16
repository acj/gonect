This is a wrapper for the [libfreenect](https://github.com/OpenKinect/libfreenect) library written in Go.

## Installation and Usage
$ go get github.com/acj/gonect/src/gonect_samples
$ go get github.com/acj/gonect/src/gonect_shell     # optional but recommended
$ gonect_samples

Tilting down
		 [...]
Tilting up
		 [...]
Tilting level
		 [...]
Changing LED status
		 [...]

TODO
====
* ~~Initialization/Shutdown~~
* ~~Tilting~~
* ~~LED colors~~
* ~~Get number of devices~~
* ~~Image capture~~
    * ~~RGB Camera~~
	* ~~Depth Camera~~
	* ~~IR Camera~~
* Video
    * RGB Camera
	* Depth Camera
	* IR Camera
* Microphones

* ~~Wrap libfreenect "C sync" functions~~
* Wrap libfreenect async functions