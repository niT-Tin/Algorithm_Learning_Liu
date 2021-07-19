package main

import (
	"fmt"
	"math"
	"math/rand"
)

var Sum = 0

// Version 1.0
func maxVal_v1(startx, starty int, mp *[][]int) {
	// get the row and col of the map
	r := len(*mp)
	c := len((*mp)[0])
	if starty == 0 && startx == 0 {
		Sum += (*mp)[startx][starty]
	}
	if startx != c-1 || starty != r-1 {
		maxV, nextI, nextJ := maxRound(startx, starty, r, c, mp)
		Sum += maxV
		maxVal_v1(nextI, nextJ, mp)
	}
}

// Version2.0
func maxVal_v2(m *[][]int) {
	var i, j = 0, 0
	Sum += (*m)[i][j]
	i1, j1 := len(*m), len((*m)[0])
	fmt.Printf("i1 = %d j1 = %d \n", i1, j1)
	for i != i1-1 || j != j1-1 {
		a, b, c, d := i+1, j, i, j+1
		fmt.Printf("i = %d j = %d \n", i, j)
		e, f := float64((*m)[a][b]), float64((*m)[c][d])
		max := math.Max(e, f)
		Sum += int(max)
		if e > f {
			i = i + 1
		} else if e < f {
			j = j + 1
		}
	}
}

// get the max value of two direction and each coordinate
// and return the coordinate and add the max value to the sum
// the coordinate will participate in the next recursive
func maxRound(x, y, row, col int, mp *[][]int) (int, int, int) {
	var v1, v2, max int
	var maxI, maxJ int
	if x >= 0 && x <= col-1 && y >= 0 && y <= row-1 {
		if (x + 1) <= col-1 {
			v1 = (*mp)[y][x+1]
		} else {
			v1 = 0
		}
		if (y + 1) <= row-1 {
			v2 = (*mp)[y+1][x]
		} else {
			v2 = -1
		}
	}
	if v1 > v2 {
		max = v1
		maxI, maxJ = x+1, y
	} else if v2 > v1 {
		max = v2
		maxI, maxJ = x, y+1
	} else {
		max = 0
		maxI, maxJ = x, y
	}
	return max, maxI, maxJ
}

func GenerateMap() [][]int {
	// get row and col of the map
	fmt.Println("Please enter r and c of the map")
	var row, col int
	_, err := fmt.Scanf("%d %d", &row, &col)
	if err != nil {
		panic(err)
	}
	//row, col = 3, 3

	// generate the specified size map
	ms := make([][]int, row)
	for i := range ms {
		ms[i] = make([]int, col)
		for j := range ms[i] {
			ms[i][j] = rand.Intn(21)
		}
	}
	return ms
}

func main() {

	ms := GenerateMap()

	fmt.Printf("The map is \n")
	showMap(ms)
	maxVal_v2(&ms)
	fmt.Printf("The max value is %d\n", Sum)

}

// As the name said
func showMap(ms [][]int) {
	for i := range ms {
		for j := range ms[i] {
			fmt.Printf("%3d", ms[i][j])
		}
		fmt.Println()
	}
}
