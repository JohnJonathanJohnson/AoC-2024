package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadBytes('\n')
	stones := parse_int_array(input)

	count := 0
	curr_count := 0
	cases := [][3]int{}

	for _, stone := range(stones) {

		curr_count, cases = split_stone(stone, 75, [][3]int{})
		count += curr_count
		fmt.Println(stone, ":", count)
	}

	fmt.Println("LOGGED CASES:", cases)
	fmt.Println("STONE COUNT:", count)
}
//Thank you, Nitro-Carwash, for the "memo" idea.
func split_stone(stone int, steps_left int, same_cases [][3]int) (int, [][3]int) {

	for _, Case := range(same_cases) {

		if stone == Case[0] && steps_left == Case[1] {

			return Case[2], same_cases
		}
	}
	if steps_left == 0 {
	
		return 1, same_cases
	}
	
	if stone == 0 {

		return split_stone(1, steps_left - 1, same_cases)
	} else if digits(stone) % 2 == 0 {

		splitting_point := digits(stone)/2
		exponent := 1

		for _ = range(splitting_point) {

			exponent *= 10
		}

		front := stone / exponent
		back := stone - front * exponent
		
		front, same_cases = split_stone(front, steps_left - 1, same_cases)
		back, same_cases = split_stone(back, steps_left - 1, same_cases)
		count := front + back
		same_cases = append(same_cases, [3]int{stone, steps_left, count})
		
		return count, same_cases
	}

	return split_stone(stone * 2024, steps_left - 1, same_cases)
}

func digits(num int) int {

	digits := 0

	for true {

		if num <= 0 {
			break
		}
		digits += 1
		num /= 10
	}

	return digits
}

func parse_int_array(input []byte) []int {

	int_array := []int{}
	number := 0
	for _, character := range(input) {

		if character == ' ' || character == '\n' {

			int_array = append(int_array, number)
			number = 0
		} else {

			number *= 10
			number += int(character) - 48
		}
	}

	return int_array
}


