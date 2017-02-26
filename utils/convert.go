package utils

import (
	"bytes"
	"fmt"
)

//In this package - common functions
//useful for providers

//return amount bytes as b,kb,Mb
//can used from providers
func DigitFormat(bytes uint64) string {
	switch {
	case bytes < 1000:
		return fmt.Sprintf("%d B", bytes)
	case bytes < 1000000:
		return fmt.Sprintf("%d Kb", bytes/1000)
	case bytes < 1000000000:
		return fmt.Sprintf("%d Mb", bytes/1000000)
	default:
		return fmt.Sprintf("%d Gb", bytes/1000000000)
	}
}

//transform to traffic in bit/second
func BitSeconds(amount uint64) string {
	var buffer bytes.Buffer
	buffer.WriteString(DigitFormat(amount * 8)) //in bit
	buffer.WriteString("it/s")
	return buffer.String()
}
