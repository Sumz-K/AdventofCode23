package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"reflect"
	"sort"
	"strconv"
	"strings"
)

var hands = map[string]int{
	"FIVE_KIND":  7,
	"FOUR_KIND":  6,
	"FULL_HOUSE": 5,
	"THREE_KIND": 4,
	"TWO_PAIR":   3,
	"ONE_PAIR":   2,
	"HIGH_CARD":  1,
}

type Hand struct {
	cards    string
	bid      int
	handType int
}

func identifyHand(cards string, part2 bool) int {
	cardMap := map[rune]int{}
	

	//fmt.Print(values,"\n")
	var htype int
	if !part2 {
		for _, c := range cards {
			cardMap[c] += 1
		}
	
		values := make([]int, 0, len(cardMap))
	
		for _, v := range cardMap {
			values = append(values, v)
		}
	
		sort.Sort(sort.Reverse(sort.IntSlice(values)))

		highest := values[0]

		second := 0
		if len(values) > 1 {
			second = values[1]
		}

		switch highest {
		case 5:
			htype = hands["FIVE_KIND"]
		case 4:
			htype = hands["FOUR_KIND"]
		case 3:
			if second == 2 {
				htype = hands["FULL_HOUSE"]
			} else if second == 1 {
				htype = hands["THREE_KIND"]
			}
		case 2:
			if second == 2 {
				htype = hands["TWO_PAIR"]
			} else {
				htype = hands["ONE_PAIR"]
			}
		case 1:
			htype = hands["HIGH_CARD"]
		default:
			htype = -1
			fmt.Print("Invalid hand\n")

		}
	} else {
		jokercount:=0
		for _,c:=range cards {
			if c=='J'{
				jokercount+=1
				continue
			}
			cardMap[c]+=1
		}

		values := make([]int, 0, len(cardMap))
	
		for _, v := range cardMap {
			values = append(values, v)
		}
	
		sort.Sort(sort.Reverse(sort.IntSlice(values)))
		

		highest:=jokercount
		if len(values) > 0  { 
			highest = values[0] + jokercount
		}

		second := 0
		if len(values) > 1 {
			second = values[1]
		}

		switch highest {
		case 5:
			htype = hands["FIVE_KIND"]
		case 4:
			htype = hands["FOUR_KIND"]
		case 3:
			if second == 2 {
				htype = hands["FULL_HOUSE"]
			} else if second == 1 {
				htype = hands["THREE_KIND"]
			}
		case 2:
			if second == 2 {
				htype = hands["TWO_PAIR"]
			} else {
				htype = hands["ONE_PAIR"]
			}
		case 1:
			htype = hands["HIGH_CARD"]
		default:
			htype = -1
			fmt.Print("Invalid hand\n")

		}

	}

	return htype

}

func compareCards(a string, b string,part2 bool) (int, int) {
	order := "AKQJT98765432"

	if part2 {
		order="AKQT98765432J"
	}

	first := 0
	second := 0

	for idx, c := range order {
		if string(c) == a {
			first = idx
		}
		if string(c) == b {
			second = idx
		}
	}
	return first, second
}
func compareHands(h1 Hand, h2 Hand,part2 bool) bool {
	if h1.handType != h2.handType {
		return h1.handType < h2.handType
	}
	card1 := h1.cards
	card2 := h2.cards

	for i := 0; i < len(card1); i++ {
		if card1[i] == card2[i] {
			continue
		} else {
			c1 := string(card1[i])
			c2 := string(card2[i])
			f, s := compareCards(c1, c2,part2)

			if f < s {
				return false
			}
			return true
		}
	}
	return true
}

func part1(hands []Hand) int {
	count := 0
	for i, hand := range hands {
		count += (i + 1) * hand.bid
	}
	return count
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	hands := []Hand{}

	if len(os.Args) != 2 {
		log.Fatal("Incorrect number of command line args,probably pasasing none. Pass 1 for part1 or 2 for part2")
	}

	arg := os.Args[1]

	if arg == "1" {

		for scanner.Scan() {
			line := scanner.Text()
			cards := strings.Split(line, " ")[0]
			bid := strings.Split(line, " ")[1]

			bidInt, err := strconv.Atoi(bid)
			if err != nil {
				log.Fatal("Error converting string to int")
			}

			//fmt.Printf("Cards: %v Bid: %v\n",cards,bidInt)

			handVal := identifyHand(cards, false)

			//fmt.Printf("Handval %v\n",handVal)

			hands = append(hands, Hand{cards, bidInt, handVal})

		}

		sort.Slice(hands, func(i, j int) bool {
			return compareHands(hands[i], hands[j],false)
		})

		ans1 := part1(hands)

		fmt.Printf("Answer to part 1 is %v\n", ans1)
	} else {
		for scanner.Scan() {
			line := scanner.Text()
			cards := strings.Split(line, " ")[0]
			bid := strings.Split(line, " ")[1]

			bidInt, err := strconv.Atoi(bid)
			if err != nil {
				log.Fatal("Error converting string to int")
			}

			//fmt.Printf("Cards: %v Bid: %v\n",cards,bidInt)

			handVal := identifyHand(cards, true)

			//fmt.Printf("Handval %v\n",handVal)

			hands = append(hands, Hand{cards, bidInt, handVal})

		}

		sort.Slice(hands, func(i, j int) bool {
			return compareHands(hands[i], hands[j],true)
		})

		ans1 := part1(hands)

		fmt.Printf("Answer to part 2 is %v\n", ans1)
	}
}
