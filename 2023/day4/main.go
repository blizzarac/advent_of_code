package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)



func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	plan := LoadPlan(scanner)
	fmt.Println("Plan", plan)
	result := 0

	fmt.Println("Result", result)
}

func LoadPlan(scanner *bufio.Scanner) [][]string{
	fmt.Println("Loading plan")
	var plan [][]string

	for scanner.Scan() {
		line := scanner.Text()
		string_array := strings.Split(line, "")
		plan = append(plan, string_array)
	}

	return plan
}

