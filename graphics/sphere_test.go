package graphics

import (
	"fmt"
	"rayTracer/vec3d"
)

func ExampleSphere_Intersect() {
	s := Sphere{
		C:vec3d.Vec3d{
			X:-10,
			Y:0,
			Z : 0,
		},
		R:10,
	}
	ray := Ray{
		Start:&vec3d.Vec3d{
			X:1,
			Y:0,
			Z:0,
		},
		Dir:&vec3d.Vec3d{
			X:1,
			Y:0,
			Z:0,
		},
	}
	h , _ := s.Intersect(ray)
	fmt.Println(h)
	//Output:
	//false
}
