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
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

func (o *output) Add(text string) {
	o.buffer = append(o.buffer, text)
}

func (o *output) Render() {
	for _, line := range o.buffer {
		fmt.Fprintf(o.writeTo, line)
	}
}
