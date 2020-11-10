package main

import (
	"fmt"
	"io"
)

type output struct {
	buffer   []string
	rendered []string
	writeTo  io.Writer
}

func NewOutput(writer io.Writer) *output {
	return &output{writeTo: writer}
}

func (o *output) Clear() {
	fmt.Fprintf(o.writeTo, "\033[2J")
	fmt.Fprintf(o.writeTo, "\033[H")
	o.Render()
}

func (o *output) Add(text string) {
	o.buffer = append(o.buffer, text)
}

func (o *output) Render() {
	for _, line := range o.buffer {
		fmt.Fprintf(o.writeTo, line)
	}
}
