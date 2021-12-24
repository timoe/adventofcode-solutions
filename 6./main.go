package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fish := parseInput()

	generations := make([]int, 9)

	for _, v := range fish {
		generations[v] += 1
	}

	for i := 0; i < 256; i++ {
		var swapping = generations[8]
		thoseFishAtZeroCycle := generations[0]
		generations[8] = thoseFishAtZeroCycle

		for j := 7; j >= 0; j-- {
			current := generations[j]
			generations[j] = swapping
			swapping = current
		}
		generations[6] += thoseFishAtZeroCycle
	}

	var sum = 0
	for _, v := range generations {
		sum += v
	}
	duration := time.Now().Sub(start)
	log.Printf("Fish: %d, took %dms (%d ns)\n", sum, duration.Milliseconds(), duration.Nanoseconds())

}

// let's assume no errors
func parseInput() []int {
	content, _ := ioutil.ReadFile("input.text")
	stringFish := strings.Split(string(content), ",")
	intFish := make([]int, len(stringFish))

	for i, v := range stringFish {
		n, _ := strconv.Atoi(v)
		intFish[i] = n
	}

	return intFish
}
