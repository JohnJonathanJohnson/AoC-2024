//Task: to check antennas and stuff idk how to explain...
//Approach: let's keep the same type of antennas in a slice of co-ordinations (x, y),
//then for each co-ordinate,
//you mirror it with another antenna coordinate (of the same type) and mark that spot as an antinode.
//Duplicates are treated as one antinode.
//Potential Problems: out of bounds, inaccurate antinode, comparing a coordiante to itself, etc. Remeber to use a different map from the coordinates for the antinodes.

package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	var grid [][]byte

	for true {

		input, _ := reader.ReadBytes('\n')
		input = input[:len(input) - 1]

		if len(input) == 0 {

			break
		} else {

			grid = append(grid, input)
		}
	}

	var sat_array [255][][2]int

	for y_cord, row := range(grid) {

		for x_cord, pixel := range(row) {

			if pixel != '.' {

				sat_array[int(pixel)] = append(sat_array[int(pixel)], [2]int{x_cord, y_cord})
			}
		}
	}

	var node_grid [][]bool

	for index, row := range(grid) {

		node_grid = append(node_grid, []bool{})
		for _, _ = range(row) {

			node_grid[index] = append(node_grid[index], false)
		}
	}

	node_count := 0

	for _, satellite := range(sat_array) {

		for index, location := range(satellite) {

			for other_index, other_location := range(satellite) {

				if other_index == index {

					continue
				}

				dist_x := other_location[0] - location[0]
				dist_y := other_location[1] - location[1]

				if location[0] - dist_x < len(grid) && location[1] - dist_y < len(grid[0]) &&
				location[0] - dist_x > -1 && location[1] - dist_y > - 1 &&
				node_grid[location[0] - dist_x][location[1] - dist_y] == false {
					
					node_grid[location[0] - dist_x][location[1] - dist_y] = true
					node_count += 1
				}
			}
		}
	}

	fmt.Println("NODE COUNT: ", node_count)
}
