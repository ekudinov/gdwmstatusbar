package main

import (
	"github.com/johhy/gdwmstatusbar/providers/audio"
	"github.com/johhy/gdwmstatusbar/providers/date"
	"github.com/johhy/gdwmstatusbar/providers/kbd"
	"github.com/johhy/gdwmstatusbar/providers/net"
	"github.com/johhy/gdwmstatusbar/providers/weather"
)

//Create bar on top of screen
//There is config for elements on bar
//And init operation
func TopBar() *Bar {
	bar := &Bar{
		Elements: []*Element{
			&Element{
				Field: "Wlan0",
				Provider: &net.Net{
					Interface: "wlan0",
					Timeout:   30000,
					Down:      "DW",
					Up:        "UP",
					Error:     "ER",
				},
			},
			&Element{
				Field: "Vol",
				Provider: &audio.VOL{
					Timeout: 1000,
				},
			},
			&Element{
				Field: "Kbd",
				Provider: &kbd.KBD{
					Timeout: 1000,
				},
			},
			&Element{
                Field: "Tmp",
				Provider: &weather.OWM{
					Timeout:  1,
					Language: "ru",
					Metric:   "C",
					City:     "Novosibirsk",
				},
			},
			&Element{
                Field: "Dat",
				Provider: &date.Dat{
					Format:  "Mon Jan _2 15:04:05",
					Timeout: 1000,
				},
			},
		},
	}
	//Init bar with elements
	bar.Init()
	return bar
}
