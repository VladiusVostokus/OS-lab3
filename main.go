package main

import (
	core "OS_lab3/Core"
)

func main() {
	c := &core.Core{}
	mmu := &core.MMU{}
	c.Start(100)
	c.CreateProcess()
	proc := c.GetProcess()
	c.GenerateWorkingSet(proc)
	//for proc, _ := range c.
	for i := 0; i < c.NReqQuantum; i++ {
		index := proc.GetPTEIndex()
		mmu.AccessPage(proc.PageTable, c, index)
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