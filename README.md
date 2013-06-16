# Wrapper for the [libfreenect](https://github.com/OpenKinect/libfreenect) library written in Go.

## Installation and Usage

First, be sure that you have installed [libfreenect](https://github.com/OpenKinect/libfreenect). If it's
installed in a strange place, you may need to edit the cgo flags in `freenect.go`.

Once the library is installed, dive in:
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
