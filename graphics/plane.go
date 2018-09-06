package graphics

import (
	"rayTracer/vec3d"
)

type Plane struct {
	P0 *vec3d.Vec3d
	N *vec3d.Vec3d
	Color vec3d.Vec3d
}
func (p Plane) Intersect(ray Ray) (bool,*vec3d.Vec3d) {
	ln := p.N.Dot(ray.Dir)
	pln := p.P0.Sub(ray.Start).Dot(p.N)
	if ln == 0 {
		return false,&vec3d.Vec3d{}
	}else {
		d:= pln/ln
		if d< 0 {
			return false,&vec3d.Vec3d{}
		}

		a:=ray.Dir.Mult(d)
		p:=ray.Start.Add(a)
		return true,p
	}
}
func (p Plane) NormalVector(d *vec3d.Vec3d)(*vec3d.Vec3d){
	//p.N.Normalize()
	return p.N
}

func (p Plane) GetColor()(*vec3d.Vec3d){
	return &p.Color
}