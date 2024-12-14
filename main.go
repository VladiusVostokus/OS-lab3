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
	for procIndex, _ := range c.RunQ {
		proc := c.RunQ[procIndex]
		newWorkingSetProb := core.Random(0, 10)
		if (newWorkingSetProb < 1) {
			c.GenerateWorkingSet(proc)
			fmt.Println("GENERATE NEW WORKING SET FOR PROCESS №", procIndex + 1)
		}

		updateStatProb := core.Random(0, 10)
		if (updateStatProb < 2) {
			c.UpdateStat()
		}
		for i := 0; i < c.NReqQuantum; i++ {
			index := proc.GetPTEIndex()
			mmu.AccessPage(proc.PageTable, c, index)
		}
	}
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