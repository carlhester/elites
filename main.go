package main

import (
	"fmt"
)

type player struct {
	Name    string
	attribs map[string]Attrib
}

type Attrib interface {
	Get()
}

func (p *player) AddAttrib(name string, attrib Attrib) {
	p.attribs[name] = attrib
}

func (p *player) HasAttrib(name string) bool {
	if _, ok := p.attribs[name]; ok {
		return true
	}
	return false
}

func (p *player) RemoveAttrib(attrib string) {
	_, ok := p.attribs[attrib]
	if ok {
		delete(p.attribs, attrib)
	}
}

type health struct {
	hp int
}

func (h health) Get() {}

type heal struct {
	hp int
}

func (h heal) Get() {}

type game struct {
	p1   *player
	p2   *player
	turn int
}

func main() {
	var players []*player
	p1 := &player{Name: "P1", attribs: make(map[string]Attrib)}
	p2 := &player{Name: "P2", attribs: make(map[string]Attrib)}

	p1.AddAttrib("health", health{hp: 10})
	p2.AddAttrib("health", health{hp: 10})
	players = append(players, p1, p2)

	for turn := 0; turn < 10; turn++ {
		for _, p := range players {
			p.AddAttrib("heal", heal{hp: 5})
			DoHeal(p)
			fmt.Printf("%+v\n", p)
		}
	}
}

func DoHeal(p *player) {
	if p.HasAttrib("heal") {
		if p.HasAttrib("health") {
			newHp := p.attribs["health"].(health).hp + p.attribs["heal"].(heal).hp
			p.RemoveAttrib("heal")
			p.RemoveAttrib("health")
			p.AddAttrib("health", health{hp: newHp})
		}
	}
}
