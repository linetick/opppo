// Package abstrakt defines abstract base types for 3D figures.
package abstrakt

import "fmt"

type AbstractFigure interface{
	GetDensity() float64
	GetOwnerName() string
	Print()
}

type Abstrakt struct {
	Density float64
	Name    string
}

func BaseAbstrakt(density float64, name string) Abstrakt {
	return Abstrakt{
		Density: density,
		Name:    name,
	}
}

func (b *Abstrakt) GetOwnerName() string {
	return b.Name
}

func (b *Abstrakt) GetDensity() float64 {
	return b.Density
}

func (a *Abstrakt) Print() {
	fmt.Println("Density:", a.Density)
	fmt.Println("Owner:", a.Name)
}