package core

import (
	"fmt"
)

type Core struct {
	RunQ                             []*Process
	FreePages                        []*PhysicalPage
	BusyPages                        []*PhysicalPage
	AddressSpaceMax, AddressSpaceMin int
	ReqPageMax, ReqPageMin           int
	ReqWorkSetMax, ReqWorkSetMin     int
	WorkSetSizeMax, WorkSetSizeMin   int
	NReqQuantum                      int
	replacementAlgoritm ReplacementAlgorithm
	algType string
}

func (c *Core) Start(n int) {
	fmt.Println("Start system...")
	c.RunQ = make([]*Process, 0)
	c.FreePages = make([]*PhysicalPage, n)

	chooseAlg := Random(0, 100) 
	if (chooseAlg >= 50) {
		c.replacementAlgoritm = &RandomAlgorithm{}
		fmt.Println("RANDOM ALGORITHM USED")
		c.algType = "RANDOM"
	} else {
		c.replacementAlgoritm = &NRUAlgorithm{}
		fmt.Println("NRU ALGORITHM USED")
		c.algType = "NRU"
	}

	for i := 0; i < n; i++ {
		pte := &PTE{}
		physPage := &PhysicalPage{PTE: pte, Number: i}
		c.FreePages[i] = physPage
	}
	c.AddressSpaceMin = 15
	c.AddressSpaceMax = 20
	c.NReqQuantum = 10
	c.ReqPageMin = 100
	c.ReqPageMax = 150
	c.ReqWorkSetMin = 10
	c.ReqWorkSetMax = 15
	fmt.Println("System is ready to work!")
}

func (c *Core) CreateProcess() {
	process := new(Process)
	addressSpace := Random(c.AddressSpaceMin, c.AddressSpaceMax)
	process.PageTable = new(PageTable)
	process.PageTable.Entries = make([]*PTE, addressSpace) 
	for i := 0; i < addressSpace; i++ {
		pte := &PTE{}
		process.PageTable.Entries[i] = pte
	}
	reqPageCount := Random(c.ReqPageMin, c.ReqPageMax)
	process.NReq = reqPageCount 
	c.RunQ = append(c.RunQ, process)
	fmt.Println("Create process")
	fmt.Println("LEN OF PROCCESS №", len(c.RunQ), " PAGE TABLE", len(process.PageTable.Entries))
}

func (c *Core) GenerateWorkingSet(process *Process) {
	workingSetCount := Random(c.ReqWorkSetMin, c.ReqWorkSetMax)
	process.WorkingSet.PageIndexies = make([]int, workingSetCount)
	for i := 0; i < workingSetCount; i++ {
		process.WorkingSet.PageIndexies[i] = i
	}
	fmt.Println("LEN OF WORKING SET", len(process.WorkingSet.PageIndexies))
}

func (c *Core) PageFault(pageTable *PageTable, idx int) {
	var physPage **PhysicalPage
	if len(c.FreePages) > 0 {
		fmt.Println("LEN OF FREE PAGES ARRAY", len(c.FreePages))
		index := Random(0, len(c.FreePages))
		physPage = &c.FreePages[index]
		c.BusyPages = append(c.BusyPages, *physPage)
		c.removeFreePage(index)
		fmt.Println("Map free page", idx)
	} else {
		fmt.Println("LEN OF BUSY PAGES ARRAY", len(c.BusyPages))
		c.replacementAlgoritm.ReplacePage(c, &physPage)
	}
	(*physPage).PTE = pageTable.Entries[idx]
	pageTable.Entries[idx].PNN = (*physPage).Number
	(*physPage).PTE.P = true
}

func (c *Core) removeFreePage(page int) {
	c.FreePages = append(c.FreePages[:page], c.FreePages[page+1:]...)
}

func (c *Core) UpdateStat() {
	pagesToUpdate := len(c.BusyPages) / 2
	if (pagesToUpdate == 0) {
		return
	}
	pageUpdateFrom := Random(0, len(c.BusyPages) - pagesToUpdate)
	pageUpdateTo := pageUpdateFrom + pagesToUpdate
	for i := pageUpdateFrom; i <= pageUpdateTo; i++ {
		c.BusyPages[i].PTE.R = false
	}

	fmt.Println("UPDATE STATE OF PAGES", pageUpdateFrom," - ", pageUpdateTo)
}

func (c *Core) PrintAlgType() {
	fmt.Println("\t   ",c.algType, "ALGORITHM STATISTIC:")
}
