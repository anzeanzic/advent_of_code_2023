package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type point struct {
	x int
	y int
}

func main() {
	ReadFile()
}

func ReadFile() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	var arr []string
	scanner := bufio.NewScanner(file)
	

	for scanner.Scan() {
		line := scanner.Text()
		arr = append(arr, line)
	}

	// go through lines
	sum := 0;

	for y, line := range arr {
		prevIsNum := false
		num := ""
		numStartPos := point{ x: 0, y: 0 }
		numEndPos := point{x: 0, y: 0 }

		for x, char := range line {
			// check if number
			if _, err := strconv.Atoi(string(char)); err == nil {
				//fmt.Println(string(char))

				if !prevIsNum {
					numStartPos.x = x
					numStartPos.y = y
				}

				num += string(char);
				prevIsNum = true
			} else {
				if (prevIsNum) {
					numEndPos.x = x - 1
					numEndPos.y = y
					signNearby := CheckIfSignNearby(arr, numStartPos, numEndPos)
					fmt.Println(num);

					if signNearby {
						fmt.Println("sign nearby")
						fullNum, err := strconv.Atoi(num)
						sum += fullNum

						if err != nil {
							panic(err)
						}
					} else {
						//fmt.Println(num);
						//fmt.Println("no sign")
					}
				}

				prevIsNum = false
				numStartPos = point{ x: 0, y: 0 }
				numEndPos = point{x: 0, y: 0 }
				num = ""
			}
		}

		if prevIsNum {
			numEndPos.x = len(line) - 1
			numEndPos.y = y
			signNearby := CheckIfSignNearby(arr, numStartPos, numEndPos)
			fmt.Println(num);

			if signNearby {
				//fmt.Println("sign nearby")
				fullNum, err := strconv.Atoi(num)
				sum += fullNum

				if err != nil {
					panic(err)
				}
			}
		}
	}

	file.Close()
	fmt.Println("---------------------------------------")
	fmt.Println("Increased counter: " + strconv.Itoa(sum))
}

func CheckIfSignNearby(arr []string, startPos point, endPos point) (bool) {
	// check up
	if startPos.y - 1 >= 0 {
		x_start := 0
		x_end := 0

		if startPos.x - 1 >= 0 {
			x_start = startPos.x - 1
		} else {
			x_start = startPos.x
		}

		if endPos.x + 1 < len(arr[startPos.y - 1]) {
			x_end = endPos.x + 1
		} else {
			x_end = endPos.x
		}

		for i := x_start; i <= x_end; i++ {
			//fmt.Println(i, startPos.y - 1, string(arr[startPos.y - 1][i]))
			isNearby := CheckIfPosIsSign(string(arr[startPos.y - 1][i]))

			if isNearby {
				return true
			}
		}
	}

	// check down
	if startPos.y + 1 < len(arr) {
		x_start := 0
		x_end := 0

		if startPos.x - 1 >= 0 {
			x_start = startPos.x - 1
		} else {
			x_start = startPos.x
		}

		if endPos.x + 1 < len(arr[startPos.y + 1]) {
			x_end = endPos.x + 1
		} else {
			x_end = endPos.x
		}

		for i := x_start; i <= x_end; i++ {
			//fmt.Println(i, startPos.y + 1, string(arr[startPos.y + 1][i]))
			isNearby := CheckIfPosIsSign(string(arr[startPos.y + 1][i]))

			if isNearby {
				return true
			}
		}
	}

	// check middle left
	if startPos.x - 1 >= 0 {
		isNearby := CheckIfPosIsSign(string(arr[startPos.y][startPos.x - 1]))

		if isNearby {
			return true
		}
	}
	
	// check middle right
	if endPos.x + 1 < len(arr[startPos.y]) {
		//fmt.Println(endPos.x + 1, startPos.y, string(arr[startPos.y][endPos.x + 1]))
		isNearby := CheckIfPosIsSign(string(arr[startPos.y][endPos.x + 1]))

		if isNearby {
			return true
		}
	}

	return false
}

func CheckIfPosIsSign(pos_str string) (bool) {
	_, err := strconv.Atoi(pos_str)

	return pos_str != "." && err != nil;
	//return pos_str == "*" || pos_str == "#" || pos_str == "+" || pos_str == "$" 
}