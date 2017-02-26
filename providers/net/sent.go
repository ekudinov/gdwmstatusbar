package net

import (
	"fmt"
	"github.com/johhy/gdwmstatusbar/utils"
	"github.com/shirou/gopsutil/net"
	"time"
)

//Amount sent provider
type SENT struct {
	//The inteface from get data
	Interface string
	//Timeout for timer in milliseconds
	Timeout time.Duration
	//Channel for values
	channel chan uint64
	//previous value
	value uint64
}

//Start provider as goroutine
func (s *SENT) Start() {
	//Prepare steps
	s.channel = make(chan uint64)
	//Start goroutine
	go func() {
		for {
			data, err := net.IOCounters(true)
			if err != nil {
				fmt.Println("Error sent:", err)
				return
			}
			for _, v := range data {
				if v.Name == s.Interface {
					s.channel <- v.BytesSent
				}
			}
			time.Sleep(s.Timeout * time.Millisecond)
		}
	}()
}

//Get value from channel - non blocking
//When new value it save in prev and return it
//if no message in channel return prev value
func (s *SENT) GetValue() string {
	select {
	case val := <-s.channel:
		s.value = val
	default:
	}
	return utils.DigitFormat(s.value)
}
