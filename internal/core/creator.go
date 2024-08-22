package core

import (
	"github.com/goexl/args/internal/param"
)

type Creator struct {
	params    *param.Arguments
	arguments []any
}

func NewCreator() *Creator {
	return &Creator{
		params: param.NewCreator(),
	}
}

func (c *Creator) Capacity(capacity int) *Creator {
	c.params.Capacity = capacity

	return c
}

func (c *Creator) Short(placeholder string) *Creator {
	c.params.Short = placeholder

	return c
}

func (c *Creator) Long(placeholder string) *Creator {
	c.params.Long = placeholder

	return c
}

func (c *Creator) Equal(placeholder string) *Creator {
	c.params.Equal = placeholder

	return c
}

func (c *Creator) Build() *Builder {
	return NewBuilder(c.params, &c.arguments)
}
