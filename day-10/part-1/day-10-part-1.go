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

				var grid_copy [][]byte
				for index, row := range(grid) {
					grid_copy = append(grid_copy, []byte{})
					for _, char := range(row) {
						grid_copy[index] = append(grid_copy[index], char)
					}
				}

				score, _ := get_score(x_cord, y_cord, grid_copy)
				total_score += score
			}
		}
	}

	fmt.Println(total_score)
}

func get_score(x_cord int, y_cord int, grid [][]byte) (int, [][]byte) {

	//recursive.
	if grid[y_cord][x_cord] == '9' {

		grid[y_cord][x_cord] = 'X'
		return 1, grid
	}

	score_sum := 0
	score := 0

	if x_cord > 0  && int(grid[y_cord][x_cord]) + 1 == int(grid[y_cord][x_cord - 1]) {

		score, grid = get_score(x_cord - 1, y_cord, grid)
		score_sum += score
		score = 0
	}

	if x_cord < len(grid[y_cord]) - 1  && int(grid[y_cord][x_cord]) + 1 == int(grid[y_cord][x_cord + 1]) {

		score, grid = get_score(x_cord + 1, y_cord, grid)
		score_sum += score
		score = 0
	}

	if y_cord > 0  && int(grid[y_cord][x_cord]) + 1 == int(grid[y_cord - 1][x_cord]) {

		score, grid = get_score(x_cord, y_cord - 1, grid)
		score_sum += score
		score = 0
	}

	if y_cord < len(grid) - 1  && int(grid[y_cord][x_cord]) + 1 == int(grid[y_cord + 1][x_cord]) {

		score, grid = get_score(x_cord, y_cord + 1, grid)
		score_sum += score
		score = 0
	}

	return score_sum, grid
}
