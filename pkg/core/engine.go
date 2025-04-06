package core

import (
	"context"

	"github.com/samber/mo"
)

type Engine interface {
	GetConfig() *EngineConfig
	GetCtx() context.Context

	GetObjectRoot() mo.Option[Object]

	Init()
	Start()
	Stop(err error)
	Wait() (err error)
}

type EngineConfig struct {
	Fps int32
}
