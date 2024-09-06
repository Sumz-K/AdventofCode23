package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func allzero(arr[]int) bool {
	for i:=0;i<len(arr);i++ {
		if arr[i]!=0 {
			return false
		}
	}
	return true
}

func part1(arr []int) int{
	var temp []int
	// fmt.Print(arr,"\n")
	// fmt.Print(len(arr),"\n")
	for i:=1;i<len(arr);i++ {
		temp = append(temp, arr[i]-arr[i-1])
	}

	
	if allzero(temp) {
		return arr[len(temp)-1]
	} else {
		return arr[len(arr)-1]+part1(temp)
	}

}

func part2(arr[]int) int{
	var temp []int

	for i:=1;i<len(arr);i++ {
		temp = append(temp, arr[i]-arr[i-1])
	}

	if allzero(temp) {
		return arr[0]
	} else {
		return arr[0]-part2(temp)
	}

}

func main() {
	file, err := os.Open("./input.txt")

	if err!=nil {
		log.Fatal("Cant open file\n")
	}

	scanner:=bufio.NewScanner(file)

	ans1:=0
	ans2:=0
	for scanner.Scan() {
		line:=scanner.Text()

		lineSlice:=strings.Fields(line)

		//fmt.Print(len(lineSlice),"\n")
		intSlice:=make([]int,len(lineSlice))
		for i,ele:=range lineSlice {
			intSlice[i],_=strconv.Atoi(ele)
		}

		//fmt.Println(len(intSlice))
		ans1+=part1(intSlice)

		ans2+=part2(intSlice)
		
	}

	fmt.Println(ans1)

	fmt.Println(ans2)

}