package net

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//Out ip provider
type OutIP struct {
	//Url service to get result
	Url string
	//Timeout for timer in milliseconds
	Timeout time.Duration
	//Wait timeout for response in seconds
	Wait time.Duration
	//Channel for values
	channel chan string
	//previous value
	value string
}

//Start provider as goroutine
func (o *OutIP) Start() {
	//Prepare steps
	o.channel = make(chan string)
	//Start goroutine
	go func() {
		for {
			//Set out ip as none
			//if ip found it updates
			outIP := "None"
			timeout := time.Duration(o.Wait * time.Second)
			client := http.Client{
				Timeout: timeout,
			}
			if resp, err := client.Get(o.Url); err == nil {
				//remoute host ok
				if outIP == "None" { //ip not define
					defer resp.Body.Close()
					if buf, err := ioutil.ReadAll(resp.Body); err == nil {
						outIP = strings.TrimSpace(string(buf))
					}
				}
			}
			o.channel <- outIP
			time.Sleep(o.Timeout * time.Millisecond)
		}
	}()
}

//Get value from channel - non blocking
//When new value it save in prev and return it
//if no message in channel return prev value
func (o *OutIP) GetValue() string {
	select {
	case msg := <-o.channel:
		o.value = msg
	default:
	}
	return o.value
}
