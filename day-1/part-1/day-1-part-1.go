package main

import (
	"fmt"
	"bufio"
	"os"
)

func main () {

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadBytes('Q')

	list1, list2 := parse_to_arrays(text)

	result := compare(list1, list2)

	fmt.Println(result)
}

func compare (list1 []int, list2 []int) int {

	//Going to do a weird kind of heap sort
	sumdiff := 0
	max_index := len(list1) - 1
	
	list1, _ = heapify(list1, 0)
	list2, _ = heapify(list2, 0)

	sumdiff += abs(list2[0] - list1[0])

	for i := 0; i < max_index; i += 1 {

		list1, _ = heapify(list1[1:], 0)
		list2, _ = heapify(list2[1:], 0)

		sumdiff += abs(list2[0] - list1[0])

		fmt.Println("SUMDIFF: ",sumdiff)
	}

	return sumdiff
}

func heapify(list []int, index int) ([]int, bool) {

	next_left := 2 * index + 1
	is_heap := (list[index] >= next_left && list[index] >= next_left + 1)
	placeholder := 0

	if next_left >= len(list) {

		fmt.Println("HEAP: ",list)
		return list, true
	} else if next_left + 1 == len(list) {

		if list[index] < list[next_left] {

			placeholder = list[index]
			list[index] = list[next_left]
			list[next_left] = placeholder
		}

		fmt.Println("HEAP: ",list)
		return list, true
	}

	for true {

	
		list, is_heap = heapify(list, next_left)	
		list, is_heap = heapify(list, next_left + 1)
	
		if list[next_left] > list[index] {
	
			placeholder = list[next_left]
			list[next_left] = list[index]
			list[index] = placeholder

			list, is_heap = heapify(list, next_left)
		}
	
		if list[next_left + 1 ] > list[index] {
	
			placeholder = list[next_left + 1]
			list[next_left + 1] = list[index]
			list[index] = placeholder

			list, is_heap = heapify(list, next_left + 1)
		}
	
		if is_heap == true && list[index] >= list[next_left] && list[index] >= list[next_left + 1] {

			break
		}
	}

	fmt.Println("HEAP: ",list)
	return list, is_heap
}

func parse_to_arrays(line []byte) ([]int, []int) {

	arr_1_turn := true
	array1 := []int{}
	array2 := []int{}
	buffer := []byte{}

	for i := 0; i < len(line); i += 1 {

		if (line[i] == ' ' || line[i] == '\n') && len(buffer) > 0 {

			if arr_1_turn == true {

				array1 = append(array1, parsenum(buffer))
				arr_1_turn = false
			} else {

				array2 = append(array2, parsenum(buffer))
				arr_1_turn = true
			}

			buffer = []byte{}
		} else if line[i] != ' ' && line[i] != '\n' {

			buffer = append(buffer, line[i])
		}
	}

	return array1, array2
}

func parsenum(word []byte) int {

	num := 0

	for i := 0; i < len(word); i += 1 {

		num *= 10
		num += (int)(word[i]) - 48
	}

	return num
}

func abs(num int) int {

	if num < 0 {

		return num * -1
	}

	return num
}
