package main


import (
	"testing"
	"fmt"
)


func TestRangeToList_start(t *testing.T) {
	rangeString := "1-3"
	list := rangeToList(rangeString)
	if list[0] == 1 {
		fmt.Println("RangeToList(1-3) = 1")
	} else {
		t.Errorf("RangeToList(1-3) = %d; want 1", list[0])
	}
}


func TestRangeToList_mid(t *testing.T) {
	rangeString := "1-3"
	list := rangeToList(rangeString)
	if list[1] == 2 {
		fmt.Println("RangeToList(1-3) = 2")
	} else {
		t.Errorf("RangeToList(1-3) = %d; want 2", list[0])
	}
}

func TestRangeToList_end(t *testing.T) {
	rangeString := "1-3"
	list := rangeToList(rangeString)
	if list[2] == 3 {
		fmt.Println("RangeToList(1-3) = 3")
	} else {
		t.Errorf("RangeToList(1-3) = %d; want 3", list[0])
	}
}

func TestRangeToList_size(t *testing.T) {
	rangeString := "1-5"
	list := rangeToList(rangeString)
	if len(list) == 5 {
		fmt.Println("RangeToList(1-5) = 5")
	} else {
		t.Errorf("RangeToList(1-5) = %d; want 5", len(list))
	}
}

func TestIntersection_fully_B_in_A(t *testing.T) {
	list1 := []int{1, 2, 3}
	list2 := []int{2, 3}

	intersection := intersectionChecker(list1, list2)

	if intersection == true {
		fmt.Println("Intersection([1,2,3], [2,3]) = true")
	} else {
		t.Errorf("Intersection([1,2,3], [2,3]) = %t; want true", intersection)
	}
}

func TestIntersection_fully_A_in_B(t *testing.T) {
	list1 := []int{2, 3}
	list2 := []int{1, 2, 3}

	intersection := intersectionChecker(list1, list2)

	if intersection == true {
		fmt.Println("Intersection([2,3], [1,2,3]) = true")
	} else {
		t.Errorf("Intersection([2,3], [1,2,3]) = %t; want true", intersection)
	}
}

func TestIntersection_partially_A_in_B(t *testing.T) {
	list1 := []int{1, 2, 3}
	list2 := []int{3, 4}

	intersection := intersectionChecker(list1, list2)
	if intersection == false {
		fmt.Println("Intersection([1,2,3], [3,4]) = false")
	} else {
		t.Errorf("Intersection([1,2,3], [3,4]) = %t; want false", intersection)
	}
}

func TestIntersection_partially_B_in_A(t *testing.T) {
	list1 := []int{3, 4}
	list2 := []int{1, 2, 3}

	intersection := intersectionChecker(list1, list2)
	if intersection == false {
		fmt.Println("Intersection([3,4], [1,2,3]) = false")
	} else {
		t.Errorf("Intersection([3,4], [1,2,3]) = %t; want false", intersection)
	}
}

func TestIntersection_no_intersection(t *testing.T) {
	list1 := []int{1, 2}
	list2 := []int{3, 4}

	intersection := intersectionChecker(list1, list2)
	if intersection == false {
		fmt.Println("Intersection([1,2], [3,4]) = false")
	} else {
		t.Errorf("Intersection([1,2], [3,4]) = %t; want false", intersection)
	}
}

func TestIntersection_singleDig_in_B(t *testing.T) {
	list1 := []int{1}
	list2 := []int{1, 2}

	intersection := intersectionChecker(list1, list2)
	if intersection == true {
		fmt.Println("Intersection([1], [1,2]) = true")
	} else {
		t.Errorf("Intersection([1], [1,2]) = %t; want true", intersection)
	}
}

func TestIntersection_singleDig_in_A(t *testing.T) {
	list1 := []int{1, 2}
	list2 := []int{1}

	intersection := intersectionChecker(list1, list2)
	if intersection == true {
		fmt.Println("Intersection([1,2], [1]) = true")
	} else {
		t.Errorf("Intersection([1,2], [1]) = %t; want true", intersection)
	}
}



