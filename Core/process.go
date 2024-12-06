package core

type Process struct {
    PageTable *PageTable
    NReq int
    WorkingSet WorkingSet
}

func (p *Process) GetPTE() *PTE {
    // 90 % із WorkingSet.Indexies , 10% - будь який номер з PageTable.PTE.lenth
    //probability := random(0, 10)
    //if (probability >= 9) {
    //   index := p.WorkingSet.PageIndexies[0]
    //    return p.PageTable.Entries[index]
    //}
    randIndex := random(0, len(p.PageTable.Entries))
    return p.PageTable.Entries[randIndex] // some index, rand
}

