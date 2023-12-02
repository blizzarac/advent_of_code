package main

import (
	"log"
	"io/ioutil"
	"os"

)


func findMarkerPosition(input []byte, searchWindow int) int {
	var windowBuffer map[string]bool

	for i := 0; i < len(input); i++ {
//		log.Printf("Current window: %s", curWindow)
		windowBuffer = map[string]bool{}
		for j := 0; j < searchWindow; j++ {
			windowBuffer[string(input[i+j])] = true
		}

		if len(windowBuffer) == searchWindow {
			return i + searchWindow
		}

	}

	return 0
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

	result := findMarkerPosition(content, 14)

	log.Printf("Result: %d", result)

}
