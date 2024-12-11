package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	var grid [][]byte
	total_score := 0

	for true {

		input, _ := reader.ReadBytes('\n')
		input = input[:len(input) - 1]

		if len(input) == 0 {

			break
		}

		grid = append(grid, input)
	}
	
	for y_cord, row := range(grid) {
		for x_cord, pixel := range(row) {

			if pixel == '0' {

				total_score += get_score(x_cord, y_cord, grid)
			}
		}
	}

	fmt.Println(total_score)
}

func get_score(x_cord int, y_cord int, grid [][]byte) (int) {

	//recursive.
	if grid[y_cord][x_cord] == '9' {

		return 1
	}

	score_sum := 0

	if x_cord > 0  && int(grid[y_cord][x_cord]) + 1 == int(grid[y_cord][x_cord - 1]) {

		score_sum += get_score(x_cord - 1, y_cord, grid)
	}

	if x_cord < len(grid[y_cord]) - 1  && int(grid[y_cord][x_cord]) + 1 == int(grid[y_cord][x_cord + 1]) {

		score_sum += get_score(x_cord + 1, y_cord, grid)
	}

	if y_cord > 0  && int(grid[y_cord][x_cord]) + 1 == int(grid[y_cord - 1][x_cord]) {

		score_sum += get_score(x_cord, y_cord - 1, grid)
	}

	if y_cord < len(grid) - 1  && int(grid[y_cord][x_cord]) + 1 == int(grid[y_cord + 1][x_cord]) {

		score_sum += get_score(x_cord, y_cord + 1, grid)
	}

	return score_sum
}
