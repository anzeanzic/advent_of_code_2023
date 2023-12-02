package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func main() {
	ReadFile()
}

func ReadFile() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	sum := 0;

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Fields(line)
		game_id, err := strconv.Atoi(strings.Replace(line_arr[1], ":", "", -1))

		if err != nil {
			panic(err)
		}

		counters := [3]int{0, 0, 0} // r, g, b
		number := 0
		color := ""
		sum_per_game := 0

		for i := 2; i < len(line_arr); i++ {
			if num, err := strconv.Atoi(line_arr[i]); err != nil {
				color = strings.Replace(strings.Replace(line_arr[i], ",", "", -1), ";", "", -1)
				
				switch {
				case color == "red":
					if number > counters[0] {
						counters[0] = number;
					}
				case color == "green":
					if number > counters[1] {
						counters[1] = number;
					}
				case color == "blue":
					if number > counters[2] {
						counters[2] = number;
					}
				}
			} else {
				number = num
			}
		}

		sum_per_game = counters[0] * counters[1] * counters[2]
		fmt.Println(game_id, sum_per_game)
		sum += sum_per_game
	}

	file.Close()
	fmt.Println("---------------------------------------")
	fmt.Println("Increased counter: " + strconv.Itoa(sum))
}
