package pgxz

type ICol interface {
	Init()
	Set(fn string, v any)
	SetIfNotNil(fn string, v any)
	TableName() string
	Keys() []string
	Mapping() map[string]any
}

type Col struct {
	mapping map[string]any
}

func (c *Col) Init() {
	c.mapping = make(map[string]any)
}
func (c *Col) SetIfNotNil(fn string, v any) {
	if v != nil {
		c.mapping[fn] = v
	}
}

func (c *Col) Set(fn string, v any) {
	c.mapping[fn] = v
}
func (c *Col) Mapping() map[string]any {
	return c.mapping
}
