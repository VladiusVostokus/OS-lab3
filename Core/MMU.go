package core

type MMU struct {}

func (m *MMU) AccessPage(pageTable *PageTable, c *Core, idx int) {
	if (pageTable.Entries[idx].P == false) {
		//c.PageFault()
	}
	pageTable.Entries[idx].R = true
	writeProb := random(0, 10)
	if (writeProb >= 7) {
		pageTable.Entries[idx].M = true
	}
}