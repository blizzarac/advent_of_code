package main



import (
	"testing"
	"fmt"
)




func TestSignValue_Rock(t *testing.T) {
	result := signValue(Rock)
	if result == 1 {
		fmt.Println("TestSignValue_Rock PASS")
		return
	} else {
		fmt.Println("TestSignValue_Rock FAIL")
	}
}

func TestSignValue_Paper(t *testing.T) {
	result := signValue(Paper)
	if result == 2 {
		fmt.Println("TestSignValue_Paper PASS")
		return
	} else {
		fmt.Println("TestSignValue_Paper FAIL")
	}
}

func TestSignValue_Scissors(t *testing.T) {
	result := signValue(Scissors)
	if result == 3 {
		fmt.Println("TestSignValue_Scissors PASS")
		return
	} else {
		fmt.Println("TestSignValue_Scissors FAIL")
	}
}


func TestTranslate_A(t *testing.T) {
	result := translate("A")
	if result == Rock {
		fmt.Println("TestTranslate_A PASS")
		return
	} else {
		fmt.Println("TestTranslate_A FAIL")
	}
}

func TestTranslate_B(t *testing.T) {
	result := translate("B")
	if result == Paper {
		fmt.Println("TestTranslate_B PASS")
		return
	} else {
		fmt.Println("TestTranslate_B FAIL")
	}
}

func TestTranslate_C(t *testing.T) {
	result := translate("C")
	if result == Scissors {
		fmt.Println("TestTranslate_C PASS")
		return
	} else {
		fmt.Println("TestTranslate_C FAIL")
	}
}

func TestTranslate_X(t *testing.T) {
	result := translate("X")
	if result == Rock {
		fmt.Println("TestTranslate_X PASS")
		return
	} else {
		fmt.Println("TestTranslate_X FAIL")
	}
}

func TestTranslate_Y(t *testing.T) {
	result := translate("Y")
	if result == Paper {
		fmt.Println("TestTranslate_Y PASS")
		return
	} else {
		fmt.Println("TestTranslate_Y FAIL")
	}
}

func TestTranslate_Z(t *testing.T) {
	result := translate("Z")
	if result == Scissors {
		fmt.Println("TestTranslate_Z PASS")
		return
	} else {
		fmt.Println("TestTranslate_Z FAIL")
	}
}

func TestOutcome_RockRock(t *testing.T) {
	result := outcome(Rock, Rock)
	if result == 0 {
		fmt.Println("TestOutcome_RockRock PASS")
		return
	} else {
		fmt.Println("TestOutcome_RockRock FAIL")
	}
}

func TestOutcome_RockPaper(t *testing.T) {
	result := outcome(Rock, Paper)
	if result == 2 {
		fmt.Println("TestOutcome_RockPaper PASS")
		return
	} else {
		fmt.Println("TestOutcome_RockPaper FAIL")
	}
}

func TestOutcome_RockScissors(t *testing.T) {
	result := outcome(Rock, Scissors)
	if result == 1 {
		fmt.Println("TestOutcome_RockScissors PASS")
		return
	} else {
		fmt.Println("TestOutcome_RockScissors FAIL")
	}
}

func TestOutcome_PaperRock(t *testing.T) {
	result := outcome(Paper, Rock)
	if result == 1 {
		fmt.Println("TestOutcome_PaperRock PASS")
		return
	} else {
		fmt.Println("TestOutcome_PaperRock FAIL")
	}
}

func TestOutcome_PaperPaper(t *testing.T) {
	result := outcome(Paper, Paper)
	if result == 0 {
		fmt.Println("TestOutcome_PaperPaper PASS")
		return
	} else {
		fmt.Println("TestOutcome_PaperPaper FAIL")
	}
}

func TestOutcome_PaperScissors(t *testing.T) {
	result := outcome(Paper, Scissors)
	if result == 2 {
		fmt.Println("TestOutcome_PaperScissors PASS")
		return
	} else {
		fmt.Println("TestOutcome_PaperScissors FAIL")
	}
}

func TestOutcome_ScissorsRock(t *testing.T) {
	result := outcome(Scissors, Rock)
	if result == 2 {
		fmt.Println("TestOutcome_ScissorsRock PASS")
		return
	} else {
		fmt.Println("TestOutcome_ScissorsRock FAIL")
	}
}

func TestOutcome_ScissorsPaper(t *testing.T) {
	result := outcome(Scissors, Paper)
	if result == 1 {
		fmt.Println("TestOutcome_ScissorsPaper PASS")
		return
	} else {
		fmt.Println("TestOutcome_ScissorsPaper FAIL")
	}
}

func TestOutcome_ScissorsScissors(t *testing.T) {
	result := outcome(Scissors, Scissors)
	if result == 0 {
		fmt.Println("TestOutcome_ScissorsScissors PASS")
		return
	} else {
		fmt.Println("TestOutcome_ScissorsScissors FAIL")
	}
}

func TestCalculatePoints_RockRock(t *testing.T) {
	result := calculatePoints(Rock, Rock)
	if result == 4 {
		fmt.Println("TestCalculatePoints_1 PASS")
		return
	} else {
		fmt.Println("TestCalculatePoints_1 FAIL")
	}
}

func TestCalculatePoints_RockPaper(t *testing.T) {
	result := calculatePoints(Rock, Paper)
	if result == 8 {
		fmt.Println("TestCalculatePoints_2 PASS")
		return
	} else {
		fmt.Println("TestCalculatePoints_2 FAIL")
	}
}

func TestCalculatePoints_RockScissors(t *testing.T) {
	result := calculatePoints(Rock, Scissors)
	if result == 3 {
		fmt.Println("TestCalculatePoints_3 PASS")
		return
	} else {
		fmt.Println("TestCalculatePoints_3 FAIL")
	}
}

func TestCalculatePoints_PaperRock(t *testing.T) {
	result := calculatePoints(Paper, Rock)
	if result == 1 {
		fmt.Println("TestCalculatePoints_4 PASS")
		return
	} else {
		fmt.Println("TestCalculatePoints_4 FAIL")
	}
}

func TestCalculatePoints_PaperPaper(t *testing.T) {
	result := calculatePoints(Paper, Paper)
	if result == 5 {
		fmt.Println("TestCalculatePoints_5 PASS")
		return
	} else {
		fmt.Println("TestCalculatePoints_5 FAIL")
	}
}

func TestCalculatePoints_PaperScissors(t *testing.T) {
	result := calculatePoints(Paper, Scissors)
	if result == 9 {
		fmt.Println("TestCalculatePoints_6 PASS")
		return
	} else {
		fmt.Println("TestCalculatePoints_6 FAIL")
	}
}


func TestIntegration_Case1(t *testing.T) {
	result := calculatePoints(translate("A"), translate("Y"))

	if result == 8 {
		fmt.Println("TestIntegration_Case1 PASS")
		return
	} else {
		fmt.Println("TestIntegration_Case1 FAIL")
	}
}

func TestIntegration_Case2(t *testing.T) {
	result := calculatePoints(translate("B"), translate("X"))

	if result == 1 {
		fmt.Println("TestIntegration_Case2 PASS")
		return
	} else {
		fmt.Println("TestIntegration_Case2 FAIL")
	}
}

func TestIntegration_Case3(t *testing.T) {
	result := calculatePoints(translate("C"), translate("Z"))

	if result == 6 {
		fmt.Println("TestIntegration_Case3 PASS")
		return
	} else {
		fmt.Println("TestIntegration_Case3 FAIL")
	}
}
