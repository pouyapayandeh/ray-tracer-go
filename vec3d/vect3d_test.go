package vec3d

import (
	"fmt"
)

func ExampleVec3d_Length() {
	v := Vec3d{
		X: 0,
		Y: 1,
		Z: 0,
	}
	fmt.Println(v.Length())
	//Output:
	//1
}
func ExampleAdd() {
	v1 := Vec3d{
		X: 0,
		Y: 1,
		Z: 0,
	}
	v2 := Vec3d{
		X: 1,
		Y: 0,
		Z: 0,
	}
	v3 :=  v1.Add(&v2)
	fmt.Println(v3)
	//Output:
	//{1 1 0}
}