package main

import (
	"mystuff"
	"fmt"
	"os"
	"bufio"
	"strings"
)

//Task: Find how many "reports" contain only ups or downs and not less than one or more than three.

func main() {

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('Q')
	strings := strings.Split(input, "\n")
	valid_report_count := 0
	fmt.Println(strings[len(strings) - 1])
	
	for i, str := range(strings[:len(strings) - 1]) {

		valid_report_count += check_report(parse.ParseToArray(str))
		fmt.Println("INDEX: ", i, "SCORE: ", valid_report_count)
	}

	fmt.Println("COUNT: ", valid_report_count)
}

func check_report(array []int) int {

	if len(array) < 3 {

		return 1
	}

	diff := array[len(array) - 1] - array[0]

	if diff > 0 {

		return check_report_increasing(array)
	} else {

		return check_report_decreasing(array)
	}
}

func check_report_increasing(array []int) (int) {

	last_num := array[0]

	for _, current_num := range(array[1:]) {

		if last_num >= current_num || current_num > last_num + 3 {

			return 0
		}

		last_num = current_num
	}

	fmt.Println(array)
	return 1
}

func check_report_decreasing(array []int) (int) {

	last_num := array[0]

	for _, current_num := range(array[1:]) {

		if last_num <= current_num || current_num < last_num - 3 {

			return 0
		}

		last_num = current_num
	}

	fmt.Println(array)
	return 1
}
