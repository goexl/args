package param

import (
	"github.com/goexl/args/internal/constant"
)

type Arguments struct {
	Capacity int
	Short    string
	Long     string
	Equal    string
}

func NewCreator() *Arguments {
	return &Arguments{
		Capacity: constant.Size,
		Short:    constant.Short,
		Long:     constant.Long,
		Equal:    constant.Equal,
	}
}
