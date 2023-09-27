package strategy

type Context struct {
	strategy SortStrategy
}

func NewContext(strategy SortStrategy) *Context {
	return &Context{strategy: strategy}
}

func (c *Context) SetStrategy(strategy SortStrategy) {
	c.strategy = strategy
}

func (c *Context) Sort(arr []int) []int {
	return c.strategy.Sort(arr)
}
