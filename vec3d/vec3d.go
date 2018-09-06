package vec3d

import "math"

type Vec3d struct {
	X float64
	Y float64
	Z float64
}
func (v Vec3d) Length() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y +v.Z*v.Z)
}
func (v1 *Vec3d) Add(v2 *Vec3d) *Vec3d {
	return &Vec3d{
		X: v1.X +v2.X,
		Y: v1.Y +v2.Y,
		Z: v1.Z +v2.Z,
	}
}

func (v1 *Vec3d) Sub(v2 *Vec3d) *Vec3d {
	return &Vec3d{
		X: v1.X -v2.X,
		Y: v1.Y -v2.Y,
		Z: v1.Z -v2.Z,
	}
}

func (v1 *Vec3d) Mult(r float64) *Vec3d {
	 	return &Vec3d{
		X: v1.X * r,
		Y: v1.Y * r,
		Z: v1.Z * r,
	}
}



func (v *Vec3d) Normalize(){
	len_ := v.Length()
	if len_ > 0{
		v.X /=len_
		v.Y /=len_
		v.Z /=len_
	}
}

func (v1 *Vec3d) Dot(v2 *Vec3d ) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func (v1 *Vec3d) Cross(v2 *Vec3d ) *Vec3d {
	return &Vec3d{
		X: v1.Y* v2.Z - v1.Z* v2.Y,
		Y: v1.Z* v2.X - v1.X* v2.Z,
		Z: v1.X* v2.Y - v1.Y* v2.X,
	}
}