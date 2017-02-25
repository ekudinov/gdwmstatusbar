package net

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"time"
)

//Amount recieve provider
type OUTSPEED struct {
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
	//previous amount sent bytes
	prev uint64
}

//Start provider as goroutine
func (o *OUTSPEED) Start() {
	//Prepare steps
	o.channel = make(chan uint64)
	o.time = time.Now() //set start time
	//Start goroutine
	go func() {
		for {
			//set timeout for avoid divide on zero
			time.Sleep(o.Timeout * time.Millisecond)
			data, err := net.IOCounters(true)
			if err != nil {
				fmt.Println("Error out speed:", err)
				return
			}
			var outspeed uint64 = 0
			for _, v := range data {
				if v.Name == o.Interface {
					sent := v.BytesSent
					//calculare speed
					//Why uint64? uint64 because the period is not so big
					outspeed = (sent - o.prev) / (uint64)(time.Since(o.time).Seconds())
					o.time = time.Now()
					o.prev = sent
				}
			}
			o.channel <- outspeed
		}
	}()
}

//Get value from channel - non blocking
//When new value it save in prev and return it
//if no message in channel return prev value
func (o *OUTSPEED) GetValue() string {
	select {
	case val := <-o.channel:
		o.value = val
	default:
	}
	return inseconds(o.value)
}
