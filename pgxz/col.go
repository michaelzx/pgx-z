package pgxz

type KyeMap map[string]struct{}
type ICol interface {

	// Col的孩子 来实现

	TableName() string

	// Col 来实现

	Init(km KyeMap)
	Set(fn string, v any)
	SetMust(fn string, v any)
	setMustKey(fn string, v any)
	Keys() []string
	HasKey(k string) bool
	IsSet(k string) bool
	Mapping() map[string]any
}

type Col struct {
	mapping map[string]any
	keyMap  KyeMap
}

func (c *Col) Init(km KyeMap) {
	c.mapping = make(map[string]any)
	c.keyMap = km
}

func (c *Col) Set(fn string, v any) {
	c.setMustKey(fn, v)
}

// SetMust When setting the value, the value must not nil
func (c *Col) SetMust(fn string, v any) {
	if v != nil {
		c.setMustKey(fn, v)
	}
}

// SetMustKey When setting the value, the key must be within the range of keys
func (c *Col) setMustKey(fn string, v any) {
	if c.HasKey(fn) {
		c.mapping[fn] = v
	}
}

func (c *Col) HasKey(key string) bool {
	_, exists := c.keyMap[key]
	return exists
}
func (c *Col) IsSet(key string) bool {
	_, exists := c.mapping[key]
	return exists
}

func (c *Col) Keys() []string {
	list := make([]string, 0)
	for k, _ := range c.keyMap {
		list = append(list, k)
	}
	return list
}

func (c *Col) Mapping() map[string]any {
	return c.mapping
}
