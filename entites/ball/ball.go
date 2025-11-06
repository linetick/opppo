package ball

import (
	"fmt"
	"pr1/entites/abstrakt"
)


type Ball struct {
	abstrakt.Abstrakt
	Radius int
}

func NewBall(radius int, ownerName string, density float64) *Ball {
	return &Ball{
		Abstrakt: abstrakt.BaseAbstrakt(density, ownerName),
		Radius:           radius,
	}
}

func (b *Ball) Print() {
	fmt.Println("=Ball=")
	b.Abstrakt.Print()
	fmt.Printf("Radius  %d\n", b.Radius)
	fmt.Println("=====")
}