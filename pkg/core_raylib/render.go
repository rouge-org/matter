package core_raylib

import (
	"github.com/rouge-org/matter/pkg/core"
)

type RenderContextImpl struct {
}

func NewRenderContext() core.RenderContext {
	ctx := new(RenderContextImpl)

	return ctx
}

func (ctx *RenderContextImpl) GetFlagRaylib() bool {
	return true
}
