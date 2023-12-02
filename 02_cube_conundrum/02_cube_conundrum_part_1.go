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
		all_games_ok := true

		for i := 2; i < len(line_arr); i++ {
			if num, err := strconv.Atoi(line_arr[i]); err != nil {
				color = strings.Replace(strings.Replace(line_arr[i], ",", "", -1), ";", "", -1)
				
				switch {
				case color == "red":
					counters[0] += number
				case color == "green":
					counters[1] += number
				case color == "blue":
					counters[2] += number
				}
			} else {
				number = num
			}

			// check if game is possible
			if strings.Contains(line_arr[i], ";") { 
				if (!(counters[0] <= MAX_RED && counters[1] <= MAX_GREEN && counters[2] <= MAX_BLUE)) {
					all_games_ok = false
					break
				}

				counters = [3]int{0, 0, 0}
			}
		}

		// check if game is possible - last one
		if (!(counters[0] <= MAX_RED && counters[1] <= MAX_GREEN && counters[2] <= MAX_BLUE)) {
			all_games_ok = false
		}

		if (all_games_ok) {
			fmt.Println(game_id)
			sum += game_id
		}
	}

	file.Close()
	fmt.Println("---------------------------------------")
	fmt.Println("Increased counter: " + strconv.Itoa(sum))
}
