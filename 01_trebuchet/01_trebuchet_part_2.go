package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	ReadFile()
}

func ReadFile() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	number := 0

	for scanner.Scan() {
		str := scanner.Text()
		fmt.Println(scanner.Text())

		firstNumber := ""
		lastNumber := ""

		for i := 0; i < len(str); i++ {
			if _, err := strconv.Atoi(string(str[i])); err == nil {
				fmt.Printf("%q looks like a number.\n", string(str[i]))

				if firstNumber == "" {
					firstNumber = string(str[i])
				}

				lastNumber = string(str[i])
			} else {
				word := string(str[i])

				for j := i + 1; j < len(str); j++ {
					if _, err := strconv.Atoi(string(str[i])); err == nil {
						word += string(str[j])
						numStr := checkIfNumber(word)

						if numStr != "" {
							if firstNumber == "" {
								firstNumber = numStr
							}

							lastNumber = numStr
						}
						break
					} else {
						word += string(str[j])
						numStr := checkIfNumber(word)

						if numStr != "" {
							if firstNumber == "" {
								firstNumber = numStr
							}

							lastNumber = numStr
							break
						}
					}
				}
			}
		}

		number, err = strconv.Atoi(firstNumber + lastNumber)

		if err != nil {
			panic(err)
		}

		sum += number

		fmt.Println("Num: " + strconv.Itoa(number))
	}

	file.Close()
	fmt.Println("---------------------------------------")
	fmt.Println("Increased counter: " + strconv.Itoa(sum))
}

func checkIfNumber(word string) string {
	number := ""

	switch {
	case word == "one":
		number = "1"
		break
	case word == "two":
		number = "2"
		break
	case word == "three":
		number = "3"
		break
	case word == "four":
		number = "4"
		break
	case word == "five":
		number = "5"
		break
	case word == "six":
		number = "6"
		break
	case word == "seven":
		number = "7"
		break
	case word == "eight":
		number = "8"
		break
	case word == "nine":
		number = "9"
		break
	}

	return number
}
