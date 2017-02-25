package net

import (
	"bytes"
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

//transform to traffic in bit/second
//used for conver from inner format to outer
func inseconds(sent uint64) string {
	var buffer bytes.Buffer
	buffer.WriteString(digitFormat(sent * 8))
	buffer.WriteString("it/s")
	return buffer.String()
}
