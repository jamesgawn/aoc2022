package main

import (
	"aoc2022/dayx"
	"aoc2022/pkg/utils"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	defer utils.TimeTrack(time.Now(), "Total execution")
	day := "4"
	if len(os.Args) > 1 {
		day = string(os.Args[1])
	}
	inputFileLocation := "./day" + day + "/input.txt"
	if len(os.Args) > 2 {
		inputFileLocation = string(os.Args[1])
	}

	log.Println("Using input file: " + inputFileLocation)
	file, err := os.Open(inputFileLocation)
	if err != nil {
		log.Println("I can't solve every problem... not yet anyway")
		log.Fatal(err)
	}

	FindSolution(day, file)

	cErr := file.Close()
	if cErr != nil {
		log.Fatal(cErr)
	}
}

func FindSolution(day string, input io.Reader) {
	log.Println("Day " + day)
	switch day {
	case "x":
		dayx.ExecuteSolution(input)
	default:
		log.Fatal("I can't solve every problem... not yet anyway")
	}
}
