package args

import (
	"github.com/goexl/args/internal/core"
)

var _ = New

// New 创建新的参数
func New() *core.Creator {
	return core.NewCreator()
}
