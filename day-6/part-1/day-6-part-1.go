package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

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

	fmt.Println("-- INITIAL MAP --")
	PrintMap(lab_map)
	fmt.Println("Guard Location: ", guard)

	path_len := 0

	for true {

		if guard[0] >= len(lab_map) || guard[1] >= len(lab_map[0]) {

			break
		}

		switch {
		case lab_map[guard[0]][guard[1]] == '#':

			guard[0] -= guard[2]
			guard[1] -= guard[3]
			placeholder := guard[2]
			guard[2] = guard[3]
			guard[3] = -1 * placeholder

			// -1,0 0,1 1,0 0,-1 when left goes to right, it flips. When right goes to left, it doesn't.
		case lab_map[guard[0]][guard[1]] == 'X':

			guard[0] += guard[2]
			guard[1] += guard[3]
		default:

			lab_map[guard[0]][guard[1]] = 'X'
			guard[0] += guard[2]
			guard[1] += guard[3]
			path_len += 1
		}

	}
	
	PrintMap(lab_map)
	fmt.Println("LEN: ", path_len)

}

func PrintMap(map_grid [][]byte) {

	for _, row := range(map_grid) {

		for _, pixel := range(row) {

			fmt.Print((string)(pixel))
		}

		fmt.Println()
	}
}
