package main


import(
	"testing"
)




func Test_Example1(t *testing.T) {
	sut := []byte("bvwbjplbgvbhsrlpgdmjqwftvncz")

	result := findMarkerPosition(sut)

	if result == 5 {
		t.Log("Test_Example1 passed")
	} else {
		t.Errorf("Test_Example1 failed. Expected 5, got %d", result)
	}
}

func Test_Example2(t *testing.T) {
	sut := []byte( "nppdvjthqldpwncqszvftbrmjlhg")

	result := findMarkerPosition(sut)

	if result == 6 {
		t.Log("Test_Example2 passed")
	} else {
		t.Errorf("Test_Example2 failed. Expected 5, got %d", result)
	}
}

func Test_Example3(t *testing.T) {
	sut := []byte( "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")

	result := findMarkerPosition(sut)

	if result == 10 {
		t.Log("Test_Example3 passed")
	} else {
		t.Errorf("Test_Example3 failed. Expected 10, got %d", result)
	}
}

func Test_Example4(t *testing.T) {
	sut := []byte( "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")

	result := findMarkerPosition(sut)

	if result == 11 {
		t.Log("Test_Example4 passed")
	} else {
		t.Errorf("Test_Example4 failed. Expected 11, got %d", result)
	}
}
