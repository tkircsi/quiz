package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	bShuffle := flag.Bool("s", true, "Shuffle questions and answers")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %q\n", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parsed the provided CSV file.")
	}

	problems := parseLines(lines)

	if *bShuffle {
		shuffle(&problems)
	}

	correct := 0
	for _, p := range problems {
		fmt.Printf("%s = ", p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("You scored %d of %d.\n", correct, len(problems))
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func shuffle(problems *[]problem) {
	rand.Seed(time.Now().UnixNano())
	l := len(*problems) - 1

	for i := 0; i < 10; i++ {
		r := rand.Intn(l)
		(*problems)[r], (*problems)[l-r] = (*problems)[l-r], (*problems)[r]
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
