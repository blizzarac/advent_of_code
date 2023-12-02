package main


import (
	"slices"
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)


func importGrid() [][]int{
	file,err := os.Open("test_input.txt")
	if err != nil {
		fmt.Println("Error", err)
	}

	scanner := bufio.NewScanner(file)

	result := [][]int{}
	row := 0
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), "")
		result = append(result, []int{})
		for i := range splitLine {
			treeHeight,_ := strconv.Atoi(splitLine[i])
			result[row] = append(result[row], treeHeight)
		}

		row++
	}
	return result
}

func findUncoveredInRow(row []int) int {
	result := 0
	//result = append(result, i)
	heighestLeft := 0
	heightestRight := 0
	for i:=1; i<len(row)-1; i++ {
//		fmt.Println("Row", row, "Index", i, "Value", row[i])
//		fmt.Println("Highest left slice", row[:i])
//		fmt.Println("Highest right slice", row[i+1:])
		heighestLeft = slices.Max(row[:i])
		heightestRight = slices.Max(row[i+1:])
//		fmt.Println("Heighest left", heighestLeft)
//		fmt.Println("Heighest right", heightestRight)

		if row[i] <= heighestLeft && row[i] <= heightestRight {
//			fmt.Println("Found", i)
			result++
		}
	}
	return result
}

func printGrid[T int|string](grid [][]T) {
	for i := range grid {
		for j := range grid[i] {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}

func findMaxScenicScore(grid [][]int) (int, [][]string) {
	result := 0
	debugGrid := make([][]string, len(grid))
	for i := 0; i < len(grid); i++ {
		debugGrid[i] = make([]string, len(grid[0]))
	}

	for i:=0; i<len(grid)-1; i++ {
		for j:=0; j<len(grid[i])-1; j++ {
			score := calculateScenicScore(grid, i, j)
			debugGrid[i][j] = "[" + strconv.Itoa(score) + "]"
			if score > result {
				result =  score
			}
		}
	}

	return result, debugGrid
}

func calculateScenicScore(grid [][]int, i int, j int) int {
	result := 0

	// search left
	for k:=j-1; k>=0; k-- {
		result++
		if grid[i][k] >= grid[i][j] {
			break
		}
	}

	// search right
	for k:=j+1; k<len(grid[i]); k++ {
		result++
		if grid[i][k] >= grid[i][j] {
			break
		}

	}

	// search up
	for k:=i-1; k>=0; k-- {
		result++
		if grid[k][j] >= grid[i][j] {
			break
		}

	}

	// search down
	for k:=i+1; k<len(grid); k++ {
		result++
		if grid[k][j] >= grid[i][j] {
			break
		}

	}

	return result
}


func findUncoveredTrees(grid [][]int) (int, [][]int) {
	result := 0
	reverseGrid := transpose(grid)

	debugGrid := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
    	debugGrid[i] = make([]int, len(grid[0]))
	}

	highestLeft := 0
	highestRight := 0
	highestUp := 0
	highestDown := 0

	for i:=1; i<len(grid)-1; i++ {
		for j:=1; j<len(grid[i])-1; j++ {
//			printGrid(grid)
//			printGrid(reverseGrid)
//			fmt.Println("Index", i, j, "Value", grid[i][j])
//			fmt.Println("Highest left slice", grid[i][:j])
//			fmt.Println("Highest right slice", grid[i][j+1:])
//			fmt.Println("Highest up slice", reverseGrid[j][:i])
//			fmt.Println("Highest down slice", reverseGrid[j][i+1:])
			highestLeft = slices.Max(grid[i][:j])
			highestRight = slices.Max(grid[i][j+1:])
			highestUp = slices.Max(reverseGrid[j][:i])
			highestDown = slices.Max(reverseGrid[j][i+1:])
//			fmt.Println("Heighest left", highestLeft)
//			fmt.Println("Heighest right", highestRight)
//			fmt.Println("Heighest up", highestUp)
//			fmt.Println("Heighest down", highestDown)

			if grid[i][j] <= highestLeft && grid[i][j] <= highestRight && grid[i][j] <= highestUp && grid[i][j] <= highestDown {
				result++
				debugGrid[i][j] = 1
			} else {
				debugGrid[i][j] = 0
			}
		}
	}

	return result, debugGrid
}

func transpose(a [][]int) [][]int {
    newArr := make([][]int, len(a))
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(a[0]); j++ {
            newArr[j] = append(newArr[j], a[i][j])
        }
    }
    return newArr
}

func main() {
	grid := importGrid()

	result, resultGrid :=findUncoveredTrees(grid)

	printGrid(resultGrid)
	fmt.Println((99*99)-result)


	result2, resultGrid2 := findMaxScenicScore(grid)
	printGrid(resultGrid2)
	fmt.Println(result2)
}
