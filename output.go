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
	/*
		for _, out := range o.writeTo {
			fmt.Fprintf(out, "\033[2J")
			fmt.Fprintf(out, "\033[H")
			o.Render()
		}
	*/
}

func (o *output) Add(text string) {
	o.buffer = append(o.buffer, text)
	o.buffer = append(o.buffer, "\n")
	o.Render()
}

func (o *output) Render() {
	for _, line := range o.buffer {
		for _, out := range o.writeTo {
			_, err := fmt.Fprintf(out, line)
			if err != nil {
				fmt.Printf("%+w", err)
			}
		}
	}
	o.buffer = []string{}
}
