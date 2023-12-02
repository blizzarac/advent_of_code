package main


import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

func calculatePriority(item1 rune) int {
	result := 0
	if unicode.IsLower(item1) {
		log.Println("Found: ", string(item1), "priority: ", int(item1)-96)
		result += int(item1) - 96
	} else {
		log.Println("Found: ", string(item1), "priority: ", int(item1)-38)
		result += int(item1) - 38
	}

	return result
}

func removeDuplicate[T string | int | rune](sliceList []T) []T {
    allKeys := make(map[T]bool)
    list := []T{}
    for _, item := range sliceList {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}

func main() {

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
	numLines := len(lines)

	result := 0

	for i := 0; i < numLines-2; i=i+3 {
		line1 := lines[i]
		line2 := lines[i+1]
		line3 := lines[i+2]

		itemList1 := []rune(line1)
		itemList2 := []rune(line2)
		itemList3 := []rune(line3)

		uniqueList1 := removeDuplicate(itemList1)
		uniqueList2 := removeDuplicate(itemList2)
		uniqueList3 := removeDuplicate(itemList3)

		log.Println("Comparing: ", string(uniqueList1), "with", string(uniqueList2), "with", string(uniqueList3))

		for _, item1 := range uniqueList1 {
			for _, item2 := range uniqueList2 {
				for _, item3 := range uniqueList3 {
					if item1 == item2 && item2 == item3 {
						result += calculatePriority(item1)
						break;
					}
				}
			}
		}

		log.Println("Result: ", result)
	}

}
