package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {

	//Purpose: parse random text to detect "mul([0-9]+,[0-9]+)"
	reader := bufio.NewReader(os.Stdin)
	in := ""
	input := ""

	for true {

		in, _ = reader.ReadString('\n')

		if in == "Q\n" {

			break
		}

		input += in
	}

	value := 0

	split_input := strings.Split(input, "mul(")

	for _, word := range(split_input) {


		value += checkmul(word)
	}

	fmt.Println("VALUE: ", value)
}
func checkmul(word string) int {

	num1, num2 := 0, 0
	mode_num1 := true

	for _, character := range(([]byte)(word)) {

		if mode_num1 == true {

			if character >= '0' && character <= '9' {

				num1 *= 10
				num1 += (int)(character) - 48
			} else if character == ',' {

				mode_num1 = false
			} else {

				return 0
			}
		} else {

			if character >= '0' && character <= '9' {

				num2 *= 10
				num2 += (int)(character) - 48
			} else if character == ')' {

				break
			} else {

				return 0
			}
		}
	}

	return num1 * num2
}
