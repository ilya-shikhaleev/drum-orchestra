package lib

import "log"

type Pool struct {
	name        string
	connections map[*Connection]bool
	AddAction   chan *Connection
	Join        chan *Connection
	Leave       chan *Connection
}

func (p *Pool) Name() string {
	return p.name
}

func (p *Pool) run() {
	for {
		log.Println("run pool select")
		select {
		case c := <-p.Leave:
			log.Println("p.leave catch", c)
			delete(p.connections, c)
			p.OnPlayerLeave(c)
		case c := <-p.Join:
			log.Println("p.join catch")
			p.connections[c] = true
			p.OnPlayerJoined(c)
		case c := <-p.AddAction:
			log.Println("p.addAction catch", c)
			p.OnActionAdded(c)
		}
	}
}

func (p *Pool) OnPlayerJoined(newPlayer *Connection) {
	for cp := range p.connections {
		cp.SendMessage("{\"event\":\"join\"}")
	}
}

func (p *Pool) OnPlayerLeave(newPlayer *Connection) {
	for c := range p.connections {
		c.SendMessage("{\"event\":\"leave\"}")
	}
}

func (p *Pool) OnActionAdded(c *Connection) {
	if c.action == "play" {
		for cp := range p.connections {
			cp.SendMessage("{\"event\":\"play\"}")
		}
	}
}

func NewPool() *Pool {
	r := &Pool{
		connections: make(map[*Connection]bool),
		AddAction:   make(chan *Connection),
		Join:        make(chan *Connection),
		Leave:       make(chan *Connection),
	}

	go r.run()
	return r
}
