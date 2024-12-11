package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	//Task: Take in a set of rules and check them to the second set of input.
	//Syntax: A|B means A must come before B, the answer is all the middle numbers combined.
	//Approach: keep a '''''map''''' of the rules and check that the previous numbers don't violate it before inserting a number.
	

	reader := bufio.NewReader(os.Stdin)
	var rules []Rule
	input := ""

	for true {

		appended := false

		input, _ = reader.ReadString('\n')

		if input == "\n" {

			break
		}

		rule := ToRule(input[:len(input) - 1])

		for index, element := range(rules) {

			if rule.before == element.before {

				rules[index].afters = append(rules[index].afters, rule.afters[0])
				appended = true
				break
			}
		}

		if appended == false {

			rules = append(rules, rule)
		}
	}

	fmt.Println("-- RULES --")
	for _, rule := range(rules) {

		fmt.Println(rule)
	}

	page_num_sum := 0

	for true {

		input, _ = reader.ReadString('\n')
		fail := false
		num_buffer := ""
		var pages []int

		if input == "\n" {

			break
		}

		for _, char := range(input) {

			if char == ',' || char == '\n' {
				num := ParseNum(num_buffer)
				fail = FailRule(pages, rules, num)
				num_buffer = ""

				if fail == true {

					break
				} else {

					pages = append(pages, num)
				}
			} else {
				num_buffer = num_buffer + (string)(char)
			}
		}

		if fail == false {
			fmt.Println("PASS: ", pages)
			page_num_sum += pages[len(pages) / 2]
		}
	}

	fmt.Println("-- RESULT --")
	fmt.Println("PAGE NUM SUM: ", page_num_sum)
}

func FailRule(pages []int, rules []Rule, num int) (bool) {

	for _, rule := range(rules) {

		if rule.before == num {

			for _, num := range(pages) {

				for _, after := range(rule.afters) {

					if num == after {

						return true
					}
				}
			}

			break
		}
	}

	return false
}

func ToRule(word string) (Rule) {

	var rule Rule

	for index, character := range(word) {

		if character == '|' {

			rule.before = ParseNum(word[:index])
			rule.afters = append(rule.afters, ParseNum(word[index + 1:]))
			break
		}
	}

	return rule
}

func ParseNum(num string) int {

	num_bytes := ([]byte)(num)
	number := 0

	for _, char := range(num_bytes) {

		number *= 10
		number += int(char) - 48
	}

	return number
}

type Rule struct {
	before int
	afters []int
}
