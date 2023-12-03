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
		for x, char := range line {
			// check if *
			if string(char) == "*" {
				num := CheckIfNumberNearby(arr, point{ x: x, y: y})
				
				if num >= 0 {
					sum += num
				}
			}
		}
	}

	file.Close()
	fmt.Println("---------------------------------------")
	fmt.Println("Increased counter: " + strconv.Itoa(sum))
}

func CheckIfNumberNearby(arr []string, pos point) (int) {
	fmt.Println(pos)
	var numbers []int
	var points []point

	// check up
	if pos.y - 1 >= 0 {
		x_start := 0
		x_end := 0

		if pos.x - 1 >= 0 {
			x_start = pos.x - 1
		} else {
			x_start = pos.x
		}

		if pos.x + 1 < len(arr[pos.y - 1]) {
			x_end = pos.x + 1
		} else {
			x_end = pos.x
		}

		for i := x_start; i <= x_end; i++ {
			num, err := strconv.Atoi(string(arr[pos.y - 1][i]))

			if err == nil {
				isChecked := CheckIfPointAlreadyChecked(points, point{ x: i, y: pos.y - 1 })

				if !isChecked {
					fmt.Println(num, point{ x: i, y: pos.y - 1 })
					full_num, temp_points := GetNumber(arr, point{ x: i, y: pos.y - 1 })
					points = temp_points
					numbers = append(numbers, full_num)
				}
			}
		}
	}

	// check down
	if pos.y + 1 < len(arr) {
		x_start := 0
		x_end := 0

		if pos.x - 1 >= 0 {
			x_start = pos.x - 1
		} else {
			x_start = pos.x
		}

		if pos.x + 1 < len(arr[pos.y + 1]) {
			x_end = pos.x + 1
		} else {
			x_end = pos.x
		}

		for i := x_start; i <= x_end; i++ {
			num, err := strconv.Atoi(string(arr[pos.y + 1][i]))

			if err == nil {
				isChecked := CheckIfPointAlreadyChecked(points, point{ x: i, y: pos.y + 1 })

				if !isChecked {
					fmt.Println("*" , num, point{ x: i, y: pos.y + 1 })
					full_num, temp_points := GetNumber(arr, point{ x: i, y: pos.y + 1 })
					points = temp_points
					numbers = append(numbers, full_num)
				}
			}
		}
	}

	// check middle left
	if pos.x - 1 >= 0 {
		num, err := strconv.Atoi(string(arr[pos.y][pos.x - 1]))

		if err == nil {
			isChecked := CheckIfPointAlreadyChecked(points, point{ x: pos.x - 1, y: pos.y })

			if !isChecked {
				fmt.Println("*" , num, point{ x: pos.x - 1, y: pos.y })
				full_num, temp_points := GetNumber(arr, point{ x: pos.x - 1, y: pos.y })
				points = temp_points
				numbers = append(numbers, full_num)
			}
		}
	}
	
	// check middle right
	if pos.x + 1 < len(arr[pos.y]) {
		num, err := strconv.Atoi(string(arr[pos.y][pos.x + 1]))

		if err == nil {
			isChecked := CheckIfPointAlreadyChecked(points, point{ x: pos.x + 1, y: pos.y })

			if !isChecked {
				fmt.Println("*" , num, point{ x: pos.x + 1, y: pos.y })
				full_num, temp_points := GetNumber(arr, point{ x: pos.x + 1, y: pos.y })
				points = temp_points
				numbers = append(numbers, full_num)
			}
		}
	}

	// check if two numbers
	fmt.Println(numbers)
	fmt.Println("============================")

	if len(numbers) == 2 {
		return numbers[0] * numbers[1];
	}

	return -1;
}

func GetNumber(arr []string, pos point) (int, []point) {
	num_str := ""
	prevIsNum := true;
	var points []point

	// before the number
	for i := pos.x - 1; i >= 0; i-- {
		_, err := strconv.Atoi(string(arr[pos.y][i]))
	
		if err == nil {
			num_str = string(arr[pos.y][i]) + num_str;
			points = append(points, point{ x: i, y: pos.y })
			prevIsNum = true
		} else {
			if (prevIsNum) {
				break
			}

			prevIsNum = false
		}
	}

	num_str += string(arr[pos.y][pos.x]);
	points = append(points, point{ x: pos.x, y: pos.y })
	prevIsNum = true

	// after the number
	for i := pos.x + 1; i < len(arr[pos.y]); i++ {
		_, err := strconv.Atoi(string(arr[pos.y][i]))

		if err == nil {
			num_str += string(arr[pos.y][i]);
			points = append(points, point{ x: i, y: pos.y })
			prevIsNum = true
		} else {
			if (prevIsNum) {
				break
			}

			prevIsNum = false
		}
	}

	full_num, err := strconv.Atoi(num_str)

	if err != nil {
		panic(err)
	}

	return full_num, points
}

func CheckIfPointAlreadyChecked(point_arr []point, point point) (bool) {
	for i := 0; i < len(point_arr); i++ {
		if point_arr[i].x == point.x && point_arr[i].y == point.y {
			return true
		}
	}

	return false
}