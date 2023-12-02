package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Sign int
type Outcome int

const (
	Rock Sign = iota
	Paper
	Scissors
)

const (
	Win Outcome = iota
	Lose
	Draw
)

func signValue(s Sign) int {
	switch s {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	return 0
}

func translate(s string) (Sign) {
	switch s {
	case "A" , "X":
		return Rock
	case "B" ,  "Y":
		return Paper
	case "C" , "Z":
		return Scissors
	}
	return 0
}

func translateOutcome(s string) (Outcome) {
	 switch s {
		case "X":
			return Lose
		case "Y":
			return Draw
		case "Z":
			return Win
		}

	return 0
}

func outcome(sign1 Sign, sign2 Sign) int {
	if sign1 == Rock && sign2 == Scissors {
		return 1
	} else if sign1 == Scissors && sign2 == Paper {
		return 1
	} else if sign1 == Paper && sign2 == Rock {
		return 1
	} else if sign1 == Scissors && sign2 == Rock {
		return 2
	} else if sign1 == Paper && sign2 == Scissors {
		return 2
	} else if sign1 == Rock && sign2 == Paper {
		return 2
	} else {
		return 0
	}
}

func calculateSign(sign1 Sign, result Outcome) (Sign) {
	switch sign1 {
	case Rock:
		if result == Win {
			return Paper
		} else if result == Lose {
			return Scissors
		} else if result == Draw {
			return Rock
		}
	case Paper:
		if result == Win {
			return Scissors
		} else if result == Lose {
			return Rock
		} else if result == Draw {
			return Paper
		}
	case Scissors:
		if result == Win {
			return Rock
		} else if result == Lose {
			return Paper
		} else if result == Draw {
			return Scissors
		}
	}
	return 0
}

func calculatePoints(sign1 Sign, sign2 Sign) (int) {
	result := signValue(sign2)

	outcome := outcome(sign1, sign2)
	switch outcome {
	case 1:
		result += 0
	case 2:
		result += 6
	case 0:
		result += 3
	}

	return result
}

func main() {

	// Read file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	totalPoints := 0
	for _, line := range lines {
		fmt.Println(line)
		signs := strings.Split(line, " ")
		sign1 := translate(signs[0])
		result := translateOutcome(signs[1])

		sign2 := calculateSign(sign1, result)

		totalPoints += calculatePoints(sign1, sign2)
	}

	fmt.Println(totalPoints)
}

