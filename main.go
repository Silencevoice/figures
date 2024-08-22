package main

import (
	"cifras/operations"
	"fmt"

	"math/rand"
)

func resolveTargetNumber(numbers []int, target int) (bool, []operations.Operation, int, []operations.Operation) {
	if len(numbers) == 1 {
		if numbers[0] == target {
			return true, []operations.Operation{}, numbers[0], []operations.Operation{}
		}
		return false, nil, numbers[0], []operations.Operation{}
	}

	bestDifference := 9999999
	bestResult := 0
	bestHistory := []operations.Operation{}

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			a, b := numbers[i], numbers[j]

			possibleResults := []operations.Operation{
				operations.Sum(a, b),
			}

			if op, valid := operations.Subtract(a, b); valid {
				possibleResults = append(possibleResults, op)
			}

			if op, valid := operations.Subtract(b, a); valid {
				possibleResults = append(possibleResults, op)
			}

			possibleResults = append(possibleResults, operations.Multiply(a, b))

			if op, valid := operations.Divide(a, b); valid {
				possibleResults = append(possibleResults, op)
			}

			if op, valid := operations.Divide(b, a); valid {
				possibleResults = append(possibleResults, op)
			}

			newList := make([]int, 0, len(numbers)-1)
			newList = append(newList, numbers[:i]...)
			newList = append(newList, numbers[i+1:j]...)
			newList = append(newList, numbers[j+1:]...)

			for _, operation := range possibleResults {
				if ok, history, _, _ := resolveTargetNumber(append(newList, operation.Result), target); ok {
					return true, append([]operations.Operation{operation}, history...), operation.Result, append([]operations.Operation{operation}, history...)
				} else {
					difference := abs(operation.Result - target)
					if difference < bestDifference {
						bestDifference = difference
						bestResult = operation.Result
						bestHistory = append([]operations.Operation{operation}, history...)
					}
				}
			}
		}
	}

	return false, nil, bestResult, bestHistory
}

// abs returns the absolute value for integers.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Select playing numbers and target number.
func selectNumbers() ([]int, int) {

	// Possible number list
	possibleNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 25, 50, 75, 100}

	// Select six numbers without repeating.
	selectedNumbers := make([]int, 0, 6)
	for i := 0; i < 6; i++ {
		idx := rand.Intn(len(possibleNumbers))
		selectedNumbers = append(selectedNumbers, possibleNumbers[idx])
		// Remove selected number to avoid repetitions
		possibleNumbers = append(possibleNumbers[:idx], possibleNumbers[idx+1:]...)
	}

	// Select a target number between 101 and 999
	target := 101 + rand.Intn(899) // 999-101+1

	return selectedNumbers, target
}

func main() {
	numbers, target := selectNumbers()

	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Target: %d\n", target)

	if ok, history, bestResult, bestHistory := resolveTargetNumber(numbers, target); ok {
		fmt.Println("Exact target number is possible")
		for _, operation := range history {
			fmt.Println(operation)
		}
	} else {
		fmt.Printf("Exact target number unreacheable. Best approx: %d\n", bestResult)
		for _, operation := range bestHistory {
			fmt.Println(operation)
		}
	}
}
