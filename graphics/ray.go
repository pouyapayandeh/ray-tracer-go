package graphics

import (
	"rayTracer/vec3d"
)

type Ray struct {
	Start *vec3d.Vec3d
	Dir   *vec3d.Vec3d
}

//func (r Ray) String() string  {
//	return fmt.Sprintf("%+v %+v " ,r.Start , r.Dir)
//}
