package pgxz

type Pagination struct {
	PageNum     int64 // 第几页
	PageSize    int64 // 每页几条
	PageTotal   int64 // 总共几页
	Total       int64 // 总共几条
	IsFirstPage bool  // 是否是第一页
	IsLastPage  bool  // 是否是最后一页
}

func NewPagination(pageParams IPageParams) Pagination {
	p := Pagination{
		PageNum:  pageParams.GetPageNum(),
		PageSize: pageParams.GetPageSize(),
	}
	if p.PageNum <= 0 {
		p.PageNum = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	if p.PageSize > 100 {
		p.PageSize = 100
	}
	return p
}

func (p *Pagination) Compute(total int64) {
	p.Total = total
	// 一共几页
	p.PageTotal = p.Total / p.PageSize // 可能是0
	if p.Total%p.PageSize > 0 {
		p.PageTotal = p.PageTotal + 1
	}
	if p.PageNum > p.PageTotal {
		p.PageNum = p.PageTotal // 可能是0
		if p.PageNum == 0 {
			p.PageNum = 1
		}
	}
	if p.PageTotal == 0 {
		p.IsFirstPage = true
		p.IsLastPage = true
	} else {
		if p.PageNum == 1 {
			p.IsFirstPage = true
		}
		if p.PageNum == p.PageTotal {
			p.IsLastPage = true
		}
	}
}
func (p *Pagination) GetSkipRows() int64 {
	return p.PageSize * (p.PageNum - 1)
}
