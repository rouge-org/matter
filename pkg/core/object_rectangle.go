package core

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/rouge-org/matter/pkg/data"
	"github.com/samber/mo"
)

type ObjectRectangle struct {
	id        ObjectId
	state     *ObjectState
	childList []Object

	pos  data.Vec2i
	size data.Vec2i
}

func NewObjectRectangle(pos data.Vec2i, size data.Vec2i) Object {
	o := new(ObjectRectangle)

	o.id = NewObjectId()
	o.state = NewObjectState()
	o.childList = make([]Object, 0)

	o.pos = pos
	o.size = size

	return o
}

func (o *ObjectRectangle) GetId() ObjectId {
	return o.id
}

func (o *ObjectRectangle) GetState() *ObjectState {
	return o.state
}

func (o *ObjectRectangle) AddChild(value Object) {
	o.childList = append(o.childList, value)
}

func (o *ObjectRectangle) GetChild(id ObjectId) mo.Option[Object] {
	for _, child := range o.childList {
		if child.GetId() == id {
			return mo.Some(child)
		}
	}

	return mo.None[Object]()
}

func (o *ObjectRectangle) ListChild() []Object {
	return o.childList
}

func (o *ObjectRectangle) RemoveChild(id ObjectId) {
	l := make([]Object, 0)

	for _, child := range o.childList {
		if child.GetId() == id {
			continue
		}

		l = append(l, child)
	}

	o.childList = l
}

func (o *ObjectRectangle) Compute(ctx ComputeContext) {

}

func (o *ObjectRectangle) Render(ctx RenderContext) {
	rl.DrawRectangle(
		o.pos.GetX(),
		o.pos.GetY(),
		o.size.GetX(),
		o.size.GetY(),
		rl.DarkBlue)
}
