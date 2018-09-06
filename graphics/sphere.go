package graphics

import (
	"math"
	"rayTracer/vec3d"
)

type Sphere struct {
	C *vec3d.Vec3d
	R float64
	Color vec3d.Vec3d
}
func (s Sphere) Intersect(ray Ray) (bool,*vec3d.Vec3d){

	oc := ray.Start.Sub(s.C)
	loc := oc.Dot(ray.Dir)
	delta := (loc*loc) - oc.Dot(oc) + s.R*s.R
	if delta < 0{
		return false,&vec3d.Vec3d{}
	}else{
		d1 := -loc + math.Sqrt(delta)
		d2 := -loc - math.Sqrt(delta)
		if d1 < 0 && d2 > 0{
			a:=ray.Dir.Mult(d2)
			p:=ray.Start.Add(a)
			return true,p
		}
		if d1 > 0 && d2 < 0{
			a:=ray.Dir.Mult(d1)
			p:=ray.Start.Add(a)
			return true,p
		}
		if d1 < 0 && d2 < 0{
			return false,&vec3d.Vec3d{}
		}
		if d1 <= d2{
			a:=ray.Dir.Mult(d1)
			p:=ray.Start.Add(a)
			return true,p
		}else{
			a:=ray.Dir.Mult(d2)
			p:=ray.Start.Add(a)
			return true,p
		}
	}
}

func (s Sphere) NormalVector(d *vec3d.Vec3d)(*vec3d.Vec3d){
	return d.Sub(s.C)
}

func (s Sphere) GetColor()(*vec3d.Vec3d){
	return &s.Color
}