package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MinTemp = 15
	MaxTemp = 30
)

var ErrUnknownOperator = errors.New("unknown operator")

func main() {
	if err := run(); err != nil {
		fmt.Printf("Execution error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	scanner := bufio.NewScanner(os.Stdin)

	departments, err := readInt(scanner)
	if err != nil {
		return fmt.Errorf("failed to read number of departments: %w", err)
	}

	employees, err := readInt(scanner)
	if err != nil {
		return fmt.Errorf("failed to read number of employees: %w", err)
	}

	for departmentIndex := 0; departmentIndex < departments; departmentIndex++ {
		if err := processDepartment(scanner, employees); err != nil {
			return fmt.Errorf("error in department %d: %w", departmentIndex+1, err)
		}
	}

	return nil
}

func readInt(scanner *bufio.Scanner) (int, error) {
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		val, err := strconv.Atoi(line)
		if err != nil {
			return 0, fmt.Errorf("invalid integer: %w", err)
		}

		return val, nil
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scanner error: %w", err)
	}

	return 0, errors.New("unexpected end of input")
}

func processDepartment(scanner *bufio.Scanner, employees int) error {
	low := MinTemp
	high := MaxTemp

	for employeeIndex := 0; employeeIndex < employees; employeeIndex++ {
		var operator string

		var value int

		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				continue
			}

			var err error

			operator, value, err = parseConstraint(line)
			if err != nil {
				return fmt.Errorf("failed to parse preference: %w", err)
			}

			break
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

		switch {
		case low > high:
			fmt.Println("-1")
		case operator == ">=":
			fmt.Println(low)
		default:
			fmt.Println(high)
		}
	}

	return nil
}

func parseConstraint(line string) (string, int, error) {
	if strings.HasPrefix(line, ">=") {
		val, err := strconv.Atoi(strings.TrimSpace(line[2:]))
		if err != nil {
			return "", 0, fmt.Errorf("invalid integer in '%s': %w", line, err)
		}

		return ">=", val, nil
	}

	if strings.HasPrefix(line, "<=") {
		val, err := strconv.Atoi(strings.TrimSpace(line[2:]))
		if err != nil {
			return "", 0, fmt.Errorf("invalid integer in '%s': %w", line, err)
		}

		return "<=", val, nil
	}

	return "", 0, fmt.Errorf("unknown operator: %s", line)
}
