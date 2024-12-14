package core

import "fmt"

type Process struct {
	PageTable  *PageTable
	NReq       int
	WorkingSet WorkingSet
}

func (p *Process) GetPTEIndex() int {
	probability := Random(0, 10)
	if probability <= 9 {
		randIndex := Random(0, len(p.WorkingSet.PageIndexies))
		fmt.Println("Get random PTE form working set", randIndex)
		return randIndex
	}
	randIndex := Random(0, len(p.PageTable.Entries))
	fmt.Println("Get random PTE form page table", randIndex)
	return randIndex // some index, rand
}
