package main

import (
	"strings"
)

//Bar represents array of elements on panel.
//Numeration from left to right
//or up to down for vertical bar
type Bar struct {
	Elements []*Element
}

//Bar is represents as string
func (b *Bar) String() string {
	delim := " " //between elements
	str := []string{}
	for _, v := range b.Elements {
		str = append(str, v.String())
	}
	return strings.Join(str, delim)
}

//Start all providers on elements
func (b *Bar) Init() {
	for _, el := range b.Elements {
		el.Init()
	}
}

//Element to show on panel bar
type Element struct {
	//Name of element on screen
	Field string
	//Provider for data
	Provider Provider
}

//Element is represents as string
func (e *Element) String() string {
	return strings.Join(
		[]string{e.Field, e.Provider.GetValue()}, ":")
}

//Start provider in element
func (e *Element) Init() {
	e.Provider.Start()
}

//Provider that get value from out service
//and keeps it on inner format and
//then convert inner format to this
//Must implement this interface
type Provider interface {
	//Start provider as goroutine
	Start()
	//Get value from provider
	GetValue() string
}
