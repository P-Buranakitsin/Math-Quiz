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
	
	randLine := genRandLine(countPtr)
	fmt.Println(randLine)

	f, err := os.Open(*filePtr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var cc int
	s := bufio.NewScanner(f)
	for s.Scan() {
		if (randLine[cc]) {
			fmt.Println(s.Text())
		}
		cc++
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	
	fmt.Println(cc)

	/*data, err := ioutil.ReadFile(*filePtr)
	if err != nil {
		log.Fatal("File reading error ", err)
	}
	fmt.Println(string(data))*/
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
func genRandLine(count *int) map[int]bool {
	randLine := map[int]bool{}

	i := 0
	for i != *count {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(10) // 10 is number of lines in a file
		if !randLine[randNum] {
			randLine[randNum] = true
			i++
		}
	}

	return randLine

}
