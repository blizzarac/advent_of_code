package main


import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)

func rangeToList(s string) []int {
	var list []int
	start := strings.Split(s, "-")[0]
	end := strings.Split(s, "-")[1]

	startNum, _ := strconv.Atoi(start)
	endNum, _ := strconv.Atoi(end)

	for i := startNum; i <= endNum; i++ {
		list = append(list, i)
	}
	return list
}

func intersectionChecker(ass1 []int, ass2 []int) bool {
	if ass1[0] >= ass2[0] && ass1[len(ass1)-1] <= ass2[len(ass2)-1] {
		return true
	}

	if ass2[0] >= ass1[0] && ass2[len(ass2)-1] <= ass1[len(ass1)-1] {
		return true
	}

	return false
}

func intersectionPartialChecker(ass1 []int, ass2 []int) bool {
	if ass1[0] >= ass2[0] && ass1[0] <= ass2[len(ass2)-1] {
		return true
	}

	if ass2[0] >= ass1[0] && ass2[0] <= ass1[len(ass1)-1] {
		return true
	}

	return false
}

func main() {

	// Open file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	// Read file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		assigments := strings.Split(scanner.Text(), ",")

		fmt.Println(assigments)

		ass1 := rangeToList(assigments[0])
		ass2 := rangeToList(assigments[1])

		fmt.Println(ass1, ass2)

		if intersectionPartialChecker(ass1, ass2) {
			result++
		}

	}

	fmt.Println(result)
}
