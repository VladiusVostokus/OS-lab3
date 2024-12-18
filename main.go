package main

import (
	core "OS_lab3/Core"
	"fmt"
)

func main() {
	c := &core.Core{}
	mmu := &core.MMU{}
	c.Start(25)
	c.CreateProcess()
	c.CreateProcess()
	c.CreateProcess()
	for _, proc := range c.RunQ {
		c.GenerateWorkingSet(proc)
	}

	numOfProcWorkSwitches := 10
	for i := 0; i < numOfProcWorkSwitches; i++ {
		for procIndex, _ := range c.RunQ {
			proc := c.RunQ[procIndex]
			newWorkingSetProb := core.Random(0, 100)
			if (newWorkingSetProb <= 10) {
				c.GenerateWorkingSet(proc)
				fmt.Println("GENERATE NEW WORKING SET FOR PROCESS №", procIndex + 1)
			}
	
			updateStatProb := core.Random(0, 100)
			if (updateStatProb <= 40) {
				c.UpdateStat()
			}
			for i := 0; i < c.NReqQuantum; i++ {
				index := proc.GetPTEIndex()
				mmu.AccessPage(proc.PageTable, c, index)
			}
		}
	}
	c.PrintFinalInfo(mmu.PageFaultCount, mmu.AccessCount)
}