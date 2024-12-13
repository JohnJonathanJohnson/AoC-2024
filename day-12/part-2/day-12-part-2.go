package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	var farm_grid [][]byte

	for true {

		input, _ := reader.ReadBytes('\n')
		input = input[:len(input)-1]

		if len(input) == 0 {
			break
		}

		farm_grid = append(farm_grid, input)
	}

	price := 0

	for y_loc := range farm_grid {
		for x_loc := range farm_grid[y_loc] {
			if farm_grid[y_loc][x_loc] != '-' {

				area, fences := calculate_fencing_price(farm_grid, y_loc, x_loc)
				fences += 4
				//Problem: +4 sides for each embedded field.
				//Make a function to scan for embedded fields?
				//Replace entire plot with 0 and +1 every time?
				fences -= 4 * check_embedded_plots(farm_grid)

				fmt.Println("AREA:", area, "FENCES", fences)
				price += area * (fences)

				replace_plot(farm_grid, y_loc, x_loc, '-')
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

	if y_loc > 0 && farm_grid[y_loc-1][x_loc] == plot {

		temp_area, temp_fences := calculate_fencing_price(farm_grid, y_loc-1, x_loc)
		area += temp_area
		fences += temp_fences
	}

	if x_loc > 0 && farm_grid[y_loc][x_loc-1] == plot {

		temp_area, temp_fences := calculate_fencing_price(farm_grid, y_loc, x_loc-1)
		area += temp_area
		fences += temp_fences
	}

	if y_loc < len(farm_grid)-1 && farm_grid[y_loc+1][x_loc] == plot {

		temp_area, temp_fences := calculate_fencing_price(farm_grid, y_loc+1, x_loc)
		area += temp_area
		fences += temp_fences
	}

	if x_loc < len(farm_grid[y_loc])-1 && farm_grid[y_loc][x_loc+1] == plot {

		temp_area, temp_fences := calculate_fencing_price(farm_grid, y_loc, x_loc+1)
		area += temp_area
		fences += temp_fences
	}

	//Based on "Folding"
	//Top-left corner _|
	if x_loc != 0 && y_loc != 0 && farm_grid[y_loc][x_loc-1] == '.' && farm_grid[y_loc-1][x_loc] == '.' && farm_grid[y_loc-1][x_loc-1] != '.' {

		fences += 2
	}

	//Bot-right corner |-
	if x_loc != len(farm_grid[y_loc])-1 && y_loc != len(farm_grid)-1 && farm_grid[y_loc][x_loc+1] == '.' && farm_grid[y_loc+1][x_loc] == '.' && farm_grid[y_loc+1][x_loc+1] != '.' {

		fences += 2
	}

	//Bot-left corner -|
	if x_loc != 0 && y_loc != len(farm_grid)-1 && farm_grid[y_loc][x_loc-1] == '.' && farm_grid[y_loc+1][x_loc] == '.' && farm_grid[y_loc+1][x_loc-1] != '.' {

		fences += 2
	}

	//Top-right corner |_
	if x_loc != len(farm_grid[y_loc])-1 && y_loc != 0 && farm_grid[y_loc][x_loc+1] == '.' && farm_grid[y_loc-1][x_loc] == '.' && farm_grid[y_loc-1][x_loc+1] != '.' {

		fences += 2
	}

	return area, fences
}

func replace_plot(farm_grid [][]byte, y_loc int, x_loc int, new_plot byte) {

	plot := farm_grid[y_loc][x_loc]
	farm_grid[y_loc][x_loc] = new_plot

	if y_loc > 0 && farm_grid[y_loc-1][x_loc] == plot {

		replace_plot(farm_grid, y_loc-1, x_loc, new_plot)
	}

	if x_loc > 0 && farm_grid[y_loc][x_loc-1] == plot {

		replace_plot(farm_grid, y_loc, x_loc-1, new_plot)
	}

	if y_loc < len(farm_grid)-1 && farm_grid[y_loc+1][x_loc] == plot {

		replace_plot(farm_grid, y_loc+1, x_loc, new_plot)
	}

	if x_loc < len(farm_grid[y_loc])-1 && farm_grid[y_loc][x_loc+1] == plot {

		replace_plot(farm_grid, y_loc, x_loc+1, new_plot)
	}
}

func check_embedded_plots(farm_grid [][]byte) int {

	plots := 0
	farm_grid_copy := [][]byte{}

	for _, row := range farm_grid {

		row_copy := []byte{}
		for _, plot := range row {

			row_copy = append(row_copy, plot)
		}
		farm_grid_copy = append(farm_grid_copy, row_copy)
	}

	//A solution most ingenious.
	//You isolate it on all borders so only the real holes and not the side pockets (dont count) would be left.
	for x := range(farm_grid_copy[0]) {

		if farm_grid_copy[0][x] != '.' {
			isolate(farm_grid_copy, 0, x)
		}
	}

	for y := range(farm_grid_copy[1:len(farm_grid_copy) - 1]) {

		if farm_grid_copy[y+1][0] != '.' && farm_grid_copy[y+1][len(farm_grid_copy[y])-1] != '.' {
			isolate(farm_grid_copy, y+1, 0)
			isolate(farm_grid_copy, y+1, len(farm_grid_copy[y])-1)
		}
	}

	for x := range(farm_grid_copy[len(farm_grid_copy) - 1]) {

		if farm_grid_copy[0][x] != '.' {
			isolate(farm_grid_copy, 0, x)
		}
	}

	for y_loc := range farm_grid_copy {
		for x_loc := range farm_grid_copy[y_loc] {
			if farm_grid_copy[y_loc][x_loc] != '.' {
				if check_embed(farm_grid_copy, y_loc, x_loc) == true {

					fmt.Println(farm_grid_copy)
					plots += 1
				}
			}
		}
	}

	return plots
}

func check_embed(farm_grid [][]byte, y_loc int, x_loc int) bool {

	//spread dots until a boundary or a dead-end is found. Use ANDs.
	farm_grid[y_loc][x_loc] = '.'
	return_bool := true

	if y_loc == 0 || x_loc == 0 || y_loc == len(farm_grid)-1 || x_loc == len(farm_grid[y_loc])-1 {

		return false
	}

	if farm_grid[y_loc-1][x_loc] == '.' && farm_grid[y_loc+1][x_loc] == '.' && farm_grid[y_loc][x_loc-1] == '.' && farm_grid[y_loc][x_loc+1] == '.' &&
		farm_grid[y_loc-1][x_loc-1] == '.' && farm_grid[y_loc-1][x_loc+1] == '.' && farm_grid[y_loc+1][x_loc-1] == '.' && farm_grid[y_loc+1][x_loc+1] == '.' {

		return true
	}

	if farm_grid[y_loc-1][x_loc] != '.' {

		return_bool = return_bool && check_embed(farm_grid, y_loc-1, x_loc)
	}

	if farm_grid[y_loc+1][x_loc] != '.' {

		return_bool = return_bool && check_embed(farm_grid, y_loc+1, x_loc)
	}

	if farm_grid[y_loc][x_loc-1] != '.' {

		return_bool = return_bool && check_embed(farm_grid, y_loc, x_loc-1)
	}

	if farm_grid[y_loc][x_loc+1] != '.' {

		return_bool = return_bool && check_embed(farm_grid, y_loc, x_loc+1)
	}

	//added diagonals
	if farm_grid[y_loc-1][x_loc-1] != '.' {

		return_bool = return_bool && check_embed(farm_grid, y_loc-1, x_loc-1)
	}
	if farm_grid[y_loc-1][x_loc+1] != '.' {

		return_bool = return_bool && check_embed(farm_grid, y_loc-1, x_loc+1)
	}
	if farm_grid[y_loc+1][x_loc-1] != '.' {

		return_bool = return_bool && check_embed(farm_grid, y_loc+1, x_loc-1)
	}
	if farm_grid[y_loc+1][x_loc+1] != '.' {

		return_bool = return_bool && check_embed(farm_grid, y_loc+1, x_loc+1)
	}

	return return_bool
}

func isolate(farm_grid [][]byte, y_loc int, x_loc int) {

	farm_grid[y_loc][x_loc] = '.'

	if y_loc != 0 {

		if farm_grid[y_loc-1][x_loc] != '.' {

			isolate(farm_grid, y_loc-1, x_loc)
		}
	}

	if y_loc != len(farm_grid)-1 {

		if farm_grid[y_loc+1][x_loc] != '.' {

			isolate(farm_grid, y_loc+1, x_loc)
		}
	}

	if x_loc != 0 {

		if farm_grid[y_loc][x_loc-1] != '.' {

			isolate(farm_grid, y_loc, x_loc-1)
		}
	}

	if x_loc != len(farm_grid[y_loc])-1 {

		if farm_grid[y_loc][x_loc+1] != '.' {

			isolate(farm_grid, y_loc, x_loc+1)
		}
	}

	//Diagonals
//	if y_loc != 0 && x_loc != 0 {
//
//		if farm_grid[y_loc-1][x_loc-1] != '.' {
//
//			isolate(farm_grid, y_loc-1, x_loc-1)
//		}
//	}
//
//	if y_loc != 0 && x_loc != len(farm_grid[0])-1 {
//
//		if farm_grid[y_loc-1][x_loc+1] != '.' {
//
//			isolate(farm_grid, y_loc-1, x_loc+1)
//		}
//	}
//
//	if y_loc != len(farm_grid)-1 && x_loc != 0 {
//
//		if farm_grid[y_loc+1][x_loc-1] != '.' {
//
//			isolate(farm_grid, y_loc+1, x_loc-1)
//		}
//	}
//
//	if y_loc != len(farm_grid)-1 && x_loc != len(farm_grid[0])-1 {
//
//		if farm_grid[y_loc+1][x_loc+1] != '.' {
//
//			isolate(farm_grid, y_loc+1, x_loc+1)
//		}
//	}
}
