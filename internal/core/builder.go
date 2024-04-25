package core

import (
	"github.com/goexl/args/internal/param"
	"github.com/goexl/gox"
)

type Builder struct {
	params    *param.Creator
	arguments *[]any
}

func NewBuilder(params *param.Creator, arguments *[]any) *Builder {
	return &Builder{
		params:    params,
		arguments: arguments,
	}
}

func (b *Builder) Insert(args ...any) *Builder {
	arguments := make([]any, 0, len(*b.arguments)+len(args))
	arguments = append(arguments, args...)
	arguments = append(arguments, *b.arguments...)
	b.arguments = &arguments

	return b
}

func (b *Builder) Add(args ...any) *Builder {
	*b.arguments = append(*b.arguments, args...)

	return b
}

func (b *Builder) Subcommand(subcommand string, subcommands ...string) *Builder {
	*b.arguments = append(*b.arguments, subcommand)
	for _, other := range subcommands {
		*b.arguments = append(*b.arguments, other)
	}

	return b
}

func (b *Builder) Argument(key string, value any) *Builder {
	argument := gox.StringBuilder(b.key(key), b.params.Equal, value).String()
	*b.arguments = append(*b.arguments, argument)

	return b
}

func (b *Builder) Option(key string, value any, others ...any) *Builder {
	*b.arguments = append(*b.arguments, b.key(key), gox.ToString(value))
	for _, other := range others {
		*b.arguments = append(*b.arguments, gox.ToString(other))
	}

	return b
}

func (b *Builder) Flag(flags ...string) *Builder {
	for _, flag := range flags {
		*b.arguments = append(*b.arguments, b.key(flag))
	}

	return b
}

func (b *Builder) Clone() (builder *Builder) {
	cloned := make([]any, len(*b.arguments))
	_ = copy(cloned, *b.arguments)

	builder = new(Builder)
	builder.params = b.params
	builder.arguments = &cloned

	return
}

func (b *Builder) Build() *Arguments {
	return NewArguments(b.params, b.arguments)
}

func (b *Builder) key(original string) (final string) {
	placeholder := b.params.Long
	if 1 == len(original) {
		placeholder = b.params.Short
	}
	final = gox.StringBuilder(placeholder, original).String()

	return
}
