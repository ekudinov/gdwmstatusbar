package main

// #cgo LDFLAGS: -lX11
// #include <X11/Xlib.h>
import "C"

import "time"
import "strings"

// opening the connection with X
var dpy = C.XOpenDisplay(nil)

//show statusbar on window
func showOnWindow(str *C.char) {
	C.XStoreName(dpy, C.XDefaultRootWindow(dpy), str)
	C.XSync(dpy, 1)
}

//convert from go strings to c
func convert(str string) *C.char {
	return C.CString(str)

}

//Create message from bars as string
func createMsg(bars ...*Bar) string {
	str := []string{}
	for _, v := range bars {
		str = append(str, v.String())
	}
	return strings.Join(str, ";")

}

func main() {
	//create bars
	tbar := TopBar()
	dbar := DownBar()
	for {
		time.Sleep(500 * time.Millisecond)
		showOnWindow(convert(createMsg(tbar, dbar)))
	}
}
