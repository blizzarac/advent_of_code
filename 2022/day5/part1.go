package main

import (
	"strconv"
	"regexp"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"slices"
)

// Stack implementation
type stack []string

func (s stack) Push(v string) stack {
    return append(s, v)
}

func (s stack) Pop() (stack, string) {
    if len(s) == 0 {
		return s, ""
	}

	l := len(s)
    return  s[:l-1], s[l-1]
}

type moveInstruction struct {
    numOfMoves int
    from int
	to int
}

func chunkString(s string, chunkSize int) []string {
	// Initialize a slice to hold the chunks
	var chunks []string

	// Loop through the string, creating chunks of the specified size
	for chunkSize < len(s) {
		chunks = append(chunks, s[:chunkSize]) // Append chunk to the slice
		s = s[chunkSize:]                      // Move to the next chunk
	}

	// Append any remaining characters as a chunk
	if len(s) > 0 {
		chunks = append(chunks, s)
	}

	return chunks
}

func assembleChunksInMap(ship map[int]stack, idx int, chunk string) map[int]stack {
	cleanChunk := strings.ReplaceAll(chunk, " ", "")
	if len(cleanChunk) > 0 {
		ship[idx] = ship[idx].Push(strings.ReplaceAll(chunk, " ", ""))
	}

	return ship
}

func readShipConfig(lines []string, lineCount int) map[int]stack {
	var ship = map[int]stack{}

	// Read Stack
	shipConfig := lines[0:lineCount]
	slices.Reverse(shipConfig)

	for _, line := range shipConfig {
		chunks := chunkString(line, 4)
		for i, chunk := range chunks {
			ship = assembleChunksInMap(ship, i, chunk)
		}
	}

	return ship
}


func move(ship map[int]stack, from int, to int) map[int]stack {
	var movedContainer string
	ship[from], movedContainer = ship[from].Pop()
	ship[to] = ship[to].Push(movedContainer)
	return ship
}

func max(numbers map[int]stack) int {
    var maxNumber int
    for maxNumber = range numbers {
        break
    }
    for n := range numbers {
        if n > maxNumber {
            maxNumber = n
        }
    }
    return maxNumber
}

func calcTopContainers(ship map[int]stack) string {
	var shipConfig []string
	for i:=0; i <= max(ship); i++ {
		stack,ok := ship[i]
		if ok {
			_, topContainer := stack.Pop()
			shipConfig = append(shipConfig, topContainer)
		}
	}

	result := strings.Join(shipConfig, "")
	result = strings.ReplaceAll(result, " ", "")
	result = strings.ReplaceAll(result, "[", "")
	result = strings.ReplaceAll(result, "]", "")

	return result
}

func multiMove(ship map[int]stack, from int, to int, numOfMoves int) map[int]stack {
	for i := 0; i < numOfMoves; i++ {
		ship = move(ship, from, to)
	}
	return ship
}

func multiMoveAtOnce(ship map[int]stack, from int, to int, numOfMoves int) map[int]stack {
	var	movedContainers stack
	var movedContainer string

	for i := 0; i < numOfMoves; i++ {
		ship[from], movedContainer = ship[from].Pop()
		movedContainers = movedContainers.Push(movedContainer)
	}

	log.Println("Moved Containers: ", movedContainers)

	for i := 0; i < numOfMoves; i++ {
		movedContainers, movedContainer = movedContainers.Pop()
		ship[to] = ship[to].Push(movedContainer)
	}

	return ship
}

func printShip(tag string, ship map[int]stack) {
	log.Println(tag)
	for i:=0; i <= max(ship); i++ {
		stack,ok := ship[i]
		if ok {
			log.Println(i, stack)
		} else {
			log.Println(i, "[]")
		}
	}
}

func readInstructions(lines []string) []moveInstruction {
	var instructions []moveInstruction
	for _, line := range lines {
		re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
		matches := re.FindAllStringSubmatch(line, -1)
		if len(matches) > 0 {
			numOfMoves,_ := strconv.Atoi(matches[0][1])
			from,_ := strconv.Atoi(matches[0][2])
			to,_ := strconv.Atoi(matches[0][3])

			instruction := moveInstruction{numOfMoves, from-1, to-1}
			instructions = append(instructions, instruction)
		}
	}
	return instructions
}

func main() {
	// Open file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read whole file
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Split by lines
	lines := strings.Split(string(content), "\n")

	// Read ship configuration
	ship := readShipConfig(lines, 8)
	printShip("Ship from Input", ship)
	calcTopContainers(ship)


	instructions := readInstructions(lines[10:])

	craneVersion := 9001

	for _, instruction := range instructions {
		if craneVersion == 9001 {
			ship = multiMoveAtOnce(ship, instruction.from, instruction.to, instruction.numOfMoves)
		} else {
			ship = multiMove(ship, instruction.from, instruction.to, instruction.numOfMoves)
		}
	}

	result := calcTopContainers(ship)
	printShip("Ship after moves", ship)
	log.Println("Endresult", result)
}

