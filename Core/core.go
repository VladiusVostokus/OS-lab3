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
		for i := 0; i < n; i++ {
    	pte := &PTE{}
    	physPage := &PhysicalPage{PTE: pte, Number: i}
    	c.FreePages[i] = physPage
	}
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
	addressSpace := random(c.AddressSpaceMin, c.AddressSpaceMax)
	process.PageTable = new(PageTable)
	process.PageTable.Entries = make([]*PTE, addressSpace) // rand val
	for i := 0; i < addressSpace; i++ {
		pte := &PTE{}
		process.PageTable.Entries[i] = pte
	}
	reqPageCount := random(c.ReqPageMin, c.ReqPageMax)
	process.NReq = reqPageCount //rand val from
	c.RunQ = append(c.RunQ, process)
	fmt.Println("Create process")
}

func (c *Core) GenerateWorkingSet(process *Process) {
	workingSetCount := random(c.ReqWorkSetMin, c.ReqWorkSetMax)
	process.WorkingSet.PageIndexies = make([]int, workingSetCount) // rand val of working set
	for i := 0; i < workingSetCount; i++ {
		process.WorkingSet.PageIndexies[i] = i // form 0 to PTE count
	}
}

func (c *Core) GetProcess() *Process {
	return c.RunQ[0]
}

func (c *Core) PageFault(pageTable *PageTable, idx int) {
	var physPage **PhysicalPage
	if (len(c.FreePages) > 0) {
		index := random(0, len(c.FreePages))
		physPage = &c.FreePages[index]
		c.BusyPages = append(c.BusyPages, *physPage)
		fmt.Println("Map free page")
	} else {
		// Algoritm of page replacement
		index := random(0, len(c.BusyPages))
		physPage = &c.BusyPages[index]
		(*physPage).PTE.P = false
		fmt.Println("Replace page")
	}
	(*physPage).PTE = pageTable.Entries[idx]
	pageTable.Entries[idx].PNN = (*physPage).Number
	(*physPage).PTE.P = true
}
