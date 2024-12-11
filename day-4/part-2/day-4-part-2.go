package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	grid := [][]byte{}
	var input string

	for true {

		input, _ = reader.ReadString('\n')
		if input == "Q\n" {

			break
		}

		grid = append(grid, ([]byte)(input[:len(input) - 1]))
	}

	fmt.Println(grid)
	amount := 0
	x, y := 0, 0
	for true  {

		if x >= len(grid[y]) {
			y += 1
			x = 0
			fmt.Println("Y++")
		}
		if y >= len(grid) {
			break
		}
		if grid[y][x] != 77 && grid[y][x] != 83 {
			x += 1
			continue
		}

		if y < len(grid) - 2 && x < len(grid[y]) - 2 {

			if (int)(grid[y + 2][x] + grid[y][x + 2]) + (int)(grid[y + 2][x + 2] + grid[y][x]) == (77 * 2 + 83 * 2) {

				if grid[y + 1][x + 1] == 'A' && grid[y + 2][x + 2] + grid[y][x] == ('M' + 'S') {

					amount += 1
				}
			}
		}

		x += 1
	}

	fmt.Println("AMOUNT: ", amount)
}
