package system

import (
	"fmt"
	"github.com/johhy/gdwmstatusbar/utils"
	"github.com/shirou/gopsutil/disk"
	"time"
)

//Date provider
type DISK struct {
	//Path where check disk
	Path string
	//Timeout for timer in milliseconds
	Timeout time.Duration
	//Channel for values
	channel chan uint64
	//previous value
	value uint64
}

//Start provider as goroutine
func (d *DISK) Start() {
	//Prepare steps
	d.channel = make(chan uint64)
	//Start goroutine
	go func() {
		for {
			res, err := disk.Usage(d.Path)
			if err != nil {
				fmt.Println("Error ger disk:", err)
				return
			}
			d.channel <- res.Free
			time.Sleep(d.Timeout * time.Millisecond)
		}
	}()
}

//Get value from channel - non blocking
//When new value it save in prev and return it
//if no message in channel return prev value
func (d *DISK) GetValue() string {
	select {
	case val := <-d.channel:
		d.value = val
	default:
	}
	return utils.DigitFormat(d.value)
}
