package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var map1=map[string] int {
	"red" :12,
	"green":13,
	"blue":14,
}

func part1(gamelog string) bool {
	data:=strings.SplitAfter(gamelog, ";")
	//data=strings.Split(data, ",")
	
	//fmt.Print(data,"\n")

	for i:=0;i<len(data);i++ {
		turnlog:=data[i]
		turnlog=strings.TrimSuffix(turnlog,";")
		turnlog=strings.TrimPrefix(turnlog," ")

		//fmt.Print(turnlog,"\n")
		records:=strings.Split(turnlog,",")
		
		//fmt.Print(records,"\n")
		for j:=0;j<len(records);j++ {
			records[j]=strings.TrimPrefix(records[j]," ")
			pair:=strings.Split(records[j], " ")
			//fmt.Print(pair,"\n")
		
			count:=pair[0]
			clr:=pair[1]

			cnt,err:=strconv.Atoi(count)
			if err!=nil {
				log.Fatal(err)
			}

			if map1[clr]<cnt {
				return false
			}
		}

		
		
	}

	
	return true


}

func part2(gamelog string) int{
	data:=strings.SplitAfter(gamelog, ";")
	//data=strings.Split(data, ",")
	
	//fmt.Print(data,"\n")

	minBlue:=0
	minRed:=0
	minGreen:=0


	for i:=0;i<len(data);i++ {
		turnlog:=data[i]
		turnlog=strings.TrimSuffix(turnlog,";")
		turnlog=strings.TrimPrefix(turnlog," ")

		//fmt.Print(turnlog,"\n")
		records:=strings.Split(turnlog,",")
		
		//fmt.Print(records,"\n")
		for j:=0;j<len(records);j++ {
			records[j]=strings.TrimPrefix(records[j]," ")
			pair:=strings.Split(records[j], " ")
			//fmt.Print(pair,"\n")
		
			count:=pair[0]
			clr:=pair[1]

			cnt,err:=strconv.Atoi(count)
			if err!=nil {
				log.Fatal(err)
			}

			if clr=="red" {
				minRed=max(minRed,cnt)
			}

			if clr=="green" {
				minGreen=max(minGreen,cnt)
			}

			if clr=="blue" {
				minBlue=max(minBlue,cnt)
			}
		
		}

		
		
	}

	return minRed*minGreen*minBlue

	
}

func main(){
	file,err:=os.Open("./cubes.txt")

	if err!=nil {
		fmt.Print(err)
		return 
	}

	defer file.Close()

	scanner:=bufio.NewScanner(file)

	power:=0
	slice1 := []int{}
	for scanner.Scan() {
		s:=scanner.Text()
		//fmt.Print(s+"\n")
		gamelog:=strings.Split(s,":")[1]
		gameTag:=strings.Split(s,":")[0]
		gameID:=strings.Split(gameTag," ")[1]
		gID,err:=strconv.Atoi(gameID)

		if err!=nil {
			log.Fatal(err)
		}

		canAdd:=part1(gamelog)

		power+=part2(gamelog)


		if canAdd {
			slice1=append(slice1, gID)
		}

		
	}

	


	part1ans:=0

	for i:=0;i<len(slice1);i++ {
		part1ans+=slice1[i]
	}

	fmt.Printf("Answer to part1 is %d\n",part1ans)

	fmt.Printf("Answer to part2 is %d\n",power)


}