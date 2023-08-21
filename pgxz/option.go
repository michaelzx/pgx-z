package pgxz

type OptionType uint8

type IOption interface {
	ToSql() string
}
type OptionGroup struct {
	wheres []*whereOption
	limit  *limitOption
	offset *offsetOption
	group  *groupOption
	order  *orderOption
}

func optionsToGroup(options []IOption) *OptionGroup {
	g := &OptionGroup{
		wheres: []*whereOption{},
	}
	for _, option := range options {
		switch o := option.(type) {
		case *whereOption:
			g.wheres = append(g.wheres, o)
		case *limitOption:
			g.limit = o
		case *offsetOption:
			g.offset = o
		case *groupOption:
			g.group = o
		case *orderOption:
			g.order = o
		}
	}
	return g
}
