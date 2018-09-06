package graphics

import "rayTracer/vec3d"

func Distance (p1 *vec3d.Vec3d ,p2 *vec3d.Vec3d) float64{
	p3 := p1.Sub(p2)
	return p3.Length()
}


type Renderable  interface  {
	NormalVector(d *vec3d.Vec3d)(*vec3d.Vec3d)
	Intersect(ray Ray) (bool,*vec3d.Vec3d)
	GetColor()(*vec3d.Vec3d)
}