package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const maxWindows = 3

func main() {
	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		inputFile.Close()
	}()

	scanner := bufio.NewScanner(inputFile)

	lines := make([]int, 0)
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, val)
	}

	var inc = 0
	for i := 0; i+4 <= len(lines); i++ {
		var sumA = 0
		var sumB = 0
		for j := i; j < i+3; j++ {
			sumA += lines[j]
		}
		for k := i + 1; k < i+4; k++ {
			sumB += lines[k]
		}

		if sumB > sumA {
			inc++
		}
		sumA = 0
		sumB = 0
	}
	log.Printf("%d increases", inc)
}
