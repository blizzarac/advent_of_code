package main

import (
	"testing"
	"log"
	"os"
	"io/ioutil"
	"strings"
)

func Test_Chunk_String_length(t *testing.T) {
	chunk := chunkString("    [M]             [P] [L] [B] [J]", 4)

	if len(chunk) == 9 {
		log.Println("Test_Chunk_String_length: OK")
	} else {
		t.Error("Test_Chunk_String_length: FAIL")
	}
}

func Test_assembleChunksInMap(t *testing.T) {
	ship := map[int]stack{
		0: {"[X]"},
		1: {"[A]", "[B]", "[C]"},
		2: {"[D]", "[E]", "[F]"},
		3: {"[G]", "[H]", "[I]"},
	}

	chunks := chunkString("    [M]             [P] [L] [B] [J]", 4)

	for i, chunk := range chunks {
		ship = assembleChunksInMap(ship, i,chunk)
	}


	if len(ship) == 8 {
		log.Println("Test_assembleChunksInMap: OK")
	} else {
		t.Error("Test_assembleChunksInMap: FAIL", len(ship))
	}
}

func TestReadShipConfig(t *testing.T) {
	file, _ := os.Open("input.txt")
	defer file.Close()

	content, _ := ioutil.ReadAll(file)
	lines := strings.Split(string(content), "\n")

	ship := readShipConfig(lines, 8)

	if len(ship) == 9 {
		log.Println("TestReadShipConfig: OK")
	} else {
		t.Error("TestReadShipConfig: FAIL")
	}
}



func TestMove_single_move(t *testing.T) {
	ship := map[int]stack{
		0: {"A", "B", "C"},
		1: {"D", "E", "F"},
		2: {"G", "H", "I"},
	}

	ship = move(ship, 1, 2)

	if ship[0][0]== "A" && ship[0][1]== "B" && ship[0][2]== "C" &&
		ship[1][0]== "D" && ship[1][1]== "E" &&
		ship[2][0]== "G" && ship[2][1]== "H" && ship[2][2]== "I" && ship[2][3]== "F" {
		log.Println("TestMove_single_move: OK")
	} else {
		t.Error("TestMove_single_move: FAIL")
	}
}

func TestMove_multiple_moves_9001(t *testing.T) {
	ship := map[int]stack{
		0: {"A", "B", "C"},
		1: {"D", "E", "F"},
		2: {"G", "H", "I"},
	}

	printShip("TestMove_multiple_moves_9001: before", ship)
	ship = multiMoveAtOnce(ship, 0, 1, 2)
	printShip("TestMove_multiple_moves_9001: after", ship)

	if ship[0][0]== "A" &&
		ship[1][0]== "D" && ship[1][1]== "E" && ship[1][2]== "F" && ship[1][3]== "B" && ship[1][4]== "C" &&
		ship[2][0]== "G" && ship[2][1]== "H" && ship[2][2]== "I" {
		log.Println("TestMove_multiple_moves_9001: OK")
	} else {
		t.Error("TestMove_multiple_moves_9001: FAIL")
	}
}

func TestMove_multiple_moves(t *testing.T) {
	ship := map[int]stack{
		0: {"A", "B", "C"},
		1: {"D", "E", "F"},
		2: {"G", "H", "I"},
	}

	ship = multiMove(ship, 0, 1, 2)

	if ship[0][0]== "A" &&
		ship[1][0]== "D" && ship[1][1]== "E" && ship[1][2]== "F" && ship[1][3]== "C" && ship[1][4]== "B" &&
		ship[2][0]== "G" && ship[2][1]== "H" && ship[2][2]== "I" {
		log.Println("TestMove_single_move: OK")
	} else {
		t.Error("TestMove_single_move: FAIL")
	}
}

func TestCalcTopContainers_Check_Correct_Endings(t *testing.T) {
	ship := map[int]stack{
		1: {"A", "B", "C"},
		2: {"D", "E", "F"},
		3: {"G", "H", "I"},
	}

	result := calcTopContainers(ship)
	if result == "CFI" {
		log.Println("TestCalcTopContainers: OK")
	} else {
		t.Error("TestCalcTopContainers: FAIL", result)
	}
}

func TestIntegration(t * testing.T) {
	file, _ := os.Open("input_integration.txt")
	defer file.Close()

	content, _ := ioutil.ReadAll(file)
	lines := strings.Split(string(content), "\n")

	ship := readShipConfig(lines, 3)
//	printShip("Integration test: readShipConfig", ship)

	if ship[0][0]== "[Z]" && ship[0][1]== "[N]" &&
		ship[1][0]== "[M]" && ship[1][1]== "[C]" && ship[1][2]== "[D]" &&
		ship[2][0]== "[P]" {
		log.Println("TestIntegration: readShipConfig OK")
	} else {
		t.Error("TestIntegration: readShipConfig FAIL")
	}

	instructions :=	readInstructions(lines[3:])

	for _, instruction := range instructions {
		ship = multiMove(ship, instruction.from, instruction.to, instruction.numOfMoves)
	}

//	printShip("Integration test: after move", ship)
	if ship[0][0]== "[C]" &&
		ship[1][0]== "[M]" &&
		ship[2][0]== "[P]" && ship[2][1]== "[D]" && ship[2][2]== "[N]" && ship[2][3]== "[Z]" {
		log.Println("TestIntegration: readShipConfig OK")
	} else {
		t.Error("TestIntegration: readShipConfig FAIL")
	}

	result := calcTopContainers(ship)
	if result == "CMZ" {
		log.Println("TestIntegration: calcTopContainers OK")
	} else {
		t.Error("TestIntegration: calcTopContainers FAIL", result)
	}
}









































