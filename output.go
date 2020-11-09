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

func (o *output) Parse() {
	var results []string
	for _, s := range o.buffer {
		results = append(results, s)
	}

	o.rendered = results
	o.buffer = nil
}

func (o *output) Render(text string) {
	fmt.Println(o.rendered)
}
