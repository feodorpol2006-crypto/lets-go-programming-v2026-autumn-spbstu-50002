package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MinTemp = 15
	MaxTemp = 30
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("Execution error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	scanner := bufio.NewScanner(os.Stdin)

	departments, err := readNumber(scanner)
	if err != nil {
		return fmt.Errorf("failed to read number of departments: %w", err)
	}

	employees, err := readNumber(scanner)
	if err != nil {
		return fmt.Errorf("failed to read number of employees: %w", err)
	}

	for i := 0; i < departments; i++ {
		if err := processDepartment(scanner, employees); err != nil {
			return fmt.Errorf("error in department %d: %w", i+1, err)
		}
	}

	return nil
}

func readNumber(scanner *bufio.Scanner) (int, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return 0, err
		}
		return 0, fmt.Errorf("unexpected end of input")
	}

	line := strings.TrimSpace(scanner.Text())
	if line == "" {
		return 0, fmt.Errorf("empty string instead of number")
	}

	num, err := strconv.Atoi(line)
	if err != nil {
		return 0, fmt.Errorf("invalid number: '%s'", line)
	}

	return num, nil
}

func readBound(scanner *bufio.Scanner) (string, int, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", 0, err
		}
		return "", 0, fmt.Errorf("unexpected end of input")
	}

	line := strings.TrimSpace(scanner.Text())
	parts := strings.Fields(line)

	if len(parts) != 2 {
		return "", 0, fmt.Errorf("invalid string format: '%s', expected 'sign number'", line)
	}

	op := parts[0]
	val, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, fmt.Errorf("invalid temperature value: '%s'", parts[1])
	}

	return op, val, nil
}

func processDepartment(scanner *bufio.Scanner, employees int) error {
	low := MinTemp
	high := MaxTemp

	for j := 0; j < employees; j++ {
		op, val, err := readBound(scanner)
		if err != nil {
			return fmt.Errorf("error reading employee %d preference: %w", j+1, err)
		}

		switch op {
		case ">=":
			if val > low {
				low = val
			}
		case "<=":
			if val < high {
				high = val
			}
		default:
			return fmt.Errorf("unknown operator '%s', expected '>=' or '<='", op)
		}

		if low > high {
			fmt.Println("-1")
		} else {
			if op == ">=" {
				fmt.Println(low)
			} else {
				fmt.Println(high)
			}
		}
	}

	return nil
}
