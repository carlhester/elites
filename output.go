package main

import "fmt"

type Output struct {
	buffer   []string
	rendered []string
}

func (o *Output) Clear() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

func (o *Output) Add(text string) {
	o.buffer = append(o.buffer, text)
}

func (o *Output) Parse() {
	var results []string
	for _, s := range o.buffer {
		results = append(results, s)
	}

	o.rendered = results
	o.buffer = nil
}

func (o *Output) Render(text string) {
	fmt.Println(o.rendered)
}
