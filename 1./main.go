package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		inputFile.Close()
	}()

	scanner := bufio.NewScanner(inputFile)

	var former *int
	var inc int = 0
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		if former == nil {
			former = &val
			continue
		}

		if val > *former {
			inc++
		}
		former = &val
	}
	log.Printf("%d increases\n", inc)
}
