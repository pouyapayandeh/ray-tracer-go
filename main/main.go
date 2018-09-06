package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"rayTracer/graphics"
	"rayTracer/vec3d"
)

//

func main() {
	w := 400
	h := 400
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	eye := &vec3d.Vec3d{
		X:-200,
		Y:0,
		Z:0,
	}
	light := &vec3d.Vec3d{
		X:0,
		Y:0,
		Z:-200,
	}

	var objects  []graphics.Renderable
	objects = append(objects, graphics.Sphere{
		R:20,
		C:&vec3d.Vec3d{100,0,0},
		Color:vec3d.Vec3d{255,0,0},
	})
	objects = append(objects, graphics.Sphere{
			R:100,
			C:&vec3d.Vec3d{400,0,0},
			Color:vec3d.Vec3d{255,255,0},
	})
	objects = append(objects, graphics.Sphere{
			R:300,
			C:&vec3d.Vec3d{800,0,5},
			Color:vec3d.Vec3d{0,0,255},
	})
	objects = append(objects, graphics.Plane{
		N:&vec3d.Vec3d{0,0,-1},
		P0:&vec3d.Vec3d{0,0,320},
		Color:vec3d.Vec3d{255,255,255},
	})
	for x := 0; x< w ;x++  {
		for y := 0; y< h ;y++  {
			yy := float64(-(h / 2) + y)
			zz := float64(-(w/2)+x)
			pixel := vec3d.Vec3d{0,zz,yy,}
			dir := pixel.Sub(eye)
			dir.Normalize()

			ray := graphics.Ray{Start:eye,Dir:dir}

			minDist := math.Inf(+1)
			minIndex := -1
			var hPoint *vec3d.Vec3d = nil
			for i := 0; i < len(objects); i++ {
				hit , point := objects[i].Intersect(ray)
				if hit == true{
					dist:=graphics.Distance(point,eye)
					//fmt.Println(point)
					if minDist > dist{
						minDist = dist
						minIndex = i
						hPoint =point
					}
				}
			}
			inShadow := false
			if hPoint != nil{
				bias := objects[minIndex].NormalVector(hPoint)
				bias.Normalize()

				shadowDir := light.Sub(hPoint)
				shadowDir.Normalize()

				shadowRay := graphics.Ray{Start:hPoint.Add(bias.Mult(1e-5)),Dir:shadowDir}
				for i := 0; i < len(objects); i++ {
					hit , pShadow := objects[i].Intersect(shadowRay)
					//
					if hit && graphics.Distance(light,hPoint) - graphics.Distance(light,pShadow) > -0.00001 {
						fmt.Println(pShadow ,hPoint )
						inShadow = true
						break
					}
				}
			}else
			{
				inShadow = true
			}
			if !inShadow{
				//fmt.Println("Color ",x," ",y)
				//d := graphics.Distance(light,hPoint)
				c := objects[minIndex].GetColor()
				img.Set(x, y, color.RGBA{uint8(c.X), uint8(c.Y), uint8(c.Z), 255})

			}else {
				img.Set(x, y, color.RGBA{0, 0, 0, 255})
			}



		}
	}
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
