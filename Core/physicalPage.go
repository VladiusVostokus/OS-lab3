package core

type PhyіsicalPage struct {
	PTE *PTE
}

type freePages []*PhyіsicalPage

type busyPages []*PhyіsicalPage