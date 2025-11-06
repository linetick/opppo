package parallelepiped

import (
	"fmt"
	"pr1/entites/abstrakt"
)

type Parallelepiped struct {
	abstrakt.Abstrakt
	Rebro1 int
	Rebro2 int
	Rebro3 int
}

func NewParallelepiped(rebro1, rebro2, rebro3 int, ownerName string, density float64) *Parallelepiped {
	return &Parallelepiped{
		Abstrakt: abstrakt.BaseAbstrakt(density, ownerName),
		Rebro1: rebro1,
		Rebro2: rebro2,
		Rebro3: rebro3,
	}
}

func (p *Parallelepiped) Print() {
	fmt.Println("=Parallelepiped=")
	p.Abstrakt.Print()
	fmt.Printf("Rebro1  %d\n", p.Rebro1)
	fmt.Printf("Rebro2  %d\n", p.Rebro2)
	fmt.Printf("Rebro3  %d\n", p.Rebro3)
	fmt.Println("=====")
}