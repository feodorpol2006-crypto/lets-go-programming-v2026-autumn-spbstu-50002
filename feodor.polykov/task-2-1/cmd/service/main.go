package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	MinTemp = 15
	MaxTemp = 30
)

var ErrUnknownOperator = errors.New("unknown operator")

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Execution error:", err)
		os.Exit(1)
	}
}

func run() error {
	reader := bufio.NewReader(os.Stdin)

	var departments int

	_, err := fmt.Fscan(reader, &departments)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}

		return fmt.Errorf("failed to read number of departments: %w", err)
	}

	var employees int

	_, err = fmt.Fscan(reader, &employees)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}

		return fmt.Errorf("failed to read number of employees: %w", err)
	}

	for i := 0; i < departments; i++ {
		if err := processDepartment(reader, employees); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}

			return fmt.Errorf("error in department %d: %w", i+1, err)
		}
	}

	return nil
}

func processDepartment(reader *bufio.Reader, employees int) error {
	low := MinTemp
	high := MaxTemp

	for i := 0; i < employees; i++ {
		var op string

		var val int

		_, err := fmt.Fscan(reader, &op, &val)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}

			return fmt.Errorf("failed to read preference: %w", err)
		}

		if err := updateBounds(op, val, &low, &high); err != nil {
			return err
		}

		printResult(op, low, high)
	}

	return nil
}

func updateBounds(op string, val int, low *int, high *int) error {
	switch op {
	case ">=":
		if val > *low {
			*low = val
		}
	case "<=":
		if val < *high {
			*high = val
		}
	default:
		return fmt.Errorf("%w: '%s'", ErrUnknownOperator, op)
	}

	return nil
}

func printResult(op string, low int, high int) {
	switch {
	case low > high:
		fmt.Println("-1")
	case op == ">=":
		fmt.Println(low)
	default:
		fmt.Println(high)
	}
}
