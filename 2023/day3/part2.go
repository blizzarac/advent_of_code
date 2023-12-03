package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/samber/lo"
	"strings"
	"strconv"
)

type MachinePart struct {
	number string
	row int
	column_indexes []int
	validated bool
	gear_row int
	gear_column int
}



func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	plan := LoadPlan(scanner)
	PrintPlan(plan)

	machine_parts := FindMachineParts(plan)
	result := 0

	for idx, _ := range machine_parts {
		ValidateMachinePart(&machine_parts[idx], plan)
	}

	machine_parts = lo.Filter(machine_parts, func(x MachinePart, index int) bool {
		return x.validated
	})

	gearSetup := make(map[string][]MachinePart)

	for idx1, machine_part_1 := range machine_parts {
		for idx2, machine_part_2 := range machine_parts {
			if idx1 == idx2 {
				continue
			}
			if machine_part_1.gear_row == machine_part_2.gear_row && machine_part_1.gear_column == machine_part_2.gear_column {
				gearSetup[string(machine_part_1.gear_row)+string(machine_part_1.gear_column)] = append(gearSetup[string(machine_part_1.gear_row)+string(machine_part_1.gear_column)], machine_part_1)
				gearSetup[string(machine_part_1.gear_row)+string(machine_part_1.gear_column)] = append(gearSetup[string(machine_part_1.gear_row)+string(machine_part_1.gear_column)], machine_part_2)
				break
			}
		}
	}

	for key, value := range gearSetup {
    	local_result := 1
		fmt.Println("Key:", key)
		value = lo.UniqBy(value, func(x MachinePart) string {
			return x.number
		})

		if len(value) == 2 {
			for _, machine_part := range value {
				vl, _ := strconv.Atoi(machine_part.number)
				fmt.Println("Value:", vl)
				local_result *= vl
				fmt.Println("Local result:", local_result)
			}
		} else {
			fmt.Println("Too Many")
		}

		result += local_result
	}

	fmt.Println("Result", result)
}

func ValidateMachinePart(machine_part *MachinePart, plan [][]string) bool {

	// Check left
	if machine_part.column_indexes[0] > 0 {
//		fmt.Println("Check left", machine_part.number, machine_part.row, machine_part.column_indexes[0] - 1, plan[machine_part.row][machine_part.column_indexes[0] - 1])
		if plan[machine_part.row][machine_part.column_indexes[0] - 1]  == "*" {
			machine_part.gear_row = machine_part.row
			machine_part.gear_column = machine_part.column_indexes[0] - 1
			machine_part.validated = true
			return true
		}
	}

	// Check upper left
	if machine_part.column_indexes[0] > 0 && machine_part.row > 0 {
//		fmt.Println("Check upper left", machine_part.number, machine_part.row - 1, machine_part.column_indexes[0] - 1, plan[machine_part.row - 1][machine_part.column_indexes[0] - 1])
		if plan[machine_part.row - 1][machine_part.column_indexes[0] - 1]  == "*" {
			machine_part.gear_row = machine_part.row - 1
			machine_part.gear_column = machine_part.column_indexes[0] - 1

			machine_part.validated = true
			return true
		}
	}

	// Check right
	if machine_part.column_indexes[len(machine_part.column_indexes) - 1] < len(plan[machine_part.row]) - 1 {
//		fmt.Println("Check right", machine_part.number, machine_part.row, machine_part.column_indexes[len(machine_part.column_indexes) - 1] + 1, plan[machine_part.row][machine_part.column_indexes[len(machine_part.column_indexes) - 1] + 1])
		if plan[machine_part.row][machine_part.column_indexes[len(machine_part.column_indexes) - 1] + 1]  == "*" {
				machine_part.gear_row = machine_part.row
				machine_part.gear_column = machine_part.column_indexes[len(machine_part.column_indexes) - 1] + 1

			machine_part.validated = true
			return true
		}
	}

	// Check upper right
	if machine_part.column_indexes[len(machine_part.column_indexes) - 1] < len(plan[machine_part.row]) - 1 && machine_part.row > 0 {
//		fmt.Println("Check upper right", machine_part.number, machine_part.row - 1, machine_part.column_indexes[len(machine_part.column_indexes) - 1] + 1, plan[machine_part.row - 1][machine_part.column_indexes[len(machine_part.column_indexes) - 1] + 1])
		if plan[machine_part.row - 1][machine_part.column_indexes[len(machine_part.column_indexes) - 1] + 1]  == "*" {
				machine_part.gear_row = machine_part.row - 1
				machine_part.gear_column = machine_part.column_indexes[len(machine_part.column_indexes) - 1] + 1

			machine_part.validated = true
			return true
		}
	}

	// Check up
	if machine_part.row > 0 {
//		fmt.Println("Check up", machine_part.number, machine_part.row - 1, machine_part.column_indexes)
		for _, column_index := range machine_part.column_indexes {
			if plan[machine_part.row - 1][column_index]  == "*" {
				//fmt.Println("Check up", machine_part.number, machine_part.row - 1, column_index, plan[machine_part.row - 1][column_index])
					machine_part.gear_row = machine_part.row - 1
					machine_part.gear_column = column_index

				machine_part.validated = true
				return true
			}
		}
	}

	// Check lower left
	if machine_part.column_indexes[0] > 0 && machine_part.row < len(plan) - 1 {
//		fmt.Println("Check lower left", machine_part.number, machine_part.row + 1, machine_part.column_indexes[0] - 1, plan[machine_part.row + 1][machine_part.column_indexes[0] - 1])
		if plan[machine_part.row + 1][machine_part.column_indexes[0] - 1]  == "*" {
				machine_part.gear_row = machine_part.row + 1
				machine_part.gear_column = machine_part.column_indexes[0] - 1

			machine_part.validated = true
			return true
		}
	}

	// Check down
	if machine_part.row < len(plan) - 1 {
//		fmt.Println("Check down", machine_part.number, machine_part.row + 1, machine_part.column_indexes)
		for _, column_index := range machine_part.column_indexes {
			if plan[machine_part.row + 1][column_index]  == "*" {
				//fmt.Println("Check down", machine_part.number, machine_part.row + 1, column_index, plan[machine_part.row + 1][column_index])
					machine_part.gear_row = machine_part.row + 1
					machine_part.gear_column = column_index

				machine_part.validated = true
				return true
			}
		}
	}

	// Check lower right
	if machine_part.column_indexes[len(machine_part.column_indexes) - 1] < len(plan[machine_part.row]) - 1 && machine_part.row < len(plan) - 1 {
//		fmt.Println("Check lower right", machine_part.number, machine_part.row + 1, machine_part.column_indexes[len(machine_part.column_indexes) - 1] + 1, plan[machine_part.row + 1][machine_part.column_indexes[len(machine_part.column_indexes) - 1] + 1])
		if plan[machine_part.row + 1][machine_part.column_indexes[len(machine_part.column_indexes) - 1] + 1]  == "*" {
				machine_part.gear_row = machine_part.row + 1
				machine_part.gear_column = machine_part.column_indexes[len(machine_part.column_indexes) - 1] + 1

			machine_part.validated = true
			return true
		}
	}

	return false
}

func FindMachineParts(plan [][]string) []MachinePart {
	var machine_parts []MachinePart

	var machine_part *MachinePart

	for row_index, row := range plan {
		for column_index, _ := range row {
			_, err := strconv.Atoi(plan[row_index][column_index])
			if err == nil {
				if machine_part == nil {
					machine_part = &MachinePart{number: plan[row_index][column_index], row: row_index, column_indexes: []int{column_index}}
				} else {
					machine_part.number += plan[row_index][column_index]
					machine_part.column_indexes = append(machine_part.column_indexes, column_index)
				}
			} else {
				if machine_part != nil {
					machine_parts = append(machine_parts, *machine_part)
					machine_part = nil
				}
			}
		}
	}

	return machine_parts
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

func PrintMachineParts(machine_part []MachinePart) {
	for _, part := range machine_part {
		PrintMachinePart(part)
	}
}

func PrintMachinePart(machine_part MachinePart) {
	fmt.Println("Machine part", machine_part.number, "row", machine_part.row,
		"columns", machine_part.column_indexes, "validated", machine_part.validated,
		"gear row", machine_part.gear_row, "gear column", machine_part.gear_column)
}

func PrintPlan(plan [][]string) {
	for _, row := range plan {
		fmt.Println(row)
	}
}
