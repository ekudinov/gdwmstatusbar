package weather

import (
	"fmt"
	owm "github.com/briandowns/openweathermap"
	"os"
	"time"
)

//OpenWeatherMap provider
type OWM struct {
	//Timeout for requests on hour
	Timeout time.Duration
	//Language for represent result
	Language string
	//Cel or Faringate //C or F
	Metric string
	//City for get temperature
	City string
	//Channel for values
	channel chan float64
	//Previous value
	value float64
}

//Start provider as goroutine
func (o *OWM) Start() {
	//Prepare steps
	if os.Getenv("OWM_API_KEY") == "" {
		fmt.Println("Error You must set env: export OWM_API_KEY=you key there")
	}

	o.channel = make(chan float64)
	w, err := owm.NewCurrent(o.Metric, o.Language)
	if err != nil {
		fmt.Println("OWM error:", err)
		return
	}
	//Start goroutine
	go func() {
		for {
			w.CurrentByName(o.City)
			o.channel <- w.Main.Temp
			time.Sleep(o.Timeout * time.Hour) // timout for requests
		}
	}()

}

//value from inner type convert to string (out type)
func (o *OWM) convert() string {
	return fmt.Sprintf("%.f %s", o.value, o.Metric)
}

//Get value from channel - non blocking
//When new value it save in prev
//and return
//if no message in channel return prev value
func (o *OWM) GetValue() string {
	select {
	case val := <-o.channel:
		o.value = val
	default:
	}
	return o.convert()
}
