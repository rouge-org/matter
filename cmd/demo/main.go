package main

import (
	"github.com/rouge-org/matter/pkg/core"
	"github.com/rouge-org/matter/pkg/core_raylib"
)

func main() {
	engine := core_raylib.NewEngine(&core.EngineConfig{
		Fps: 120,
	})

	engine.Init()
	engine.Start()
	engine.Wait()
}
