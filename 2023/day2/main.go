package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"regexp"
)

type game struct {
	number int
	subsets []string
}

type subset struct {
	blue int
	green int
	red int
}

func ParseGames(scanner *bufio.Scanner) []game{
	var games []game

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, " ", "", -1)

		cubes := strings.Split(line, ":")
		subsets := strings.Split(cubes[1], ";")

		regex_game := regexp.MustCompile(`Game(?P<game>\d+):`)
		gameIndex := regex_game.SubexpIndex("game")
		matches := regex_game.FindStringSubmatch(line)

		number := 0
		if len(matches) > 0 {
			number,_ = strconv.Atoi(matches[gameIndex])
		} else {
			fmt.Println("No game number found in line", line)
		}

		games = append(games, game{number, subsets})
	}

	return games
}

func checkIfGamePossible(game game, blue int, green int, red int) bool {

	regex_red := regexp.MustCompile(`(?P<count>\d+)red`)
	regex_green := regexp.MustCompile(`(?P<count>\d+)green`)
	regex_blue := regexp.MustCompile(`(?P<count>\d+)blue`)

	redIndex := regex_red.SubexpIndex("count")
	greenIndex := regex_green.SubexpIndex("count")
	blueIndex := regex_blue.SubexpIndex("count")

	for _, subset := range game.subsets {
		matches_red := regex_red.FindStringSubmatch(subset)
		matches_green := regex_green.FindStringSubmatch(subset)
		matches_blue := regex_blue.FindStringSubmatch(subset)
		if len(matches_blue) > 0 {
			blue_hint, _ := strconv.Atoi(matches_blue[blueIndex])
			if blue_hint > blue {
				return false
			}
		}

		if len(matches_red) > 0 {
			red_hint, _ := strconv.Atoi(matches_red[redIndex])
			if red_hint > red {
				return false
			}
		}

		if len(matches_green) > 0 {
			green_hint, _ := strconv.Atoi(matches_green[greenIndex])
			if green_hint > green {
				return false
			}
		}

	}

	return true
}

func findMinCubePower(game game) int {
	regex_red := regexp.MustCompile(`(?P<count>\d+)red`)
	regex_green := regexp.MustCompile(`(?P<count>\d+)green`)
	regex_blue := regexp.MustCompile(`(?P<count>\d+)blue`)

	redIndex := regex_red.SubexpIndex("count")
	greenIndex := regex_green.SubexpIndex("count")
	blueIndex := regex_blue.SubexpIndex("count")

	min_red :=0
	min_green :=0
	min_blue := 0

	fmt.Println("Game", game.number)

	for _, subset := range game.subsets {
		fmt.Println("Subset", subset)
		matches_red := regex_red.FindStringSubmatch(subset)
		matches_green := regex_green.FindStringSubmatch(subset)
		matches_blue := regex_blue.FindStringSubmatch(subset)
		if len(matches_blue) > 0 {
			blue_hint, _ := strconv.Atoi(matches_blue[blueIndex])
			if blue_hint > min_blue {
				min_blue = blue_hint
			}
		}

		if len(matches_red) > 0 {
			red_hint, _ := strconv.Atoi(matches_red[redIndex])
			if red_hint > min_red {
				min_red = red_hint
			}
		}

		if len(matches_green) > 0 {
			green_hint, _ := strconv.Atoi(matches_green[greenIndex])
			if green_hint > min_green {
				min_green = green_hint
			}
		}

	}

	fmt.Println("Min red", min_red, "Min green", min_green, "Min blue", min_blue)

	return min_red * min_green * min_blue
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	games := ParseGames(scanner)
	result := 0

	fmt.Println("Found", len(games), "games")

	min_cube_power := 0

	for _, game := range games {
		if checkIfGamePossible(game, 14, 13, 12) {
			result += game.number
		}

		min_cube_power += findMinCubePower(game)
	}

	fmt.Println("Result", result)

	fmt.Println("Min cube power", min_cube_power)
}
