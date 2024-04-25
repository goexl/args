package param

import (
	"github.com/goexl/args/internal/constant"
)

type Creator struct {
	Capacity int
	Short    string
	Long     string
	Equal    string
}

func NewCreator() *Creator {
	return &Creator{
		Capacity: constant.Size,
		Short:    constant.Short,
		Long:     constant.Long,
		Equal:    constant.Equal,
	}
}
