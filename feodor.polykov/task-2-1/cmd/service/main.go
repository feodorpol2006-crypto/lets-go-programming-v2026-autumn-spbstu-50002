package main

import (
	"errors"
	"fmt"
	"io"
)

const (
	MinTemp = 15
	MaxTemp = 30
)

var ErrUnknownOperator = errors.New("unknown operator")

func main() {
	if err := run(); err != nil && !errors.Is(err, io.EOF) {
		fmt.Printf("Execution error: %v\n", err)
	}
}

func run() error {
	var departments int
	var employees int

	if _, err := fmt.Scan(&departments); err != nil {
		return fmt.Errorf("failed to read number of departments: %w", err)
	}

	if _, err := fmt.Scan(&employees); err != nil {
		return fmt.Errorf("failed to read number of employees: %w", err)
	}

	for departmentIndex := 0; departmentIndex < departments; departmentIndex++ {
		if err := processDepartment(employees); err != nil {
			return fmt.Errorf("error in department %d: %w", departmentIndex+1, err)
		}
	}

	return nil
}

func processDepartment(employees int) error {
	low := MinTemp
	high := MaxTemp

	for employeeIndex := 0; employeeIndex < employees; employeeIndex++ {
		var operator string

		var value int

		if _, err := fmt.Scan(&operator, &value); err != nil {
			return fmt.Errorf("failed to read employee preference: %w", err)
		}

		switch operator {
		case ">=":
			if value > low {
				low = value
			}
		case "<=":
			if value < high {
				high = value
			}
		default:
			return fmt.Errorf("%w: '%s'", ErrUnknownOperator, operator)
		}

		if low > high {
			fmt.Println("-1")
		} else if operator == ">=" {
			fmt.Println(low)
		} else {
			fmt.Println(high)
		}
	}

	return nil
}
