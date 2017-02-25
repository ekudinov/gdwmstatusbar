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
			resp, err1 := http.Get(o.Url)
			defer resp.Body.Close()
			buf, err2 := ioutil.ReadAll(resp.Body)
			if err1 == nil && err2 == nil {
				outIP = strings.TrimSpace(string(buf))
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
