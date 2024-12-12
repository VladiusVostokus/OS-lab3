package core

import "fmt"

type MMU struct {}

func (m *MMU) AccessPage(pageTable *PageTable, c *Core, idx int) {
	fmt.Println("Trying to access virtual page num =",idx)
	if (pageTable.Entries[idx].P == false) {
		fmt.Println("Page fault: vitrual page ", idx, " is not mammped")
		c.PageFault(pageTable, idx)
	}
	fmt.Println("Regular access to page:", idx)
	pageTable.Entries[idx].R = true
	writeProb := random(0, 10)
	if (writeProb >= 7) {
		pageTable.Entries[idx].M = true
		fmt.Println("Write to page", idx)
	} else {
		fmt.Println("Read from page", idx)
	}
}