package main


import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func getFirstLast(line string) string {
		numOnly := ""
		for i := 0; i < len(line); i++ {
			_, err := strconv.Atoi(string(line[i]))
			if err == nil {
				numOnly += string(line[i])
			}
		}
		firstLast :=  string(numOnly[0])+string(numOnly[len(numOnly)-1])
		return firstLast
}

func replaceStringNumber(line string) string {
	line = strings.Replace(line, "one", "1", -1)
	line = strings.Replace(line, "two", "2", -1)
	line = strings.Replace(line, "three", "3", -1)
	line = strings.Replace(line, "four", "4", -1)
	line = strings.Replace(line, "five", "5", -1)
	line = strings.Replace(line, "six", "6", -1)
	line = strings.Replace(line, "seven", "7", -1)
	line = strings.Replace(line, "eight", "8", -1)
	line = strings.Replace(line, "nine", "9", -1)
	line = strings.Replace(line, "zero", "0", -1)

	return line
}


func ReplaceStringNumbers(line string) string {
	for i := 1; i < len(line) + 1 ; i++ {
		line = replaceStringNumber(line[0:i]) + line[i:]
	}

	return line
}

func CalculateLineCount(line string) int {
	line = ReplaceStringNumbers(line)
	count, _ := strconv.Atoi(getFirstLast(line))

	return count
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error", err)
	}

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		count := CalculateLineCount(scanner.Text())
		fmt.Println(scanner.Text(), "Result", result+count, "=", result ,"+",count)
		result += count
	}

	fmt.Println("Result", result) //55686
}
