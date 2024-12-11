package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	//Task: To find how many ways there are to place a single obstruction and create a loop.
	//Idea: To got hrough the normal path and while tracing it find a spot where if you turn right you go back to somewhere in the path.

	fmt.Println("-- Map Input --")
	reader := bufio.NewReader(os.Stdin)
	lab_map := [][]byte{}
	var guard [4]int

	for true {

		input, err := reader.ReadBytes('\n')

		if err != nil {

			panic("Map doko?")
		}

		if input[0] == (byte)('\n') {

			break
		}

		lab_map = append(lab_map, input[:len(input) - 1])
	}

	for row_num, row := range(lab_map) {

		for column_num, pixel := range(row) {

			switch {

			case pixel == '^':
				guard = [4]int{row_num, column_num, -1, 0}
				break
			case pixel == '<':
				guard = [4]int{row_num, column_num, 0, -1}
				break
			case pixel == '>':
				guard = [4]int{row_num, column_num, 0, 1}
				break
			case pixel == 'v':
				guard = [4]int{row_num, column_num, 1, 0}
				break
			}
		}
	}

	orig_map := CopyGrid(lab_map)
	orig_map[guard[0]][guard[1]] = '.'
	lab_map[guard[0]][guard[1]] = 'O'
	fmt.Println("-- INITIAL MAP --")
	PrintMap(lab_map)
	fmt.Println("Guard Location: ", guard)
	orig_guard := guard
	loop_opportunity := 0

	for true {

		if guard[0] >= len(lab_map) || guard[1] >= len(lab_map[0]) || guard[0] < 0 || guard[1] < 0 {

			break
		}
		map_copy := CopyGrid(orig_map)
		if guard[0] + guard[2] < len(lab_map) &&
		guard[1] + guard[3] < len(lab_map[guard[0]]) &&
		guard[0] + guard[2] >= 0 && guard[1] + guard[3] >= 0 &&
		lab_map[guard[0] + guard[2]][guard[1] + guard[3]] == '.' &&
		CheckLoop(guard, map_copy, orig_guard) == true {

				loop_opportunity += 1
				lab_map[guard[0] + guard[2]][guard[1] + guard[3]] = 'O'
			}

		switch {
		case lab_map[guard[0]][guard[1]] == '#':

			guard[0] -= guard[2]
			guard[1] -= guard[3]
			placeholder := guard[2]
			guard[2] = guard[3]
			guard[3] = -1 * placeholder

		default:
	
			guard[0] += guard[2]
			guard[1] += guard[3]
		}

	}
	
	fmt.Println("Opportunities: ", loop_opportunity)
	PrintMap(lab_map)
}

func CopyGrid(map_grid [][]byte) [][]byte {

	returned_map := [][]byte{}
	for _, row := range(map_grid) {

		row_copy := []byte{}

		for _, pixel := range(row) {

			row_copy = append(row_copy, pixel)
		}
		returned_map = append(returned_map, row_copy)
	}

	return returned_map
}
func CheckLoop(guard_curr [4]int, map_grid [][]byte, guard [4]int) bool {

//	if guard[0] + guard[2] >= len(map_grid) || guard[1] + guard[3] >= len(map_grid[guard[0]]) || guard[0] + guard[2] < 0 || guard[1] + guard[3] < 0 {
//
//		return false
//	}
	map_grid[guard_curr[0] + guard_curr[2]][guard_curr[1] + guard_curr[3]] = '#'
	same_x_count := -1
	same_y_count := -1
	last_y := 0
	last_x := 0

	for true {

		if guard[0] + guard[2] == len(map_grid) || guard[1] + guard[3] == len(map_grid[0]) || guard[0] + guard[2] == -1 || guard[1] + guard[3] == -1 {

			return false
		}

		switch {
		case map_grid[guard[0] + guard[2]][guard[1] + guard[3]] == '#':

			placeholder := guard[2]
			guard[2] = guard[3]
			guard[3] = -1 * placeholder
			//Checking for thin loops
			if same_x_count > 4 {

				return true
			}
			if guard[1] == last_x {

				same_x_count += 1
			} else {

				same_x_count = 0
				last_x = guard[1]
			}

			if same_y_count > 4 {

				return true
			}
			if guard[0] == last_y {

				same_y_count += 1
			} else {

				same_y_count = 0
				last_y = guard[0]
			}

		case map_grid[guard[0]][guard[1]] == '^' && guard[2] == -1 && guard[3] == 0:
			return true
		case map_grid[guard[0]][guard[1]] == '<' && guard[2] == 0 && guard[3] == -1:
			return true
		case map_grid[guard[0]][guard[1]] == '>' && guard[2] == 0 && guard[3] == 1:
			return true
		case map_grid[guard[0]][guard[1]] == 'v' && guard[2] == 1 && guard[3] == 0:
			return true

		default:

			switch {
			case guard[2] == -1 && guard[3] == 0:
				map_grid[guard[0]][guard[1]] = '^'
			case guard[2] == 0 && guard[3] == -1:
				map_grid[guard[0]][guard[1]] = '<'
			case guard[2] == 0 && guard[3] == 1:
				map_grid[guard[0]][guard[1]] = '>'
			case guard[2] == 1 && guard[3] == 0:
				map_grid[guard[0]][guard[1]] = 'v'
			}
			guard[0] += guard[2]
			guard[1] += guard[3]
		}

	
	}

	return false
}

func PrintMap(map_grid [][]byte) {

	fmt.Println("-- MAP --")
	for _, row := range(map_grid) {

		for _, pixel := range(row) {

			fmt.Print((string)(pixel))
		}

		fmt.Println()
	}
}
