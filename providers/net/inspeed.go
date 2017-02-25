package net

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"time"
)

//Amount recieve provider
type INSPEED struct {
	//The interface from get data
	Interface string
	//Timeout for timer in milliseconds
	Timeout time.Duration
	//Channel for values
	channel chan uint64
	//previous value of speed
	value uint64
	//time when get value
	time time.Time
	//previous amount recieve bytes
	prev uint64
}

//Start provider as goroutine
func (i *INSPEED) Start() {
	//Prepare steps
	i.channel = make(chan uint64)
	i.time = time.Now() //set start time
	//Start goroutine
	go func() {
		for {
			//set timeout for avoid divide on zero
			time.Sleep(i.Timeout * time.Millisecond)
			data, err := net.IOCounters(true)
			if err != nil {
				fmt.Println("Error in speed:", err)
				return
			}
			var inspeed uint64 = 0
			for _, v := range data {
				if v.Name == i.Interface {
					recv := v.BytesRecv
					//calculare speed
					//Why uint64? uint64 because the period is not so big
					inspeed = (recv - i.prev) / (uint64)(time.Since(i.time).Seconds())
					i.time = time.Now()
					i.prev = recv
				}
			}
			i.channel <- inspeed
		}
	}()
}

//Get value from channel - non blocking
//When new value it save in prev and return it
//if no message in channel return prev value
func (i *INSPEED) GetValue() string {
	select {
	case val := <-i.channel:
		i.value = val
	default:
	}
	return inseconds(i.value)
}
