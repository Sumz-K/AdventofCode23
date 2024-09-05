package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func manageWhitespace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func gcd(a int,b int) int {
	if b==0{
		return a
	}
	return gcd(b,a%b)
}

func lcm(a,b int) int {
	return (a*b)/gcd(a,b)
}

func part1(dirs string,mp map[string][]string) int{
	// fmt.Println(dirs," ",len(dirs))

	// fmt.Println(mp)

	curr:="AAA"

	steps:=0

	idx:=0

	for curr!="ZZZ" {
		dir:=dirs[idx]
		if dir=='L' {
			curr=mp[curr][0]
		} else if dir=='R' {
			curr=mp[curr][1]
		}
		steps+=1
		idx=(idx+1)%len(dirs)
		
	}
	return steps

}

func part2(dirs string,mp map[string][]string) int {
	startNodes:=[]string{}

	endNodes:=[]string{}
	for key,_:=range mp {
		if strings.HasSuffix(key,"A") {
			startNodes = append(startNodes, key)
		}

		if strings.HasSuffix(key,"Z") {
			endNodes = append(endNodes, key)
		}
	}

	//fmt.Print(startNodes,"\n",endNodes)

	res:=[]int{}

	idx:=0

	

	for _,node:=range startNodes {
		steps:=0
		curr:=node
		for slices.Contains(endNodes,curr)==false {
			if dirs[idx]=='L' {
				curr=mp[curr][0]
			} else {
				curr=mp[curr][1]
			}
			steps+=1
			idx=(idx+1)%len(dirs)
		}

		res = append(res, steps)
	}
	
	l:=1

	for _,ele:=range res {
		l=lcm(l,ele)
	}
	return l
}

func main() {

	file, err := os.Open("./input.txt")
	if err!=nil {
		log.Fatal("Error opening file\n")
	}

	defer file.Close()

	scanner:=bufio.NewScanner(file)

	var dirs string

	mp:=map[string][]string{}

	for scanner.Scan() {
		line:=scanner.Text()

		if !strings.Contains(line,"=") && line!="" {
			dirs=line
		} else if line!=""{
			line=manageWhitespace(line)
			//fmt.Println(line)

			sep:=strings.Split(line,"=")
			//fmt.Print(sep)

			key:=sep[0]
			vals:=sep[1]

			key=strings.ReplaceAll(key," ","")
			vals=strings.ReplaceAll(vals," ","")

			vals=strings.ReplaceAll(vals,"(","")
			vals=strings.ReplaceAll(vals,")","")

			left:=strings.Split(vals, ",")[0]
			right:=strings.Split(vals, ",")[1]


			
			tempSlice:=make([]string,2)

			tempSlice[0]=left
			tempSlice[1]=right

			mp[key]=tempSlice
			
		}
	}

	
	// ans1:=part1(dirs,mp)

	// fmt.Printf("Answer to part1 is %v\n",ans1)

	ans2:=part2(dirs,mp)


	fmt.Printf("Answer to part2 is %v\n",ans2)

}
