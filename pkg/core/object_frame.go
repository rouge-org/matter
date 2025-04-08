package core

import (
	"github.com/samber/mo"
)

type ObjectFrame struct {
	id        ObjectId
	state     *ObjectState
	childList []Object
}

func NewObjectFrame() Object {
	o := new(ObjectFrame)

	o.id = NewObjectId()
	o.state = NewObjectState()
	o.childList = make([]Object, 0)

	return o
}

func (o *ObjectFrame) GetId() ObjectId {
	return o.id
}

func (o *ObjectFrame) GetState() *ObjectState {
	return o.state
}

func (o *ObjectFrame) AddChild(value Object) {
	o.childList = append(o.childList, value)
}

func (o *ObjectFrame) GetChild(id ObjectId) mo.Option[Object] {
	for _, child := range o.childList {
		if child.GetId() == id {
			return mo.Some(child)
		}
	}

	return mo.None[Object]()
}

func (o *ObjectFrame) ListChild() []Object {
	return o.childList
}

func (o *ObjectFrame) RemoveChild(id ObjectId) {
	l := make([]Object, 0)

	for _, child := range o.childList {
		if child.GetId() == id {
			continue
		}

		l = append(l, child)
	}

	o.childList = l
}

func (o *ObjectFrame) Compute(ctx ComputeContext) {

}

func (o *ObjectFrame) Render(ctx RenderContext) {

}
