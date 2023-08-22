package pgxz

func Page[T any](db *PgDb, selectSql, filterSql, orderSql string, pageParams IPageParams) (*PageResult[T], error) {
	p := &PageResult[T]{
		Pagination: NewPagination(pageParams),
		List:       make([]T, 0),
	}
	err := p.doQuery(db, selectSql, filterSql, orderSql, pageParams)
	if err != nil {
		return nil, err
	}
	return p, nil
}
