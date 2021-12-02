package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Returns window size from CLI args.
func getWindowSize(args []string) int {
	switch len(args) {
	case 1, 2:
		return 1
	default:
		var result, _ = strconv.Atoi(os.Args[2])
		return result
	}
}

// Returns a sum of a slice of provided window size in a given array
func getSumOfNumWindow(numArr []int) int {
	var windowSize = getWindowSize(os.Args)

	if windowSize > len(numArr) {
		log.Fatal("Window size cannot be larger than provided array.")
	}

	sum := 0

	for _, i := range numArr[:windowSize] {
		sum = sum + i
	}

	return sum
}

// Given two CLI arguments:
// - path to a text file consisting of N numbers per line
// - an integer of sliding window size.
//
// Counts the number of times the sum of measurements in this sliding window increases from the previous sum.
//
// Example window of a size 3:
//   199  A
//   200  A B
//   208  A B C
//   210    B C D
//   200  E   C D
//   207  E F   D
//   240  E F G
//   269    F G H
//   260      G H
//   263        H
//
//   The measurements in the first window are marked A (199, 200, 208);
//   their sum is 199 + 200 + 208 = 607. The second window is marked
//   B (200, 208, 210); its sum is 618. The sum of measurements in the second
//   window is larger than the sum of the first, so this first comparison increased.
// Example usage:
//   go run input-day-1.txt 3
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
	var previousNum = getSumOfNumWindow(numArray)

	// Set counter to 0.
	var increasedCount = 0

	// Increment counter and set new previous number.
	for index := range numArray {
		var sum = 0
		var size = getWindowSize(os.Args)
		var sliceEndOffset = index + size - 1

		// Prevent from slicing if there are not enough items at the end.
		if sliceEndOffset < len(numArray) {
			var slice = numArray[index : index+size]

			sum = getSumOfNumWindow(slice)
		}

		if sum > previousNum {
			increasedCount++
		}

		// Set new starting value to be compared against.
		previousNum = sum
	}

	// Print result.
	fmt.Println("Number of times num sequence increased:", increasedCount)
}
