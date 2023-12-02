package main


import (
	"sort"
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
)

func Map[T any, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func mapReduce(line string) int {
	words := strings.Split(line, ",")

	// Map
	numbers := Map(words, func(s string) int {
		num, _ := strconv.Atoi(s)
		return num
	})

	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}


func main() {
	// Read the file, line by line
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error reading file")
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }

	lines := strings.Split(string(content), ";")

	calList := make([]int, 0)
	for _, line := range lines {
		cur := mapReduce(line)
		calList = append(calList, cur)
	}

	fmt.Println(calList)
	sort.Sort(sort.Reverse(sort.IntSlice(calList)))
	fmt.Println(calList[0]+calList[1]+calList[2])
}
