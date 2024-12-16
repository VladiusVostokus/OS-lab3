package core

import "fmt"

type MMU struct{
	AccessCount, PageFaultCount int
}

func (m *MMU) AccessPage(pageTable *PageTable, c *Core, idx int) {
	fmt.Println("Trying to access virtual page num =", idx)
	m.AccessCount++
	if pageTable.Entries[idx].P == false {
		m.PageFaultCount++
		fmt.Println("Page fault: vitrual page ", idx, " is not mammped")
		c.PageFault(pageTable, idx)
	}
	fmt.Println("Regular access to page:", idx)
	pageTable.Entries[idx].R = true
	writeProb := Random(0, 100)
	if writeProb >= 70 {
		pageTable.Entries[idx].M = true
		fmt.Println("Write to page", idx)
	} else {
		fmt.Println("Read from page", idx)
	}
}
