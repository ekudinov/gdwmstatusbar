package net

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"time"
)

//Local ip provider
type LocalIP struct {
	//The interface from local ip get
	//wlan1,wlan0 for examplle
	Interface string
	//Timeout for timer in milliseconds
	Timeout time.Duration
	//Channel for values
	channel chan string
	//previous value
	value string
}

//Start provider as goroutine
func (l *LocalIP) Start() {
	//Prepare steps
	l.channel = make(chan string)
	//Start goroutine
	go func() {
		for {
			//Set local ip as none
			//if exists it changes to real
			localIP := "None"
			//Get local ip
			ifaces, err := net.Interfaces()
			if err != nil {
				fmt.Println("Error get net interfaces:", err)
				return
			}
			for _, v := range ifaces {
				if v.Name == l.Interface && v.Flags[0] == "up" {
					if len(v.Addrs) > 0 {
						addr := v.Addrs[0]
						localIP = addr.Addr
					}
				}
			}
			l.channel <- localIP
			time.Sleep(l.Timeout * time.Millisecond)
		}
	}()
}

//Get value from channel - non blocking
//When new value it save in prev and return it
//if no message in channel return prev value
func (l *LocalIP) GetValue() string {
	select {
	case msg := <-l.channel:
		l.value = msg
	default:
	}
	return l.value
}
