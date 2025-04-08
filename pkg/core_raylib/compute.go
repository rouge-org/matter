package core_raylib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/rouge-org/matter/pkg/core"
	"github.com/rouge-org/matter/pkg/data"
)

type ComputeContextImpl struct {
	cursorPos data.Vec2i
}

func NewComputeContext() core.ComputeContext {
	ctx := new(ComputeContextImpl)

	ctx.cursorPos = data.NewVec2i(rl.GetMouseX(), rl.GetMouseY())

	return ctx
}

func (ctx *ComputeContextImpl) GetFlagRaylib() bool {
	return true
}
