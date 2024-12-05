package core

import "fmt"

type Core struct {
	Process *Process
	RunQ []*Process
}

func (c *Core) Start() {
	fmt.Println("Start system...")
	c.RunQ = make([]*Process, 0)
	fmt.Println("System is ready to work!")
}