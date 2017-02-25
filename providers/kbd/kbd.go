package kbd

/*
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
#include <X11/XKBlib.h>
#include <kbd.h>
*/
import "C"

import "time"

//Keyboard status provider
type KBD struct {
	//Timeout for scan keyboard in milliseconds
	Timeout time.Duration
	//Channel to data
	channel chan string
	//Previous value
	value string
}

// opening the connection with X
var dpyKbd = C.XOpenDisplay(nil)

//Start keyboard provider as goroutine
func (k *KBD) Start() {
	//Prepare and init
	k.channel = make(chan string)
	C.XkbAllocKeyboard()
	go func() {
		for {
			k.channel <- C.GoString(C.get_kbd_state(dpyKbd))
			time.Sleep(k.Timeout * time.Millisecond)
		}
	}()
}

//Get value from channel - non blocking
//When new value it save in prev and return
//if no message in channel return prev value
func (k *KBD) GetValue() string {
	select {
	case val := <-k.channel:
		k.value = val
	default:
	}
	return k.value
}
