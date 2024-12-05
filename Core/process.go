package core

type Process struct {
    PageTable *PageTable
    NReq int
    WorkingSet WorkingSet
}

func (p *Process) GetPTE() *PTE {
    // 90 % із WorkingSet.Indexies , 10% - будь який номер з PageTable.PTE.lenth
    return p.PageTable.Entries[0] // some index, rand
}

