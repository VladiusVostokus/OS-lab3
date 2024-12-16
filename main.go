package main

import (
	core "OS_lab3/Core"
	"fmt"
)

func main() {
	c := &core.Core{}
	mmu := &core.MMU{}
	c.Start(100)
	c.CreateProcess()
	c.CreateProcess()
	c.CreateProcess()
	for _, proc := range c.RunQ {
		c.GenerateWorkingSet(proc)
	}

	numOfProcWorkSwitches := 3
	for i := 0; i < numOfProcWorkSwitches; i++ {
		for procIndex, _ := range c.RunQ {
			proc := c.RunQ[procIndex]
			newWorkingSetProb := core.Random(0, 100)
			if (newWorkingSetProb <= 10) {
				c.GenerateWorkingSet(proc)
				fmt.Println("GENERATE NEW WORKING SET FOR PROCESS №", procIndex + 1)
			}
	
			updateStatProb := core.Random(0, 100)
			if (updateStatProb <= 20) {
				c.UpdateStat()
			}
			for i := 0; i < c.NReqQuantum; i++ {
				index := proc.GetPTEIndex()
				mmu.AccessPage(proc.PageTable, c, index)
			}
		}
	}

	fmt.Println("Total count of accesses to pages:", mmu.AccessCount)
	fmt.Println("Total count of page faults:", mmu.PageFaultCount)
	fmt.Println("Total page fault ratio =", float32(mmu.PageFaultCount) / float32(mmu.AccessCount) * 100)
	/*
	сounter := 0
	for proc, _ := range c.RunQ {
	// прокручувати процеси із списку по колу
	counter++
	десь тут згенерувати новий робочий набір для процесу
	if (counter % 10 == 0) {
		
	}
		counter - кожну 10 ітерацію оновлює стату(скидає біти звернення R для частини сторінок, напр для наступних 10)
		for i := 0; i < proc.NReq; i++ {
			index := proc.GetPTEIndex()
			mmu.AccessPage(proc.PageTable, c, index)
		}
	}
	*/
	// додати ф.-процес ядра для статистики - скадає біти звернення(R)
	// 
}