package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	//"strconv"
	"strings"
)

func manageWhitespace(s string) string {
	return strings.Join(strings.Fields(s)," ")
}

func getCommon(s string) int {
	cards:=strings.Split(s, ":")[1]
	

	//fmt.Print(cards,"\n")

	winning:=strings.Split(cards, "|")[0]
	have:=strings.Split(cards, "|")[1]

	haveSlice:=strings.Fields(have)
	winSlice:=strings.Fields(winning)
	
	mp:=make(map [string]int)

	count:=0
	for _,c:=range winSlice {
		mp[c]=1
	}
	
	for _,d:=range haveSlice {
		if mp[d] ==1 {
			count+=1
		}
	}

	return count
}

func part1(s string) int {
	s=manageWhitespace(s)
	count:=getCommon(s)

	return int(math.Pow(2,float64(count-1)))
}

var cardMap=make(map[int]int)

func part2(s string){
	s=manageWhitespace(s)
	cardID:=strings.Split(strings.Split(s, ":")[0]," ")[1]
	cardIDint,err:=strconv.Atoi(cardID)

	if err!=nil {
		log.Fatal("Couldny convert string to int")
	}

	

	cardMap[cardIDint]+=1

	common:=getCommon(s)

	for i:=cardIDint+1;i<=cardIDint+common;i++ {
		cardMap[i]+=cardMap[cardIDint]
	}


}

func main() {
	file, err := os.Open("./cards.txt")

	if err!=nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner:=bufio.NewScanner(file)

	part1ans:=0
	for scanner.Scan() {
		s:=scanner.Text()

		part1ans+=part1(s)

		part2(s)
	}

	fmt.Print("Answer to part1 is ",part1ans,"\n")

	///fmt.Print(cardMap)

	part2ans:=0

	for _,v:= range cardMap {
		part2ans+=v
	}

	fmt.Print("Answer to part1 is ",part2ans,"\n")





}