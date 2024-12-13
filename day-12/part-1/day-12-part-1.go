package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	var farm_grid [][]byte

	for true {

		input, _ := reader.ReadBytes('\n')
		input = input[:len(input) - 1]

		if len(input) == 0 {
			break
		}

		farm_grid = append(farm_grid, input)
	}

	price := 0

	for y_loc := range(farm_grid) {
		for x_loc := range(farm_grid[y_loc]) {
			if farm_grid[y_loc][x_loc] != '-' {

				area, fences := calculate_fencing_price(farm_grid, y_loc, x_loc)
				fmt.Println("AREA:", area, "FENCES", fences)
				price += area * fences
				
				replace_plot(farm_grid, y_loc, x_loc)
			}
		}
	}

	fmt.Println("TOTAL PRICE:", price)
}

func calculate_fencing_price(farm_grid [][]byte, y_loc int, x_loc int) (int, int) {

	area := 1
	fences := 0
	plot := farm_grid[y_loc][x_loc]
	farm_grid[y_loc][x_loc] = '.'

	if y_loc > 0 && farm_grid[y_loc - 1][x_loc] == plot {

		temp_area, temp_fences := calculate_fencing_price(farm_grid, y_loc - 1, x_loc)
		area += temp_area
		fences += temp_fences
	} else if y_loc <= 0 || (farm_grid[y_loc - 1][x_loc] != plot && farm_grid[y_loc - 1][x_loc] != '.') {

		fences += 1
	}

	if x_loc > 0 && farm_grid[y_loc][x_loc - 1] == plot {

		temp_area, temp_fences := calculate_fencing_price(farm_grid, y_loc, x_loc - 1)
		area += temp_area
		fences += temp_fences
	} else if x_loc <= 0 || (farm_grid[y_loc][x_loc - 1] != plot && farm_grid[y_loc][x_loc - 1] != '.') {

		fences += 1
	}

	if y_loc < len(farm_grid) - 1 && farm_grid[y_loc + 1][x_loc] == plot {

		temp_area, temp_fences := calculate_fencing_price(farm_grid, y_loc + 1, x_loc)
		area += temp_area
		fences += temp_fences
	} else if y_loc >= len(farm_grid) - 1 || (farm_grid[y_loc + 1][x_loc] != plot && farm_grid[y_loc + 1][x_loc] != '.') {

		fences += 1
	}

	if x_loc < len(farm_grid[y_loc]) - 1 && farm_grid[y_loc][x_loc + 1] == plot {

		temp_area, temp_fences := calculate_fencing_price(farm_grid, y_loc, x_loc + 1)
		area += temp_area
		fences += temp_fences
	} else if x_loc >= len(farm_grid[y_loc]) - 1 || (farm_grid[y_loc][x_loc + 1] != plot && farm_grid[y_loc][x_loc + 1] != '.') {

		fences += 1
	}

	return area, fences
}

func replace_plot(farm_grid [][]byte, y_loc int, x_loc int) {

	plot := farm_grid[y_loc][x_loc]
	farm_grid[y_loc][x_loc] = '-'

	if y_loc > 0 && farm_grid[y_loc - 1][x_loc] == plot {

		replace_plot(farm_grid, y_loc - 1, x_loc)
	}

	if x_loc > 0 && farm_grid[y_loc][x_loc - 1] == plot {

		replace_plot(farm_grid, y_loc, x_loc - 1)
	}

	if y_loc < len(farm_grid) - 1 && farm_grid[y_loc + 1][x_loc] == plot {

		replace_plot(farm_grid, y_loc + 1, x_loc)
	}

	if x_loc < len(farm_grid[y_loc]) - 1 && farm_grid[y_loc][x_loc + 1] == plot {

		replace_plot(farm_grid, y_loc, x_loc + 1)
	}
}
