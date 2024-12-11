package main

import (
	"fmt"
	"bufio"
	"os"
	"slices"
)

func main() {

	//Task: Look for an X in a grid and expand it in 8 directions to find XMAS.

	reader := bufio.NewReader(os.Stdin)
	grid := [][]byte{}
	var input string

	for true {

		input, _ = reader.ReadString('\n')
		if input == "Q\n" {

			break
		}

		grid = append(grid, ([]byte)(input)[:len(input) - 1])
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
		if grid[y][x] != 'X' {
			x += 1
			continue
		}
		if y > 2 {
			//check the top
			if slices.Equal([]byte{grid[y - 1][x], grid[y - 2][x], grid[y - 3][x]}, []byte{'M', 'A', 'S'}) {
				amount += 1
				fmt.Println("FOUND")
			}
			if x > 2 {
				//check top left
				if slices.Equal([]byte{grid[y - 1][x - 1], grid[y - 2][x - 2], grid[y - 3][x - 3]}, []byte{'M', 'A', 'S'}) {
					amount += 1
					fmt.Println("FOUND")
				}
			}
			if x < len(grid[y]) - 3 {
				//check top right
				if slices.Equal([]byte{grid[y - 1][x + 1], grid[y - 2][x + 2], grid[y - 3][x + 3]}, []byte{'M', 'A', 'S'}) {
					amount += 1
					fmt.Println("FOUND")
				}
			}
		}
		if y < len(grid) - 3 {
			//check the bot
			if slices.Equal([]byte{grid[y + 1][x], grid[y + 2][x], grid[y + 3][x]}, []byte{'M', 'A', 'S'}) {
				amount += 1
				fmt.Println("FOUND")
			}
			if x > 2 {
				//check bot left
				if slices.Equal([]byte{grid[y + 1][x - 1], grid[y + 2][x - 2], grid[y + 3][x - 3]}, []byte{'M', 'A', 'S'}) {
					amount += 1
					fmt.Println("FOUND")
				}
			}
			if x < len(grid[y]) - 3 {
				//check bot right
				if slices.Equal([]byte{grid[y + 1][x + 1], grid[y + 2][x + 2], grid[y + 3][x + 3]}, []byte{'M', 'A', 'S'}) {
					amount += 1
					fmt.Println("FOUND")
				}
			}
		}
		if x > 2 {
			//check left
			if slices.Equal([]byte{grid[y][x - 1], grid[y][x - 2], grid[y][x - 3]}, []byte{'M', 'A', 'S'}) {
				amount += 1
				fmt.Println("FOUND")
			}
		}
		if x < len(grid[y]) - 3 {
			//check right
			if slices.Equal([]byte{grid[y][x + 1], grid[y][x + 2], grid[y][x + 3]}, []byte{'M', 'A', 'S'}) {
				amount += 1
				fmt.Println("FOUND")
			}
		}

		x += 1
	}

	fmt.Println("AMOUNT: ", amount)
}
