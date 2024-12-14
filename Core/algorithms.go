package core

import "fmt"

type ReplacementAlgorithm interface {
	ReplacePage(c *Core, page **PhysicalPage)
}

type RandomAlgorithm struct{}

func (r *RandomAlgorithm) ReplacePage(c *Core, page **PhysicalPage) {
	index := Random(0, len(c.BusyPages))
	page = &c.BusyPages[index]
	(*page).PTE.P = false
	fmt.Println("Replace page", index)
}

type NRUAlgorithm struct{}

func (nru *NRUAlgorithm) ReplacePage(c *Core, page **PhysicalPage) {
	index := Random(0, len(c.BusyPages))
	page = &c.BusyPages[index]
	(*page).PTE.P = false
	fmt.Println("Replace page", index)
	// NRU
		// Відсортувати всі фіз. сторінки
		// NRU = біт звернення + біт модифікації
		// лекція 22 NRU
}