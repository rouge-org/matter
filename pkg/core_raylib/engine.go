package core_raylib

import (
	"context"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/rouge-org/matter/pkg/core"
	"github.com/samber/mo"
)

type EngineImpl struct {
	config *core.EngineConfig

	ctx    context.Context
	cancel context.CancelCauseFunc

	objectRoot      core.Object
	objectCursor    core.Object
	objectFloatList []core.Object
}

func NewEngine(config *core.EngineConfig) core.Engine {
	e := new(EngineImpl)

	e.config = config

	e.ctx, e.cancel = context.WithCancelCause(context.Background())

	e.objectFloatList = make([]core.Object, 0)

	return e
}

func (e *EngineImpl) GetConfig() *core.EngineConfig {
	return e.config
}

func (e *EngineImpl) GetCtx() context.Context {
	return e.ctx
}

func (e *EngineImpl) GetObjectRoot() mo.Option[core.Object] {
	if e.objectRoot == nil {
		return mo.None[core.Object]()
	}

	return mo.Some(e.objectRoot)
}

func (e *EngineImpl) GetObjectCursor() mo.Option[core.Object] {
	if e.objectCursor == nil {
		return mo.None[core.Object]()
	}

	return mo.Some(e.objectCursor)
}

func (e *EngineImpl) ListObjectFloat() []core.Object {
	return e.objectFloatList
}

func (e *EngineImpl) runUpdate() {
	rl.InitWindow(1280, 720, "matter")
	defer rl.CloseWindow()

	rl.SetTargetFPS(e.GetConfig().Fps)
	rl.HideCursor()

	for {
		select {
		case <-e.ctx.Done():
			return

		default:
			break
		}

		if !e.doUpdate() {
			e.Stop(fmt.Errorf("window stop"))
			return
		}
	}
}

func (e *EngineImpl) Init() {
	e.objectRoot = core.NewObjectFrame()
	e.objectCursor = core.NewObjectCursor()
}

func (e *EngineImpl) Start() {
	go e.runUpdate()
}

func (e *EngineImpl) Stop(err error) {
	e.cancel(err)
}

func (e *EngineImpl) Wait() (err error) {
	<-e.ctx.Done()

	return e.ctx.Err()
}

func (e *EngineImpl) doUpdate() bool {
	if rl.WindowShouldClose() {
		return false
	}

	e.compute()

	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)
	e.render()
	rl.DrawFPS(8, 8)

	rl.EndDrawing()

	return true
}

func (e *EngineImpl) compute() {

}

func (e *EngineImpl) computeObject(ctx core.ComputeContext, object core.Object) {

}

func (e *EngineImpl) render() {
	for _, object := range e.ListObjectFloat() {
		e.renderObjectFloat(object)
	}

	e.GetObjectCursor().
		Map(func(value core.Object) (core.Object, bool) {
			e.renderObject(NewRenderContext(), value)

			return nil, false
		})

}

func (e *EngineImpl) renderObject(ctx core.RenderContext, object core.Object) {
	object.Render(ctx)
}

func (e *EngineImpl) renderObjectFloat(object core.Object) {

}
