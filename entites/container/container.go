// Package container holds a collection of 3D figures.
package container

import (
	"fmt"
	"pr1/entites/threedfigure"
)

type Container struct {
	figures []*threedfigure.ThreeDFigure
}

func New() *Container {
	return &Container{figures: make([]*threedfigure.ThreeDFigure, 0)}
}

func (c *Container) Count() int {
	return len(c.figures)
}

func (c *Container) Add(figure *threedfigure.ThreeDFigure) {
	c.figures = append(c.figures, figure)
}

func (c *Container) Remove(condition func(*threedfigure.ThreeDFigure) bool) {
	newFigures := make([]*threedfigure.ThreeDFigure, 0)
	for _, f := range c.figures {
		if !condition(f) {
			newFigures = append(newFigures, f)
		}
	}
	c.figures = newFigures
}

func (c *Container) Print() {
	if len(c.figures) == 0 {
		fmt.Println("Container is empty")
		return
	}
	for i, f := range c.figures {
		fmt.Printf("Figure %d:\n", i+1)
		f.Print()
		fmt.Println()
	}
}

func (c *Container) GetAll() []*threedfigure.ThreeDFigure {
	result := make([]*threedfigure.ThreeDFigure, len(c.figures))
	copy(result, c.figures)
	return result
}