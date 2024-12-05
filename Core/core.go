package core

import (
	"fmt"
)

type Core struct {
	RunQ []*Process
	FreePages []*PhysicalPage
	BusyPages []*PhysicalPage
	AddressSpaceMax, AddressSpaceMin int
	ReqPageMax, ReqPageMin int
	ReqWorkSetMax, ReqWorkSetMin int
	WorkSetSizeMax, WorkSetSizeMin int
}

func (c *Core) Start(n int) {
	fmt.Println("Start system...")
	c.RunQ = make([]*Process, 0)
	c.FreePages = make([]*PhysicalPage, n)
	c.AddressSpaceMin = 10
	c.AddressSpaceMax = 20
	c.ReqPageMin = 100
	c.ReqPageMax = 150
	c.ReqWorkSetMin = 10
	c.ReqWorkSetMax = 15
	fmt.Println("System is ready to work!")
	// can add a few processes
	// розмір адр. простору, робочий набір, кількість звернень для генерації набору
	// квант часу - скільки буде звернень в цьому процесі.
	// статично - кількість сторінок
}

func (c *Core) CreateProcess() {
	process := new(Process)
	process.PageTable.Entries = make([]*PTE, 5) // 5 is rand val
	process.NReq = 125 //rand val from
	c.RunQ = append(c.RunQ, process)
}

func (c *Core) GenerateWorkingSet(process *Process) {
	process.WorkingSet.PageIndexies = make([]int, 12) //12 is rand val of working set
	for i := 0; i < 12; i++ {
		process.WorkingSet.PageIndexies[i] = 5 // form 0 to PTE count
	}
}
