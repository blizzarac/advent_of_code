package main

import (
	"golang.org/x/text/search"
	"golang.org/x/text/language"
	"fmt"
	"bufio"
	"os"
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

		if len(numOnly) == 0 {
			return ""
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

func SubstringIndex(str string, substr string) (int, bool) {
    m := search.New(language.English, search.IgnoreCase)
    start, _ := m.IndexString(str, substr)
    if start == -1 {
        return start, false
    }
    return start, true
}

func ReplaceStringNumbers(line string) string {
	fmt.Println("ReplaceStringNumbers", line)

	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

	resultList := make([]string, len(line))

	marker := 0
	for i := 1; i < len(line) +1; i++ {
		for _, number := range numbers {
			fmt.Println("Checking sub", line[marker:i], number)
			idx, ok := SubstringIndex(line[marker:i], number)
			if ok {
				marker+=idx+1
				fmt.Println("Found", number, "at", idx)
				resultList = append(resultList, number)
			}
		}
	}

	fmt.Println("resultList", resultList)

	for i := 0; i < len(resultList) ; i++ {
		resultList[i] = replaceStringNumber(resultList[i])
	}

	fmt.Println("resultList", resultList)

	return strings.Join(resultList, "")
}

func CalculateLineCount(line string) int {
	line = ReplaceStringNumbers(line)
	count, _ := strconv.Atoi(getFirstLast(line))

	return count
}

func CalculateLines(scanner *bufio.Scanner) int {
	result := 0
	for scanner.Scan() {
		result += CalculateLineCount(scanner.Text())
	}
	return result
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	result := CalculateLines(scanner)

	fmt.Println("Result", result) //55686
}
