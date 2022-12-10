package day1

import (
	"bufio"
	"io"
	"log"
	"sort"
	"strconv"
)

type elf struct {
	food []int
}

func ExecuteSolution(input io.Reader) {

	elves := make([]elf, 0)
	currentElf := elf{food: make([]int, 0)}
	scanner := bufio.NewScanner(input)
	text := ""
	for scanner.Scan() {
		text = scanner.Text()
		if len(text) == 0 {
			elves = append(elves, currentElf)
			currentElf = elf{food: make([]int, 0)}
		} else {
			currentElf.food = append(currentElf.food, ExtractFoodCalories(text))
		}
	}
	currentElf.food = append(currentElf.food, ExtractFoodCalories(text))
	elves = append(elves, currentElf)

	sort.Slice(elves, func(i, j int) bool {
		return SumCal(elves[i]) > SumCal(elves[j])
	})

	fatElfCals := SumCal(elves[0])
	log.Printf("Calories on elf with most food: %d", fatElfCals)

	fatestElfCals := SumCal(elves[0]) + SumCal(elves[1]) + SumCal(elves[2])
	log.Printf("Calories on top 3 elves with most food: %d", fatestElfCals)
}

func ExtractFoodCalories(text string) int {
	foodcal, err := strconv.Atoi(text)
	if err != nil {
		log.Fatalln("Unexpected non-integer")
	}
	return foodcal
}

func FindMostCal(elves []elf) int {
	maxCals := 0
	for _, elf := range elves {
		cals := SumCal(elf)
		if cals > maxCals {
			maxCals = cals
		}
	}
	return maxCals
}

func SumCal(elf elf) int {
	cals := 0
	for _, cal := range elf.food {
		cals += cal
	}
	return cals
}
