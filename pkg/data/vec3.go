package data

type Vec3i struct {
	x int32
	y int32
	z int32
}

func NewVec3i(x int32, y int32, z int32) Vec3i {
	return Vec3i{
		x, y, z,
	}
}

func (v *Vec3i) GetX() int32 {
	return v.x
}

func (v *Vec3i) GetY() int32 {
	return v.y
}

func (v *Vec3i) GetZ() int32 {
	return v.z
}
