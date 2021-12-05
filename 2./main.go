package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

const forward = "forward"
const up = "up"
const down = "down"

func main() {
	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		inputFile.Close()
	}()

	scanner := bufio.NewScanner(inputFile)

	r, err := regexp.Compile("^([a-z]+)\\s([0-9]+)$")
	if err != nil {
		log.Fatal(err)
	}

	var h, d = 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		matches := r.FindSubmatch([]byte(line))

		amount, _ := strconv.Atoi(string(matches[2]))
		switch string(matches[1]) {
		case forward:
			h += amount
		case down:
			d += amount
		case up:
			d -= amount
		}
	}
	log.Printf("h: %d, d: %d", h, d)
}
