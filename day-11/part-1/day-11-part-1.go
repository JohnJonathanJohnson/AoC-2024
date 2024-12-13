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

	for _ = range(25) {

		stones = split_stones(stones)
		fmt.Println(len(stones))
	}
	fmt.Println("STONE COUNT:", len(stones))

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

func split_stones(stones []int) []int {

	var new_stones []int

	for _, stone := range(stones) {

		new_stones = append(new_stones, split_stone(stone)...)
	}

	return new_stones
}

func split_stone(stone int) []int {

	if stone == 0 {

		return []int{1}
	} else if digits(stone) % 2 == 0 {

		splitting_point := digits(stone)/2
		exponent := 1

		for _ = range(splitting_point) {

			exponent *= 10
		}

		front := stone / exponent
		back := stone - front * exponent
		
		return []int{front, back}
	}

	return []int{stone * 2024}
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
