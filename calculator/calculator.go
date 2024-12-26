package calculator

import (
	"errors"
	"strconv"
)

func GetResult(op []string) (int, error) {
	opInt1, err := strconv.Atoi(op[0])
	if err != nil {
		return 0, err
	}
	opInt2, err := strconv.Atoi(op[2])
	if err != nil {
		return 0, err
	}
	if opInt2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	var result int
	switch op[1] {
	case "+":
		result = opInt1 + opInt2
	case "-":
		result = opInt1 - opInt2
	case "*":
		result = opInt1 * opInt2
	case "/":
		result = opInt1 / opInt2
	}
	return result, nil
}
