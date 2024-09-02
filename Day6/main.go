package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func manageWhitespace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

type Pair struct {
	time     int
	distance int
}

func preprocess(t string,dist string) []Pair {
	times := strings.Fields(t)
		//fmt.Print(times)

	dists := strings.Fields(dist)
	//fmt.Print("\n",dists)

	var races []Pair

	for i := 0; i < len(times); i++ {
		ti, err := strconv.Atoi(times[i])
		if err != nil {
			log.Fatal("Error converting string to int")
		}
		di, err := strconv.Atoi(dists[i])
		if err != nil {
			log.Fatal("Error converting string to int")
		}
		races = append(races, Pair{ti, di})
	}

	return races

}
func part1(races []Pair) int {
	var pos []int

	for _, race := range races {
		start := 1
		end := race.time - 1

		for (race.time-start)*start <= race.distance && (race.time-end)*end <= race.distance {
			start += 1
			end -= 1
		}

		pos = append(pos, end-start+1)

	}

	prod := 1

	for _, ele := range pos {
		prod *= ele
	}

	return prod
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Failed to open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var t string
	var dist string
	for scanner.Scan() {
		line := scanner.Text()
		line = manageWhitespace(line)

		if strings.Contains(line, "Time") {
			t = strings.Split(line, ":")[1]

		}

		if strings.Contains(line, "Distance") {
			dist = strings.Split(line, ":")[1]

		}
	}

	if len(os.Args)!=2 {
		log.Fatal("Incorrect number of command line args,probably pasasing none. Pass 1 for part1 or 2 for part2")
	}

	arg := os.Args[1]

	if arg == "1" {
		races:=preprocess(t,dist)

		ans1 := part1(races)
		fmt.Printf("Answer to part 1 is %v\n", ans1)

	} else {

		t = strings.ReplaceAll(t, " ", "")
		dist = strings.ReplaceAll(dist, " ", "")

		races:=preprocess(t,dist)

		ans2:=part1(races)
		fmt.Printf("Answer to part2 is %v\n",ans2)

	}
}
