// Package process handles command parsing and file processing.
package process

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pr1/entites/ball"
	"pr1/entites/cylinder"
	"pr1/entites/parallelepiped"
	"pr1/entites/threedfigure"
	"strconv"
	"strings"
)

type Container struct {
	figures []*threedfigure.ThreeDFigure
}

func NewContainer() *Container {
	return &Container{
		figures: make([]*threedfigure.ThreeDFigure, 0),
	}
}

func (c *Container) Count() int {
	return len(c.figures)
}

func (c *Container) Add(figure *threedfigure.ThreeDFigure) {
	c.figures = append(c.figures, figure)
}

func (c *Container) Remove(condition func(*threedfigure.ThreeDFigure) bool) {
	newFigures := make([]*threedfigure.ThreeDFigure, 0)
	for _, figure := range c.figures {
		if !condition(figure) {
			newFigures = append(newFigures, figure)
		}
	}
	c.figures = newFigures
}

func (c *Container) Print() {
	if len(c.figures) == 0 {
		fmt.Println("Container is empty")
		return
	}

	for i, figure := range c.figures {
		fmt.Printf("Figure %d:\n", i+1)
		figure.Print()
		fmt.Println()
	}
}

func ParseCommand(line string) (string, []string) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return "", nil
	}
	return parts[0], parts[1:]
}

func ProcessADD(args []string, container *Container) error {
	if len(args) < 2 {
		return fmt.Errorf("invalid ADD command format. Usage: ADD FIGURE_TYPE params...")
	}

	figureType := strings.ToUpper(args[0])
	params := args[1:]

	threeDFigure := threedfigure.AddThreeDFigure()

	switch figureType {
	case "BALL":
		return addBall(params, threeDFigure, container)
	case "CYLINDER":
		return addCylinder(params, threeDFigure, container)
	case "PARALLELEPIPED":
		return addParallelepiped(params, threeDFigure, container)
	default:
		return fmt.Errorf("unknown figure type: %s", figureType)
	}
}

func addBall(params []string, figure *threedfigure.ThreeDFigure, container *Container) error {
	if len(params) < 3 {
		return fmt.Errorf("invalid BALL parameters. Usage: ADD BALL radius density owner")
	}

	radius, err := strconv.Atoi(params[0])
	if err != nil {
		return fmt.Errorf("invalid radius: %v", err)
	}
	
	density, err := strconv.ParseFloat(params[1], 64)
	if err != nil {
		return fmt.Errorf("invalid density: %v", err)
	}
	
	ownerName := params[2]
	ball := ball.NewBall(radius, ownerName, density)
	figure.SetBall(ball)
	container.Add(figure)
	
	return nil
}

func addCylinder(params []string, figure *threedfigure.ThreeDFigure, container *Container) error {
	if len(params) < 6 {
		return fmt.Errorf("invalid CYLINDER parameters. Usage: ADD CYLINDER centerX centerY radius height density owner")
	}

	centerX, err := strconv.ParseFloat(params[0], 64)
	if err != nil {
		return fmt.Errorf("invalid center X: %v", err)
	}
	
	centerY, err := strconv.ParseFloat(params[1], 64)
	if err != nil {
		return fmt.Errorf("invalid center Y: %v", err)
	}
	
	radius, err := strconv.ParseFloat(params[2], 64)
	if err != nil {
		return fmt.Errorf("invalid radius: %v", err)
	}
	
	height, err := strconv.ParseFloat(params[3], 64)
	if err != nil {
		return fmt.Errorf("invalid height: %v", err)
	}
	
	density, err := strconv.ParseFloat(params[4], 64)
	if err != nil {
		return fmt.Errorf("invalid density: %v", err)
	}
	
	ownerName := params[5]
	cylinder := cylinder.NewCylinder(centerX, centerY, radius, height, ownerName, density)
	figure.SetCylinder(cylinder)
	container.Add(figure)
	
	return nil
}

func addParallelepiped(params []string, figure *threedfigure.ThreeDFigure, container *Container) error {
	if len(params) < 5 {
		return fmt.Errorf("invalid PARALLELEPIPED parameters. Usage: ADD PARALLELEPIPED edgeA edgeB edgeC density owner")
	}

	edgeA, err := strconv.Atoi(params[0])
	if err != nil {
		return fmt.Errorf("invalid edge A: %v", err)
	}
	
	edgeB, err := strconv.Atoi(params[1])
	if err != nil {
		return fmt.Errorf("invalid edge B: %v", err)
	}
	
	edgeC, err := strconv.Atoi(params[2])
	if err != nil {
		return fmt.Errorf("invalid edge C: %v", err)
	}
	
	density, err := strconv.ParseFloat(params[3], 64)
	if err != nil {
		return fmt.Errorf("invalid density: %v", err)
	}
	
	ownerName := params[4]
	parallelepiped := parallelepiped.NewParallelepiped(edgeA, edgeB, edgeC, ownerName, density)
	figure.SetParallelepiped(parallelepiped)
	container.Add(figure)
	
	return nil
}

func ProcessREM(args []string, container *Container) error {
	if len(args) < 3 {
		return fmt.Errorf("invalid REM command format. Usage: REM field operator value")
	}

	field := args[0]
	operator := args[1]
	valueStr := strings.Join(args[2:], " ") 
	switch strings.ToLower(field) {
	case "owner":
		return processOwnerCondition(operator, valueStr, container)
	case "density":
		return processDensityCondition(operator, valueStr, container)
	case "radius":
		return processRadiusCondition(operator, valueStr, container)
	case "rebro1", "edgea":
		return processEdgeCondition(operator, valueStr, container, "edgeA")
	case "rebro2", "edgeb":
		return processEdgeCondition(operator, valueStr, container, "edgeB")
	case "rebro3", "edgec":
		return processEdgeCondition(operator, valueStr, container, "edgeC")
	case "height":
		return processHeightCondition(operator, valueStr, container)
	default:
		return fmt.Errorf("unknown field: %s", field)
	}
}

func processOwnerCondition(operator, valueStr string, container *Container) error {
	switch operator {
	case "=", "==":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.GetOwnerName() == valueStr
		})
	case "!=":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.GetOwnerName() != valueStr
		})
	case "contains", "CONTAINS":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return strings.Contains(strings.ToLower(f.GetOwnerName()), strings.ToLower(valueStr))
		})
	default:
		return fmt.Errorf("invalid operator for owner field: %s", operator)
	}
	return nil
}

func processDensityCondition(operator, valueStr string, container *Container) error {
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return fmt.Errorf("invalid density value: %v", err)
	}

	switch operator {
	case "=", "==":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.GetDensity() == value
		})
	case "!=":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.GetDensity() != value
		})
	case ">":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.GetDensity() > value
		})
	case ">=":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.GetDensity() >= value
		})
	case "<":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.GetDensity() < value
		})
	case "<=":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.GetDensity() <= value
		})
	default:
		return fmt.Errorf("invalid operator for density field: %s", operator)
	}
	return nil
}

func processRadiusCondition(operator, valueStr string, container *Container) error {
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return fmt.Errorf("invalid radius value: %v", err)
	}

	switch operator {
	case "=", "==":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.Ball != nil && float64(f.Ball.Radius) == value
		})
	case "!=":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.Ball != nil && float64(f.Ball.Radius) != value
		})
	case ">":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.Ball != nil && float64(f.Ball.Radius) > value
		})
	case ">=":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.Ball != nil && float64(f.Ball.Radius) >= value
		})
	case "<":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.Ball != nil && float64(f.Ball.Radius) < value
		})
	case "<=":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.Ball != nil && float64(f.Ball.Radius) <= value
		})
	default:
		return fmt.Errorf("invalid operator for radius field: %s", operator)
	}
	return nil
}

func processEdgeCondition(operator, valueStr string, container *Container, edgeType string) error {
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return fmt.Errorf("invalid %s value: %v", edgeType, err)
	}

	getEdgeValue := func(f *threedfigure.ThreeDFigure) (float64, bool) {
		if f.Parallelepiped == nil {
			return 0, false
		}
		
		switch edgeType {
		case "edgeA":
			return float64(f.Parallelepiped.Rebro1), true
		case "edgeB":
			return float64(f.Parallelepiped.Rebro2), true
		case "edgeC":
			return float64(f.Parallelepiped.Rebro3), true
		default:
			return 0, false
		}
	}

	switch operator {
	case "=", "==":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			edgeValue, ok := getEdgeValue(f)
			return ok && edgeValue == value
		})
	case "!=":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			edgeValue, ok := getEdgeValue(f)
			return ok && edgeValue != value
		})
	case ">":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			edgeValue, ok := getEdgeValue(f)
			return ok && edgeValue > value
		})
	case ">=":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			edgeValue, ok := getEdgeValue(f)
			return ok && edgeValue >= value
		})
	case "<":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			edgeValue, ok := getEdgeValue(f)
			return ok && edgeValue < value
		})
	case "<=":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			edgeValue, ok := getEdgeValue(f)
			return ok && edgeValue <= value
		})
	default:
		return fmt.Errorf("invalid operator for %s field: %s", edgeType, operator)
	}
	return nil
}

func processHeightCondition(operator, valueStr string, container *Container) error {
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return fmt.Errorf("invalid height value: %v", err)
	}

	switch operator {
	case "=", "==":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.Cylinder != nil && f.Cylinder.Height == value
		})
	case "!=":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.Cylinder != nil && f.Cylinder.Height != value
		})
	case ">":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.Cylinder != nil && f.Cylinder.Height > value
		})
	case ">=":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.Cylinder != nil && f.Cylinder.Height >= value
		})
	case "<":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.Cylinder != nil && f.Cylinder.Height < value
		})
	case "<=":
		container.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.Cylinder != nil && f.Cylinder.Height <= value
		})
	default:
		return fmt.Errorf("invalid operator for height field: %s", operator)
	}
	return nil
}

func ReadAndProcessFile(filename string, container *Container) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer func() {
    if err := file.Close(); err != nil {
        log.Printf("Failed to close file: %v", err)
    }
	}()

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		command, args := ParseCommand(line)

		switch strings.ToLower(command) {
		case "add":
			if err := ProcessADD(args, container); err != nil {
				fmt.Printf("Error processing ADD at line %d: %v\n", lineNumber, err)
			}
		case "rem":
			if err := ProcessREM(args, container); err != nil {
				fmt.Printf("Error processing REM at line %d: %v\n", lineNumber, err)
			}
		case "print":
			container.Print()
		default:
			fmt.Printf("Unknown command at line %d: %s\n", lineNumber, command)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	return nil
}