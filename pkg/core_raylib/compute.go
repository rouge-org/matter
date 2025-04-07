package core_raylib

import (
	"github.com/rouge-org/matter/pkg/core"
)

type ComputeContextImpl struct {
}

func NewComputeContext() core.ComputeContext {
	ctx := new(ComputeContextImpl)

	return ctx
}
