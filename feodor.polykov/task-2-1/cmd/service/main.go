package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	MinTemp = 15
	MaxTemp = 30
)

var (
	ErrUnknownOperator = errors.New("unknown operator")
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("Execution error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	var departments, employees int

	if _, err := fmt.Scan(&departments); err != nil {
		return fmt.Errorf("failed to read number of departments: %w", err)
	}

	if _, err := fmt.Scan(&employees); err != nil {
		return fmt.Errorf("failed to read number of employees: %w", err)
	}

	for i := 0; i < departments; i++ {
		if err := processDepartment(employees); err != nil {
			return fmt.Errorf("error in department %d: %w", i+1, err)
		}
	}

	return nil
}

func processDepartment(employees int) error {
	low := MinTemp
	high := MaxTemp

	for j := 0; j < employees; j++ {
		var operator string
		var val int

		if _, err := fmt.Scan(&operator, &val); err != nil {
			return fmt.Errorf("error reading employee %d preference: %w", j+1, err)
		}

		switch operator {
		case ">=":
			if val > low {
				low = val
			}
		case "<=":
			if val < high {
				high = val
			}
		default:
			return fmt.Errorf("%w: '%s', expected '>=' or '<='", ErrUnknownOperator, operator)
		}

		if low > high {
			fmt.Println("-1")
		} else {
			if operator == ">=" {
				fmt.Println(low)
			} else {
				fmt.Println(high)
			}
		}
	}

	return nil
}
