// // Package process handles command parsing and file processing.
package process

import (
	"fmt"
	"pr1/entites/container"
	"pr1/entites/figurefactory"
	"pr1/entites/file"
	"strings"
)

func ParseCommand(line string) (string, []string) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return "", nil
	}
	return strings.ToUpper(parts[0]), parts[1:]
}

func ExecuteFile(filename string, c *container.Container) error {
	lines, err := file.ReadCommands(filename)
	if err != nil {
		return err
	}

	for i, line := range lines {
		cmd, args := ParseCommand(line)
		if err := ExecuteCommand(cmd, args, c); err != nil {
			return fmt.Errorf("line %d: %w", i+1, err)
		}
	}
	return nil
}

func ExecuteCommand(command string, args []string, c *container.Container) error {
	switch command {
	case "ADD":
		return handleADD(args, c)
	case "REM":
		return handleREM(args, c)
	case "PRINT":
		c.Print()
		return nil
	default:
		return fmt.Errorf("unknown command: %s", command)
	}
}

func handleADD(args []string, c *container.Container) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: ADD FIGURE_TYPE params...")
	}

	figType := figurefactory.FigureType(args[0])
	figure, err := figurefactory.Create(figType, args[1:])
	if err != nil {
		return fmt.Errorf("ADD failed: %w", err)
	}

	c.Add(figure)
	return nil
}