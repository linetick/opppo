// Package figurefactory creates ThreeDFigure instances from string parameters.
package figurefactory

import (
	"fmt"
	"pr1/entites/ball"
	"pr1/entites/cylinder"
	"pr1/entites/parallelepiped"
	"pr1/entites/threedfigure"
	"strconv"
)

type FigureType string

const (
	BallType           FigureType = "BALL"
	CylinderType       FigureType = "CYLINDER"
	ParallelepipedType FigureType = "PARALLELEPIPED"
)

func Create(figType FigureType, params []string) (*threedfigure.ThreeDFigure, error) {
	switch figType {
	case BallType:
		return createBall(params)
	case CylinderType:
		return createCylinder(params)
	case ParallelepipedType:
		return createParallelepiped(params)
	default:
		return nil, fmt.Errorf("unknown figure type: %s", figType)
	}
}

func createBall(params []string) (*threedfigure.ThreeDFigure, error) {
	if len(params) < 3 {
		return nil, fmt.Errorf("invalid BALL parameters. Usage: radius density owner")
	}
	radius, err := strconv.Atoi(params[0])
	if err != nil {
		return nil, fmt.Errorf("invalid radius: %v", err)
	}
	density, err := strconv.ParseFloat(params[1], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid density: %v", err)
	}
	owner := params[2]

	fig := threedfigure.AddThreeDFigure()
	fig.SetBall(ball.NewBall(radius, owner, density))
	return fig, nil
}

func createCylinder(params []string) (*threedfigure.ThreeDFigure, error) {
	if len(params) < 6 {
		return nil, fmt.Errorf("invalid CYLINDER parameters. Usage: centerX centerY radius height density owner")
	}
	centerX, _ := strconv.ParseFloat(params[0], 64)
	centerY, _ := strconv.ParseFloat(params[1], 64)
	radius, _ := strconv.ParseFloat(params[2], 64)
	height, _ := strconv.ParseFloat(params[3], 64)
	density, _ := strconv.ParseFloat(params[4], 64)
	owner := params[5]

	fig := threedfigure.AddThreeDFigure()
	fig.SetCylinder(cylinder.NewCylinder(centerX, centerY, radius, height, owner, density))
	return fig, nil
}

func createParallelepiped(params []string) (*threedfigure.ThreeDFigure, error) {
	if len(params) < 5 {
		return nil, fmt.Errorf("invalid PARALLELEPIPED parameters. Usage: edgeA edgeB edgeC density owner")
	}
	edgeA, _ := strconv.Atoi(params[0])
	edgeB, _ := strconv.Atoi(params[1])
	edgeC, _ := strconv.Atoi(params[2])
	density, _ := strconv.ParseFloat(params[3], 64)
	owner := params[4]

	fig := threedfigure.AddThreeDFigure()
	fig.SetParallelepiped(parallelepiped.NewParallelepiped(edgeA, edgeB, edgeC, owner, density))
	return fig, nil
}