/*
Max Increase to Keep City Skyline
Example:
Input: grid = [[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]]
Output: 35
Explanation:
The grid is:
[ [3, 0, 8, 4],
  [2, 4, 5, 7],
  [9, 2, 6, 3],
  [0, 3, 1, 0] ]

The skyline viewed from top or bottom is: [9, 4, 8, 7]
The skyline viewed from left or right is: [8, 7, 9, 3]

The grid after increasing the height of buildings without affecting skylines is:

gridNew = [ [8, 4, 8, 7],
            [7, 4, 7, 7],
            [9, 4, 8, 7],
            [3, 3, 3, 3] ]
*/

package main

import (
	"fmt"
        "math"
)

func maxIncreaseKeepingSkyline(grid [][]int) int {
	var sum int = 0
	row := make([]int, len(grid))    //from left or right
	col := make([]int, len(grid[0])) // from top or bottom
	for i, v := range grid {
		for j, k := range v {
			row[i] = int(math.Max(float64(row[i]), float64(k)))
			col[j] = int(math.Max(float64(col[j]), float64(k)))
		}
	}
	fmt.Println("row is", row, "col is ", col)
	for i, v := range grid {
		for j, k := range v {
			sum += (int(math.Min(float64(row[i]), float64(col[j]))) - k)
		}
	}
	return sum
}

func main() {
	grid := [][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}
	fmt.Println(maxIncreaseKeepingSkyline(grid))
}
