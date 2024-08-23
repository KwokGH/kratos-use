package dbtool

// Pager paging setting
type Pager struct {
	Page     int
	PageSize int
}

func NewPager(page int32, pageSize int32) Pager {
	if pageSize > 200 {
		pageSize = 200
	}
	return Pager{
		Page:     int(page),
		PageSize: int(pageSize),
	}
}

// Offset pager to offset
func (p Pager) Offset() (int, int) {
	return (p.Page - 1) * p.PageSize, p.PageSize
}

// Range pager to [start, end)
func (p Pager) Range() (int, int) {
	return (p.Page - 1) * p.PageSize, p.Page*p.PageSize - 1
}

// Enable enable paging if page and pageSize is not zero
func (p Pager) Enable() bool {
	return p.Page != 0 && p.PageSize != 0
}

// NoPager do not paging
var NoPager = Pager{Page: 0, PageSize: 0}
