// Package about threedfigure
package threedfigure

import (
	"pr1/entites/abstrakt"
	"pr1/entites/ball"
	"pr1/entites/cylinder"
	"pr1/entites/parallelepiped"
)

type ThreeDFigure struct {
	Ball *ball.Ball
	Cylinder *cylinder.Cylinder
	Parallelepiped *parallelepiped.Parallelepiped
}

func AddThreeDFigure() *ThreeDFigure{
	return &ThreeDFigure{}
}

func (t *ThreeDFigure) SetBall(b* ball.Ball){
	t.Ball = b
	t.Cylinder = nil
	t.Parallelepiped = nil
}

func (t *ThreeDFigure) SetParallelepiped(p *parallelepiped.Parallelepiped) {
	t.Parallelepiped = p
	t.Ball = nil
	t.Cylinder = nil
}

func (t *ThreeDFigure) SetCylinder(c *cylinder.Cylinder) {
	t.Cylinder = c
	t.Ball = nil
	t.Parallelepiped = nil
}

func (t *ThreeDFigure) GetDensity() float64 {
	if fig := t.GetAbstractFigure(); fig != nil {
		return fig.GetDensity()
	}
	return 0
}

func (t *ThreeDFigure) GetAbstractFigure() abstrakt.AbstractFigure {
	if t.Ball != nil {
		return &t.Ball.Abstrakt
	}
	if t.Cylinder != nil {
		return &t.Cylinder.Abstrakt
	}
	if t.Parallelepiped != nil {
		return &t.Parallelepiped.Abstrakt
	}
	return nil
}

func (t *ThreeDFigure) GetOwnerName() string {
	if fig := t.GetAbstractFigure(); fig != nil {
		return fig.GetOwnerName()
	}
	return ""
}

func (t *ThreeDFigure) Print() {
	if t.Ball != nil {
		t.Ball.Print()
	} else if t.Cylinder != nil {
		t.Cylinder.Print()
	} else if t.Parallelepiped != nil {
		t.Parallelepiped.Print()
	}
}