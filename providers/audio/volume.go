package audio

/*
#cgo LDFLAGS: -lasound
#include <volume.h>
*/
import "C"

import "time"

//Sound volume status provider
type VOL struct {
	//Timeout for scan volume in milliseconds
	Timeout time.Duration
	//Channel to data
	channel chan string
	//Previous value
	value string
}

//Start sound volume goroutine
func (v *VOL) Start() {
	//Prepare and init
	v.channel = make(chan string)
	go func() {
		for {
			v.channel <- C.GoString(C.get_vol())
			time.Sleep(v.Timeout * time.Millisecond)
		}
	}()
}

//Get value from channel - non blocking
//When new value it save in prev and return
//if no message in channel return prev value
func (v *VOL) GetValue() string {
	select {
	case val := <-v.channel:
		v.value = val
	default:
	}
	return v.value
}
