package net

import (
	"github.com/shirou/gopsutil/net"
	"time"
)

//Net status provider
type Net struct {
	//The interface status we need know
	//wlan1,wlan0 for example
	Interface string
	//Timeout for timer in milliseconds
	Timeout time.Duration
	//Show on bar status down
	Down string
	//Show on bar status error
	Error string
	//Show on bar status up
	Up string
	//Channel for values
	channel chan string
	//previous value
	value string
}

//Start provider as goroutine
func (n *Net) Start() {
	//Prepare steps
	n.channel = make(chan string)
	//Start goroutine
	go func() {
		for {
			status := n.Down
			ifaces, err := net.Interfaces()
			if err != nil {
				status = n.Error
			}
			for _, v := range ifaces {
				if v.Name == n.Interface && v.Flags[0] == "up" {
					status = n.Up
				}
			}
			n.channel <- status
			time.Sleep(n.Timeout * time.Millisecond)
		}
	}()
}

//Get value from channel - non blocking
//When new value it save in prev and return it
//if no message in channel return prev value
func (n *Net) GetValue() string {
	select {
	case msg := <-n.channel:
		n.value = msg
	default:
	}
	return n.value
}
