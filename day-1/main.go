package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Initialize empty array to store numbers for each line.
	var numArray []int

	// Open file based on CLI args.
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	// Init scanner.
	scanner := bufio.NewScanner(file)

	// Split file into lines.
	scanner.Split(bufio.ScanLines)

	// Push each line to numArray.
	for scanner.Scan() {
		var text = scanner.Text()
		var num, _ = strconv.Atoi(text)

		numArray = append(numArray, num)
	}

	// Close the file.
	file.Close()

	// There is no measurement before the first measurement.
	var previousNum = numArray[0]

	// Set counter to 0.
	var increasedCount = 0

	// Increment counter and set new previous number.
	for _, i := range numArray {
		if i > previousNum {
			increasedCount++
		}
		previousNum = i

	}

	// Print result.
	fmt.Println("Number of times num sequence increased:", increasedCount)
}
