package main


import (
	"testing"
	"log"
)

func Test_CalculateScenicScore_Edge_Top_Left(t *testing.T) {
	sut := [][]int {
		{3,0,3,7,3},
		{2,5,5,1,2},
		{6,5,3,3,2},
		{3,3,5,4,9},
		{3,5,5,9,0},
	}

	result := calculateScenicScore(sut, 0, 0)

	if result == 4{
		log.Println("Test_CalculateScenicScore_Edge_Top_Left passed")
	} else {
		t.Error("Test_CalculateScenicScore_Edge_Top_Left failed", result)
	}
}

func Test_CalculateScenicScore_Edge_Top_Right(t *testing.T) {
	sut := [][]int {
		{3,0,3,7,3},
		{2,5,5,1,2},
		{6,5,3,3,2},
		{3,3,5,4,9},
		{3,5,5,9,0},
	}

	result := calculateScenicScore(sut, 4, 0)

	if result == 4{
		log.Println("Test_CalculateScenicScore_Edge_Top_Right passed")
	} else {
		t.Error("Test_CalculateScenicScore_Edge_Top_Right failed", result)
	}
}

func Test_CalculateScenicScore_Edge_Bottom_Left(t *testing.T) {
	sut := [][]int {
		{3,0,3,7,3},
		{2,5,5,1,2},
		{6,5,3,3,2},
		{3,3,5,4,9},
		{3,5,5,9,0},
	}

	result := calculateScenicScore(sut, 0, 4)

	if result == 4{
		log.Println("Test_CalculateScenicScore_Edge_Bottom_Left passed")
	} else {
		t.Error("Test_CalculateScenicScore_Edge_Bottom_Left failed", result)
	}
}

func Test_CalculateScenicScore_Edge_Bottom_Right(t *testing.T) {
	sut := [][]int {
		{3,0,3,7,3},
		{2,5,5,1,2},
		{6,5,3,3,2},
		{3,3,5,4,9},
		{3,5,5,9,0},
	}

	result := calculateScenicScore(sut, 4, 4)

	if result == 2{
		log.Println("Test_CalculateScenicScore_Edge_Bottom_Right passed")
	} else {
		t.Error("Test_CalculateScenicScore_Edge_Bottom_Right failed", result)
	}
}

func Test_ScenicScore_Integration_1(t *testing.T) {
	sut := [][]int {
		{3,0,3,7,3},
		{2,5,5,1,2},
		{6,5,3,3,2},
		{3,3,5,4,9},
		{3,5,5,9,0},
	}

	result, debugGrid := findMaxScenicScore(sut)

	printGrid(sut)
	printGrid(debugGrid)

	if result == 8{
		log.Println("Test_ScenicScore_Integration_1 passed")
	} else {
		t.Error("Test_ScenicScore_Integration_1 failed", result)
	}
}

func Test_findUncoveredInRow_integ_row_2(t *testing.T) {
	sut := []int {2,5,5,1,2}

	result := findUncoveredInRow(sut)
	if result == 3 {
		t.Log("Test_findUncoveredInRow_integ_row_2 passed")
	} else {
		t.Error("Test_findUncoveredInRow_integ_row_2 failed", result)
	}

}

func Test_findUncoveredInRow_integ_row_3(t *testing.T) {
	sut := []int {6,5,3,3,2}

	result := findUncoveredInRow(sut)
	if result == 1{
		t.Log("Test_findUncoveredInRow_integ_row_3 passed")
	} else {
		t.Error("Test_findUncoveredInRow_integ_row_3 failed", result)
	}

}

func Test_findUncoveredInRow_integ_row_4(t *testing.T) {
	sut := []int {3,3,5,4,9}

	result := findUncoveredInRow(sut)
	if result == 1 {
		t.Log("Test_findUncoveredInRow_integ_row_4 passed")
	} else {
		t.Error("Test_findUncoveredInRow_integ_row_4 failed", result)
	}

}
