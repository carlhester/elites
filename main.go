package main

import (
    "fmt"
)


type player struct {
    Name string
    attribs map[string]Attrib
}

type Attrib interface {
    Get()
}

func (p *player)AddAttrib(name string, attrib Attrib) {
    p.attribs[name] = attrib
}

func (p *player)HasAttrib(name string) bool{
    if _, ok := p.attribs[name]; ok  {
        return true
    }
    return false
}

func (p *player)RemoveAttrib(attrib string) { 
    _, ok := p.attribs[attrib]
    if ok { 
        delete(p.attribs, attrib)
    }
}

type health struct {
    hp int
}

func (h health)Get() {}

type game struct {
    p1  *player
    p2  *player
    turn int
}

func main() {
    var players []*player
    p1 := &player{Name: "P1", attribs: make(map[string]Attrib)}
    p2 := &player{Name: "P2", attribs: make(map[string]Attrib)}

    p1.AddAttrib("health", health{hp: 10})
    p2.AddAttrib("health", health{hp: 10})
    players = append(players, p1, p2)

    //g := &game{p1: p1, p2: p2,}

    for turn := 0; turn < 10; turn++ {
       // g.ShowStatus()
       for _, p := range players{
           fmt.Println(p.Name)
       // P1Turn()
       // ShowStatus()
       // P2Turn()

        }
    }

}
