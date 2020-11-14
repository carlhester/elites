package main

import (
	"fmt"
	"io"
)

type output struct {
	buffer   []string
	rendered []string
	writeTo  []io.Writer
}

func NewOutput() *output {
	return &output{}
}

func (o *output) addWriteTo(writer io.Writer) {
	o.writeTo = append(o.writeTo, writer)
}

func (o *output) Clear() {
	for _, out := range o.writeTo {
		fmt.Fprintf(out, "\033[2J")
		fmt.Fprintf(out, "\033[H")
		o.Render()
	}
}

func (o *output) Add(text string) {
	o.buffer = append(o.buffer, text)
}

func (o *output) Render() {
	for _, out := range o.writeTo {
		for _, line := range o.buffer {
			fmt.Fprintf(out, line)
		}
	}
	o.buffer = []string{}
}
