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
	fmt.Println(len(list1), len(list2))

	result := compare(list1, list2)

	fmt.Println("RESULT: ", result)
}

func compare (list1 []int, list2 []int) int {

	//Going to do a weird kind of heap sort
	//TODO: Convert to a (value, repeats) dual value and when you calculare you just value * rep_left * rep_right
	//Changed goal: just search the heap until smaller (search in branches) and then cut off the front part that's bigger than the search point.
	var score int = 0
	count_left := 0
	count_right := 0
	num := 0

	for true {

		list1, _ = heapify(list1, 0)
		num = list1[0]

		if num == -1 {

			break
		}

		count_left, list1 = heap_search(list1, num, 0)

		list2, _ = heapify(list2, 0)
		count_right, list2 = heap_search(list2, num, 0)

		score += count_left * count_right * num

		if count_left > 0 && count_right > 0 {
			fmt.Println("NUM: ", num)
			fmt.Println("COUNT (LEFT): ", count_left)
			fmt.Println("COUNT (RIGHT): ", count_right)
			fmt.Println("SCORE: ", score)
			//fmt.Println("LIST 1: ", list1)
			//fmt.Println("LIST 2: ", list2)
		}
	}

	return score
}

func heap_search(list []int, num int, index int) (int, []int) {

	if index >= len(list) {

		return 0, list
	}

	if list[index] >= num {

		count_left := 0
		count_right := 0
		count_left, list = heap_search(list, num, index * 2 + 1) 
		count_right, list = heap_search(list, num, index * 2 + 2)
		if list[index] == num {

			list[index] = -1
			return 1 + count_left + count_right, list
		} else {

			return count_left + count_right, list
		}
	} else {

		return 0, list
	}
}

func heapify(list []int, index int) ([]int, bool) {

	next_left := 2 * index + 1
	is_heap := (list[index] >= next_left && list[index] >= next_left + 1)
	placeholder := 0

	if next_left >= len(list) {

		//fmt.Println("HEAP: ", list)
		return list, true
	} else if next_left + 1 == len(list) {

		if list[index] < list[next_left] {

			placeholder = list[next_left]
			list[next_left] = list[index]
			list[index] = placeholder
		}

		//fmt.Println("HEAP: ", list)
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

	//fmt.Println("HEAP: ",list)
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
