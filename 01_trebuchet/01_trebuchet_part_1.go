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
		//depth, err := strconv.Atoi(scanner.Text())
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
