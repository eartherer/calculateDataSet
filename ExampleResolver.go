package main

import "strconv"

type ExampleReslover struct{}

func (r *ExampleReslover) isValid(input DataSetInput) bool {
	return true
}

func (r *ExampleReslover) resloveProblem(input DataSetInput) (ResloveResult, error) {
	result := ResloveResult{DataSetInput: input, Result: make(map[int]string)}
	for _, pos := range input.FindPosition {
		result.Result[pos] = strconv.Itoa(pos + 50)
	}
	return result, nil
}
