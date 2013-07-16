// By starting at the top of the triangle below and moving to adjacent numbers
// on the row below, the maximum total from top to bottom is 23.
//
//
// That is, 3 + 7 + 4 + 9 = 23.
// Find the maximum total from top to bottom of the triangle below:


package main

import "fmt"
import "strconv"
import "strings"

const triangle_str = `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`


func max_path(x, y int, grid [][]int) int {
	if x == len(grid) ||  y == len(grid[x]) {
		return 0
	}

	fmt.Println(x, y)
	n1 := max_path(x + 1, y, grid)
	n2 := max_path(x + 1, y + 1, grid)

	if n1 > n2 {
		return grid[x][y] + n1
	} else {
		return grid[x][y] + n2
	}
}

func max_total_path(grid [][]int) int {
	return max_path(0, 0, grid)
}

func main() {
	triangle_lines := strings.Split(triangle_str,"\n")
	triangle := make([][]int, len(triangle_lines))


	for i, line := range triangle_lines {
		for _, num_str := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(num_str)
			triangle[i] = append(triangle[i], num)
		}
	}

	fmt.Println(max_total_path(triangle))
}
