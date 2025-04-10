package core

import (
	"context"

	"github.com/samber/mo"
)

type Engine interface {
	GetConfig() *EngineConfig
	GetCtx() context.Context

	GetObjectRoot() mo.Option[Object]
	GetObjectCursor() mo.Option[Object]
	ListObjectFloat() []Object
	AddObjectFloat(value Object)

	Init()
	Start()
	Stop(err error)
	Wait() (err error)
}

type EngineConfig struct {
	Fps int32
}
