package main

import (
	"fmt"
	"strconv"

	"math/rand"
)

func sum(a, b int) int {
	return a + b
}

func subtract(a, b int) (int, bool) {
	if a >= b {
		return a - b, true
	}
	return 0, false
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) (int, bool) {
	if b != 0 && a%b == 0 {
		return a / b, true
	}
	return 0, false
}

func resolveTargetNumber(numbers []int, target int) (bool, []string, int, []string) {
	// Base case: if only one number remains, check if we have reached the target
	if len(numbers) == 1 {
		if numbers[0] == target {
			return true, []string{}, numbers[0], []string{}
		}
		return false, nil, numbers[0], []string{}
	}

	bestDifference := 9999999
	bestResult := 0
	bestHistory := []string{}

	// Run through all pairs of numbers
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			a, b := numbers[i], numbers[j]

			// Possible returns of the operations
			posiblesResultados := map[int]string{
				sum(a, b): strconv.Itoa(a) + " + " + strconv.Itoa(b),
			}

			if resultado, valido := subtract(a, b); valido {
				posiblesResultados[resultado] = strconv.Itoa(a) + " - " + strconv.Itoa(b)
			}

			if resultado, valido := subtract(b, a); valido {
				posiblesResultados[resultado] = strconv.Itoa(b) + " - " + strconv.Itoa(a)
			}

			posiblesResultados[multiply(a, b)] = strconv.Itoa(a) + " * " + strconv.Itoa(b)

			if resultado, valido := divide(a, b); valido {
				posiblesResultados[resultado] = strconv.Itoa(a) + " / " + strconv.Itoa(b)
			}

			if resultado, valido := divide(b, a); valido {
				posiblesResultados[resultado] = strconv.Itoa(b) + " / " + strconv.Itoa(a)
			}

			// Generate a new list without a and b
			newList := make([]int, 0, len(numbers)-1)
			newList = append(newList, numbers[:i]...)
			newList = append(newList, numbers[i+1:j]...)
			newList = append(newList, numbers[j+1:]...)

			// Check every possible result recursively
			for result, operation := range posiblesResultados {
				if ok, history, _, _ := resolveTargetNumber(append(newList, result), target); ok {
					return true, append([]string{operation + " = " + strconv.Itoa(result)}, history...), result, append([]string{operation + " = " + strconv.Itoa(result)}, history...)
				} else {
					difference := abs(result - target)
					if difference < bestDifference {
						bestDifference = difference
						bestResult = result
						bestHistory = append([]string{operation + " = " + strconv.Itoa(result)}, history...)
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
