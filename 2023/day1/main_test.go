package main






import (
	"testing"
)

func TestCalculateLineCount(t *testing.T) {
	sut := "two1nine"

	result := CalculateLineCount(sut)
	if result == 29 {
		t.Log("TestCalculateLineCount passed")
	} else {
		t.Errorf("TestCalculateLineCount failed, expected %v, got %v", 29, result)
	}
}

func TestCalculateLineCount2(t *testing.T) {
	sut := "eightwothree"

	result := CalculateLineCount(sut)
	if result == 83 {
		t.Log("TestCalculateLineCount2 passed")
	} else {
		t.Errorf("TestCalculateLineCount2 failed, expected %v, got %v", 83, result)
	}
}

func TestCalculateLineCount3(t *testing.T) {
	sut := "abcone2threexyz"

	result := CalculateLineCount(sut)
	if result == 13 {
		t.Log("TestCalculateLineCount3 passed")
	} else {
		t.Errorf("TestCalculateLineCount3 failed, expected %v, got %v", 13, result)
	}
}

func TestCalculateLineCount4(t *testing.T) {
	sut := "xtwone3four"

	result := CalculateLineCount(sut)
	if result == 24 {
		t.Log("TestCalculateLineCount4 passed")
	} else {
		t.Errorf("TestCalculateLineCount4 failed, expected %v, got %v", 24, result)
	}
}

func TestCalculateLineCount5(t *testing.T) {
	sut := "54306"

	result := CalculateLineCount(sut)
	if result == 56 {
		t.Log("TestCalculateLineCount5 passed")
	} else {
		t.Errorf("TestCalculateLineCount5 failed, expected %v, got %v", 56, result)
	}
}

func TestCalculateLineCount6(t *testing.T) {
	sut := "8c"

	result := CalculateLineCount(sut)
	if result == 88 {
		t.Log("TestCalculateLineCount6 passed")
	} else {
		t.Errorf("TestCalculateLineCount6 failed, expected %v, got %v", 77, result)
	}
}

func TestCalculateLineCount6a(t *testing.T) {
	sut := "4"

	result := CalculateLineCount(sut)
	if result == 44 {
		t.Log("TestCalculateLineCount6a passed")
	} else {
		t.Errorf("TestCalculateLineCount6a failed, expected %v, got %v", 88, result)
	}
}

func TestCalculateLineCount7(t *testing.T) {
	sut := "4twoknkb39two"

	result := CalculateLineCount(sut)
	if result == 42 {
		t.Log("TestCalculateLineCount7 passed")
	} else {
		t.Errorf("TestCalculateLineCount7 failed, expected %v, got %v", 42, result)
	}
}

func TestMain(t *testing.T) {
	sut := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	result := 0
	for _, v := range sut {
		result += CalculateLineCount(v)
	}

	if result == 281 {
		t.Log("TestMain passed")
	} else {
		t.Errorf("TestMain failed, expected %v, got %v", 281, result)
	}
}

func TestMain2(t *testing.T) {
	sut := []string{
		"11abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	result := 0
	for _, v := range sut {
		result += CalculateLineCount(v)
	}

	if result == 142 {
		t.Log("TestMain2 passed")
	} else {
		t.Errorf("TestMain2 failed, expected %v, got %v", 142, result)
	}
}
