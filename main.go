package main

import (
    "fmt"
)


type player struct {
    ID int
    attribs map[string]Attrib
}

type Attrib interface { 
    Get()
}

func (p *player)AddAttrib(name string, attrib Attrib) { 
    p.attribs[name] = attrib
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

func (g game)ShowStatus() { 
    fmt.Printf("P1: %+v\n", g.p1)
    fmt.Printf("P2: %+v\n", g.p2)
}

func main() {
    p1 := &player{ID: 1, attribs: make(map[string]Attrib)}
    p2 := &player{ID: 2, attribs: make(map[string]Attrib)}

    p1.AddAttrib("health", health{hp: 10})
    p2.AddAttrib("health", health{hp: 10})
    
    
    g := &game{p1: p1, p2: p2,}

    for turn := 0; turn < 10; turn++ {
        g.ShowStatus()
       // P1Turn()
       // ShowStatus()
       // P2Turn()
    }

}
