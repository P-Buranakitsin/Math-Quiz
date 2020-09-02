package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)


func main() {
	filePtr, _, countPtr := comFlag()
	countAndRead(filePtr, countPtr)

}

// Command-Line Flags
func comFlag() (*string, *int, *int) {
	//file flag
	filePtr := flag.String("file", "./questions.csv", "file path to read from")
	//timeout flag
	timePtr := flag.Int("timeout", 180, "duration in seconds for each question")
	//count flag
	countPtr := flag.Int("count", 5, "number of questions")

	flag.Parse()
	return filePtr, timePtr, countPtr
}

// Generate unique random numbers of lines to read from
func genRandLine(countPtr *int, numLine int) map[int]bool {
	randLine := map[int]bool{}

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

// Count lines from a file and read random lines from a file
func countAndRead (filePtr *string, countPtr *int) {
	// Count lines from a file
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
	for s.Scan() {
		numLine++
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	// After counting lines, select random lines from a file
	randLine := genRandLine(countPtr, numLine)
	fmt.Println(randLine)
	f.Seek(0,0)
	s = bufio.NewScanner(f)// Reset Scanner 
	var line int
	for s.Scan() {
		if (randLine[line]) {
			fmt.Println(s.Text())
		}
		line++
	}
}