//Task: See if it's possible for the number on the left to be formed from the numbers from the right.
//Idea: take the number from the left and compare it right to left.
//If it's divisible, divide. If it's not, subtract.
//Then, if at the final step, it ends at zero, it must be right.
//Potential Problems: Numbers that are divisible that are actually added.

package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	type Equation struct {result int; components []int}

	reader := bufio.NewReader(os.Stdin)
	valid_equations_sum := 0

	for true {

		input, _ := reader.ReadBytes('\n')
		
		if input[0] == '\n' {

			break
		}

		var equation Equation
		equation.result = 0
		equation.components = []int{}

		for index, digit := range(input) {

			if digit == ':' {

				input = input[index + 2:]
				break
			}

			equation.result *= 10
			equation.result += int(digit) - 48
		}

		number := 0
		for _, digit := range(input) {

			if digit == ' ' || digit == '\n' {

				equation.components = append(equation.components, number)
				number = 0
			} else {

				number *= 10
				number += int(digit) - 48
			}
		}

		result_copy := equation.result
		components_copy := equation.components
		for true {

			last_ele := len(equation.components) - 1

			if equation.components[last_ele] != 0 && equation.result % equation.components[last_ele] == 0 {

				equation.result /= equation.components[last_ele]
			} else {

				equation.result -= equation.components[last_ele]
			}

			if len(equation.components) == 2 {

				break
			} else {

				equation.components = equation.components[:last_ele]
			}
		}

		if equation.result == equation.components[0] || exhaustive_check(result_copy, components_copy[1:], components_copy[0]) {

			valid_equations_sum += result_copy
		} else {

			fmt.Println("FALSE: ", equation)
		}
	}

	fmt.Println("VALID EQUATIONS SUM: ", valid_equations_sum)
}

func exhaustive_check(expected_result int, components []int, result int) bool {

	if result == expected_result && len(components) == 0 {

		fmt.Println(expected_result, components, result)
		return true
	} else if len(components) == 0 {

		fmt.Println(expected_result, components, result)
		return false
	}

	return exhaustive_check(expected_result, components[1:], result + components[0]) || exhaustive_check(expected_result, components[1:], result * components[0])
}
