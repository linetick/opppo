package cylinder

import (
	"fmt"
	"pr1/entites/abstrakt"
)

type Point struct {
	X float64
	Y float64
}

type Cylinder struct {
	abstrakt.Abstrakt
	Center Point
	Radius float64
	Height float64
}

func NewCylinder(centerX, centerY, radius, height float64, name string, density float64) *Cylinder {
	return &Cylinder{
		Abstrakt: abstrakt.BaseAbstrakt(density, name),
		Center:          Point{X: centerX, Y: centerY},
		Radius:          radius,
		Height:          height,
	}
}

func (c *Cylinder) Print(){
	fmt.Println("=== Cylinder ===")
	c.Abstrakt.Print()
	fmt.Printf("Center: (%.2f, %.2f)\n", c.Center.X, c.Center.Y)
	fmt.Printf("Radius: %.2f\n", c.Radius)
	fmt.Printf("Height: %.2f\n", c.Height)
	fmt.Println("================")
}


