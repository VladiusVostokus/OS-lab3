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
	class1 := make([]int,0)
	class2 := make([]int,0)
	class3 := make([]int,0)
	for i, v := range c.BusyPages {
		if (v.PTE.R == false && v.PTE.M == false) {
			page = &c.BusyPages[i]
			(*page).PTE.P = false
			fmt.Println("Replace page", i)
			return
		} else {
			if (v.PTE.R == false && v.PTE.M == true) {
				class1 = append(class1, i)
			} else {
				if (v.PTE.R == true && v.PTE.M == false) {
					class2 = append(class2, i)
				} else {
					if (v.PTE.R == true && v.PTE.M == true) {
						class3 = append(class3, i)
					}
				}
			}
		}
	}

	var pageIndex int
	if (len(class1) > 0) {
		index := Random(0, len(class1) - 1)
		pageIndex = class1[index]
	} else {
		if (len(class2) > 0) {
			index := Random(0, len(class2) - 1)
			pageIndex = class2[index]
		} else {
			if (len(class3) > 0) {
				index := Random(0, len(class3) - 1)
				pageIndex = class3[index]
			}
		}
	}
	page = &c.BusyPages[pageIndex]
	(*page).PTE.P = false
	fmt.Println("Replace page", pageIndex)
	// NRU
		// Відсортувати всі фіз. сторінки
		// NRU = біт звернення + біт модифікації
		// лекція 22 NRU
}