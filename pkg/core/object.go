package core

import (
	"github.com/google/uuid"
	"github.com/samber/mo"
)

type ObjectId uuid.UUID

type Object interface {
	GetId() ObjectId
	AddChild(value Object)
	GetChild(id ObjectId) mo.Option[Object]
	ListChild() []Object
	RemoveChild(id ObjectId)

	Compute(ctx ComputeContext)
	Render(ctx RenderContext)
}

func NewObjectId() ObjectId {
	return ObjectId(uuid.New())
}
