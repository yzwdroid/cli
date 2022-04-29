package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	// Define a boolen flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	flag.Parse()
	// Calling the count function to the number of words
	// received from the Standard Input and printing it out
	fmt.Println("Total words: " + strconv.Itoa(count(os.Stdin, *lines)))
}

func count(r io.Reader, countLines bool) int {
	// Creating a new Scanner to read the input
	scanner := bufio.NewScanner(r)
	// Define the scanner split type to words(default is split by lines)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0
	// For every word or line scanned, add 1 to the counter
	for scanner.Scan() {
		wc++
	}
	return wc
}
