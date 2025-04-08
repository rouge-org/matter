package core_raylib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/rouge-org/matter/pkg/core"
	"github.com/rouge-org/matter/pkg/data"
)

type RenderContextImpl struct {
	cursorPos data.Vec2i
}

func NewRenderContext() core.RenderContext {
	ctx := new(RenderContextImpl)

	ctx.cursorPos = data.NewVec2i(rl.GetMouseX(), rl.GetMouseY())

	return ctx
}

func (ctx *RenderContextImpl) GetFlagRaylib() bool {
	return true
}
