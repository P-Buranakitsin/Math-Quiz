package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	filePtr, timePtr, countPtr := comFlag()
	indexVal, numLine := readFile(filePtr)
	randLine := genRandLine(countPtr, numLine)

	var correctAns int // Number of correct answer
	c1 := make(chan string)

	var start time.Time
	var duration time.Duration
	// Start another go routine
	go func() {
		start = time.Now() // Measure execution time
		// Randomly select lines from a file by using map
		for k := range randLine {
			split := strings.Split(indexVal[k], ",") // split into slice of strings, index 0 is question while index 1 is answer
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(split[0], " = ")

			text, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}

			// String returned from reader.ReadString includes Carriage Return (ASCII 13) and Line Feed (ASCII 10)
			// Slicing is used to remove Carriage Return and Line Feed

			// **** When run in integrated VS code terminal, Carriage Return is included in the returned value from reader.ReadString whereas
			// in normal terminal Carriage Return is not included

			// Keep removing character from text variable until its length is equal to split[1] length
			for len(text) != len(split[1]) {
				text = text[:len(text)-1]
			}

			if text == split[1] {
				correctAns++
			}
		}
		duration = time.Since(start)
		c1 <- "Done"
	}()

	// Select blocks until either case is satisfied
	var timeUsed float64
	select {
	case <-c1:
		timeUsed = duration.Seconds()
	case <-time.After(time.Duration(*timePtr) * time.Second):
		timeUsed = float64(*timePtr)
		fmt.Println("\nTime's up!!!")
	}
	fmt.Println("\ncorrect =", correctAns)
	fmt.Println("time used =", timeUsed, "seconds")
}

// Command-Line Flags
func comFlag() (*string, *int, *int) {
	//file flag
	filePtr := flag.String("file", "./questions.csv", "file path to read from")
	//timeout flag
	timePtr := flag.Int("timeout", 180, "duration in seconds to answer all questions")
	//count flag
	countPtr := flag.Int("count", 5, "number of questions")

	flag.Parse()
	return filePtr, timePtr, countPtr
}

// Read from a file
func readFile(filePtr *string) (map[int]string, int) {

	f, err := os.Open(*filePtr)
	var numLine int
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)

	var indexVal = map[int]string{} // Create map to store string read in each line as well as its no. of line
	for s.Scan() {
		indexVal[numLine] = s.Text()
		numLine++ // numLine count no. of lines from a file
	}
	err = s.Err() // Check an error while scanning
	if err != nil {
		log.Fatal(err)
	}

	return indexVal, numLine
}

// Generate unique random numbers of lines to read from
func genRandLine(countPtr *int, numLine int) map[int]bool {
	randLine := map[int]bool{}

	if *countPtr > numLine {
		*countPtr = numLine
	}

	i := 0
	for i != *countPtr {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(numLine) // 10 is number of lines in a file
		if !randLine[randNum] {
			randLine[randNum] = true
			i++
		}
	}

	return randLine
}
