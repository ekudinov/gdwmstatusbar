package date

import (
	"time"
)

//Date provider
type Dat struct {
	//Format presentation
	Format string
	//Timeout for timer in milliseconds
	Timeout time.Duration
	//Channel for values
	channel chan string
	//previous value
	value string
}

//Start provider as goroutine
func (d *Dat) Start() {
	//Prepare steps
	d.channel = make(chan string)
	//Start goroutine
	go func() {
		for {
			d.channel <- time.Now().Format(d.Format)
			time.Sleep(d.Timeout * time.Millisecond)
		}
	}()
}

//Get value from channel - non blocking
//When new value it save in prev and return it
//if no message in channel return prev value
func (d *Dat) GetValue() string {
	select {
	case val := <-d.channel:
		d.value = val
	default:
	}
	return d.value
}
