package data

type Vec2i struct {
	x int32
	y int32
}

func NewVec2i(x int32, y int32) Vec2i {
	return Vec2i{
		x, y,
	}
}

func (v *Vec2i) GetX() int32 {
	return v.x
}

func (v *Vec2i) GetY() int32 {
	return v.y
}
