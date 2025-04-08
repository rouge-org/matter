package main

import (
	"github.com/rouge-org/matter/pkg/core"
	"github.com/rouge-org/matter/pkg/core_raylib"
	"github.com/rouge-org/matter/pkg/data"
)

func main() {
	engine := core_raylib.NewEngine(&core.EngineConfig{
		Fps: 120,
	})

	engine.Init()

	engine.AddObjectFloat(core.NewObjectRectangle(
		data.NewVec2i(64, 64),
		data.NewVec2i(720, 360)))

	engine.Start()
	engine.Wait()
}
