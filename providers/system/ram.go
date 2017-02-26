package system

import (
	"fmt"
	"github.com/johhy/gdwmstatusbar/utils"
	"github.com/shirou/gopsutil/mem"
	"time"
)

//Date provider
type RAM struct {
	//Timeout for timer in milliseconds
	Timeout time.Duration
	//Channel for values
	channel chan uint64
	//previous value
	value uint64
}

//Start provider as goroutine
func (r *RAM) Start() {
	//Prepare steps
	r.channel = make(chan uint64)
	//Start goroutine
	go func() {
		for {
			v, err := mem.VirtualMemory()
			if err != nil {
				fmt.Println("Error ger ram:", err)
				return
			}
			r.channel <- v.Available
			time.Sleep(r.Timeout * time.Millisecond)
		}
	}()
}

//Get value from channel - non blocking
//When new value it save in prev and return it
//if no message in channel return prev value
func (r *RAM) GetValue() string {
	select {
	case val := <-r.channel:
		r.value = val
	default:
	}
	return utils.DigitFormat(r.value)
}
