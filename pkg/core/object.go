package core

import (
	"github.com/google/uuid"
	"github.com/samber/mo"
)

type ObjectId uuid.UUID

type Object interface {
	GetId() ObjectId
	GetState() *ObjectState
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

type ObjectState struct {
	flagVisible bool
	flagHover   bool
	flagPressed bool
	flagClick   bool
}

func NewObjectState() *ObjectState {
	return &ObjectState{
		flagVisible: true,
		flagHover:   false,
		flagPressed: false,
		flagClick:   false,
	}
}

func (s *ObjectState) Reset() {
	s.flagVisible = true
	s.flagHover = false
	s.flagPressed = false
	s.flagClick = false
}

func (s *ObjectState) GetFlagVisible() bool {
	return s.flagVisible
}
