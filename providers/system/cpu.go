package system

import (
	"fmt"
	"github.com/shirou/gopsutil/load"
	"time"
)

//Date provider
type CPU struct {
	//Timeout for timer in milliseconds
	Timeout time.Duration
	//Channel for values
	channel chan float64
	//previous value
	value float64
}

//Start provider as goroutine
func (c *CPU) Start() {
	//Prepare steps
	c.channel = make(chan float64)
	//Start goroutine
	go func() {
		for {
			l, err := load.Avg()
			if err != nil {
				fmt.Println("Error cpu:", err)
				return
			}
			c.channel <- l.Load1
			time.Sleep(c.Timeout * time.Millisecond)
		}
	}()
}

//Get value from channel - non blocking
//When new value it save in prev and return it
//if no message in channel return prev value
func (c *CPU) GetValue() string {
	select {
	case val := <-c.channel:
		c.value = val
	default:
	}
	return convert(c.value)

}
