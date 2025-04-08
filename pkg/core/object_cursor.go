package core

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/samber/mo"
)

type ObjectCursor struct {
	id        ObjectId
	state     *ObjectState
	childList []Object
}

func NewObjectCursor() Object {
	o := new(ObjectCursor)

	o.id = NewObjectId()
	o.state = NewObjectState()
	o.childList = make([]Object, 0)

	return o
}

func (o *ObjectCursor) GetId() ObjectId {
	return o.id
}

func (o *ObjectCursor) GetState() *ObjectState {
	return o.state
}

func (o *ObjectCursor) AddChild(value Object) {
	o.childList = append(o.childList, value)
}

func (o *ObjectCursor) GetChild(id ObjectId) mo.Option[Object] {
	for _, child := range o.childList {
		if child.GetId() == id {
			return mo.Some(child)
		}
	}

	return mo.None[Object]()
}

func (o *ObjectCursor) ListChild() []Object {
	return o.childList
}

func (o *ObjectCursor) RemoveChild(id ObjectId) {
	l := make([]Object, 0)

	for _, child := range o.childList {
		if child.GetId() == id {
			continue
		}

		l = append(l, child)
	}

	o.childList = l
}

func (o *ObjectCursor) Compute(ctx ComputeContext) {

}

func (o *ObjectCursor) Render(ctx RenderContext) {
	if ctx.GetFlagRaylib() {
		rl.DrawRectangle(rl.GetMouseX()-4, rl.GetMouseY()-4, 8, 8, rl.White)
	}
}
