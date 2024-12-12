package main

import (
	core "OS_lab3/Core"
)

func main() {
	c := &core.Core{}
	mmu := &core.MMU{}
	c.Start(1000)
	c.CreateProcess()
	proc := c.GetProcess()
	c.GenerateWorkingSet(proc)
	for i := 0; i < proc.NReq; i++ {
		index := proc.GetPTEIndex()
		mmu.AccessPage(proc.PageTable, c, index)
	}
}