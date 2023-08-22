package pgxz

var _ IPageParams = PageParams{}

type IPageParams interface {
	GetPageNum() int64
	GetPageSize() int64
}

type PageParams struct {
	PageNum  int64 `json:"page_num,omitempty"`
	PageSize int64 `json:"page_size,omitempty"`
}

func (p PageParams) GetPageNum() int64 {
	return p.PageNum
}

func (p PageParams) GetPageSize() int64 {
	return p.PageSize
}
