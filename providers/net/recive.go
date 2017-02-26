package net

import (
	"fmt"
	"github.com/johhy/gdwmstatusbar/utils"
	"github.com/shirou/gopsutil/net"
	"time"
)

//Amount recieve provider
type RECEIVE struct {
	//The interface from get data
	Interface string
	//Timeout for timer in milliseconds
	Timeout time.Duration
	//Channel for values
	channel chan uint64
	//previous value
	value uint64
}

//Start provider as goroutine
func (r *RECEIVE) Start() {
	//Prepare steps
	r.channel = make(chan uint64)
	//Start goroutine
	go func() {
		for {
			data, err := net.IOCounters(true)
			if err != nil {
				fmt.Println("Error receive:", err)
				return
			}
			for _, v := range data {
				if v.Name == r.Interface {
					r.channel <- v.BytesRecv
				}
			}
			time.Sleep(r.Timeout * time.Millisecond)
		}
	}()
}

//Get value from channel - non blocking
//When new value it save in prev and return it
//if no message in channel return prev value
func (r *RECEIVE) GetValue() string {
	select {
	case val := <-r.channel:
		r.value = val
	default:
	}
	return utils.DigitFormat(r.value)
}
