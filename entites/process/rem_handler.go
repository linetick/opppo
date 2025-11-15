// Package process handles command parsing and execution.
package process

import (
	"fmt"
	"pr1/entites/container"
	"pr1/entites/threedfigure"
	"strconv"
	"strings"
)

func handleREM(args []string, c *container.Container) error {
	if len(args) < 3 {
		return fmt.Errorf("invalid REM command format. Usage: REM field operator value")
	}

	field := args[0]
	operator := args[1]
	valueStr := strings.Join(args[2:], " ")

	switch strings.ToLower(field) {
	case "owner":
		return processOwnerCondition(operator, valueStr, c)
	case "density":
		return processDensityCondition(operator, valueStr, c)
	case "radius":
		return processRadiusCondition(operator, valueStr, c)
	case "rebro1", "edgea":
		return processEdgeCondition(operator, valueStr, c, "edgeA")
	case "rebro2", "edgeb":
		return processEdgeCondition(operator, valueStr, c, "edgeB")
	case "rebro3", "edgec":
		return processEdgeCondition(operator, valueStr, c, "edgeC")
	case "height":
		return processHeightCondition(operator, valueStr, c)
	default:
		return fmt.Errorf("unknown field: %s", field)
	}
}

func processOwnerCondition(operator, valueStr string, c *container.Container) error {
	switch operator {
	case "=", "==":
		c.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.GetOwnerName() == valueStr
		})
	case "!=":
		c.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return f.GetOwnerName() != valueStr
		})
	case "contains", "CONTAINS":
		c.Remove(func(f *threedfigure.ThreeDFigure) bool {
			return strings.Contains(strings.ToLower(f.GetOwnerName()), strings.ToLower(valueStr))
		})
	default:
		return fmt.Errorf("invalid operator for owner field: %s", operator)
	}
	return nil
}

func processDensityCondition(operator, valueStr string, c *container.Container) error {
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return fmt.Errorf("invalid density value: %v", err)
	}

	cond := getNumericCondition(operator, value)
	if cond == nil {
		return fmt.Errorf("invalid operator for density field: %s", operator)
	}

	c.Remove(func(f *threedfigure.ThreeDFigure) bool {
		return cond(f.GetDensity())
	})
	return nil
}

func processRadiusCondition(operator, valueStr string, c *container.Container) error {
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return fmt.Errorf("invalid radius value: %v", err)
	}

	cond := getNumericCondition(operator, value)
	if cond == nil {
		return fmt.Errorf("invalid operator for radius field: %s", operator)
	}

	c.Remove(func(f *threedfigure.ThreeDFigure) bool {
		return f.Ball != nil && cond(float64(f.Ball.Radius))
	})
	return nil
}

func processEdgeCondition(operator, valueStr string, c *container.Container, edgeType string) error {
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return fmt.Errorf("invalid %s value: %v", edgeType, err)
	}

	cond := getNumericCondition(operator, value)
	if cond == nil {
		return fmt.Errorf("invalid operator for %s field: %s", edgeType, operator)
	}

	getEdge := func(f *threedfigure.ThreeDFigure) (float64, bool) {
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

	c.Remove(func(f *threedfigure.ThreeDFigure) bool {
		if val, ok := getEdge(f); ok {
			return cond(val)
		}
		return false
	})
	return nil
}

func processHeightCondition(operator, valueStr string, c *container.Container) error {
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return fmt.Errorf("invalid height value: %v", err)
	}

	cond := getNumericCondition(operator, value)
	if cond == nil {
		return fmt.Errorf("invalid operator for height field: %s", operator)
	}

	c.Remove(func(f *threedfigure.ThreeDFigure) bool {
		return f.Cylinder != nil && cond(f.Cylinder.Height)
	})
	return nil
}

// getNumericCondition возвращает функцию-условие для числовых операторов.
func getNumericCondition(op string, value float64) func(float64) bool {
	switch op {
	case "=", "==":
		return func(x float64) bool { return x == value }
	case "!=":
		return func(x float64) bool { return x != value }
	case ">":
		return func(x float64) bool { return x > value }
	case ">=":
		return func(x float64) bool { return x >= value }
	case "<":
		return func(x float64) bool { return x < value }
	case "<=":
		return func(x float64) bool { return x <= value }
	default:
		return nil
	}
}