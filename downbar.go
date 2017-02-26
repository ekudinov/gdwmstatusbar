package main

import (
	"github.com/johhy/gdwmstatusbar/providers/net"
	"github.com/johhy/gdwmstatusbar/providers/system"
)

//Create bar on down of screen
//There is config for elements on bar
//and init operation
func DownBar() *Bar {
	bar := &Bar{
		Elements: []*Element{
			&Element{
				Field: "IP_L",
				Provider: &net.LocalIP{
					Interface: "wlan0",
					Timeout:   600000,
				},
			},
			&Element{
				Field: "IP_O",
				Provider: &net.OutIP{
					Url:     "http://ipinfo.io/ip",
					Wait:    1,
					Timeout: 100000000,
				},
			},
			&Element{
				Field: "CPU",
				Provider: &system.CPU{
					Timeout: 1000,
				},
			},
			&Element{
				Field: "MF",
				Provider: &system.RAM{
					Timeout: 1000,
				},
			},
			&Element{
				Field: "DF",
				Provider: &system.DISK{
					Path:    "/home/johhy",
					Timeout: 60000,
				},
			},
			&Element{
				Field: "Sent",
				Provider: &net.SENT{
					Interface: "wlan0",
					Timeout:   10000,
				},
			},
			&Element{
				Field: "Recv",
				Provider: &net.RECEIVE{
					Interface: "wlan0",
					Timeout:   10000,
				},
			},
			&Element{
				Field: "In",
				Provider: &net.INSPEED{
					Interface: "wlan0",
					Timeout:   1000,
				},
			},
			&Element{
				Field: "Out",
				Provider: &net.OUTSPEED{
					Interface: "wlan0",
					Timeout:   1000,
				},
			},
		},
	}
	//init elements inside bar
	bar.Init()
	return bar
}
