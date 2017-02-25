package system

import (
	"fmt"
)

//return amount bytes as b,kb,Mb
//can used from providers
func digitFormat(bytes uint64) string {
	if bytes < 1000 {
		return fmt.Sprintf("%d B", bytes)
	}
	if bytes < 1000000 {
		return fmt.Sprintf("%d Kb", bytes/1000)
	}
	if bytes < 1000000000 {
		return fmt.Sprintf("%d Mb", bytes/1000000)
	}
	return fmt.Sprintf("%d Gb", bytes/1000000000)
}

//conver to string
func convert(val float64) string {
	return fmt.Sprintf("%.2f", val)
}
