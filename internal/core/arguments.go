package core

import (
	"strings"

	"github.com/goexl/args/internal/constant"
	"github.com/goexl/args/internal/param"
	"github.com/goexl/gox"
)

type Arguments struct {
	params    *param.Creator
	arguments *[]any

	_ gox.CannotCopy
}

func NewArguments(params *param.Creator, arguments *[]any) *Arguments {
	return &Arguments{
		params:    params,
		arguments: arguments,
	}
}

func (a *Arguments) String() (arguments []string) {
	arguments = make([]string, 0, len(*a.arguments))
	for _, argument := range *a.arguments {
		arguments = append(arguments, gox.ToString(argument))
	}

	return
}

func (a *Arguments) Cli() string {
	return strings.Join(a.String(), constant.Space)
}

func (a *Arguments) Rebuild() *Builder {
	return NewBuilder(a.params, a.arguments)
}
