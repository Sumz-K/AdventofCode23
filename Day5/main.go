package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func manageWhitespace(s string) string {
	return strings.Join(strings.Fields(s)," ")
}


func part1(seeds []int,maps [][][]int) int{

	var locs []int
	for _,seed:=range seeds {
		//fmt.Printf("\nUsing seed %v\n",seed)
		for _,temmap:=range maps {
			for _,ele:=range temmap {
				//fmt.Printf("Seed %v %v\n",seed,ele)
				dest:=ele[0]
				src:=ele[1]
				rng:=ele[2]

				if seed>=src && seed<=src+rng-1 {
					seed=dest+seed-src
					//fmt.Printf("Newval is %v\n",seed)
					break
				} 
				
			}
			//fmt.Printf("At end seed is %v\n",seed)
			

		}
		locs = append(locs, seed)

		
	}

	//fmt.Printf("\nFinal locations are %v\n",locs)

	return slices.Min(locs)
}

func part2(seeds []Pair,maps [][][]int) {
	fmt.Println(seeds)

	// fmt.Print("Maps:\n")
	// for _,ele:=range maps {
	// 	fmt.Println(ele)
	// }

	var locs []Pair 
	flag:=0
	for len(seeds) >0 {
		seed:=seeds[0]
		seeds=seeds[1:]
		s:=seed.start
		e:=seed.end
		for _,tempmap:=range maps {
			for _,ele:=range tempmap {
				a,b,c:=ele[0],ele[1],ele[2]
				os:=max(b,s)
				oe:=min(b+c,e)

				if os<oe {
					flag=1
					locs = append(locs, Pair{os-b+a,oe-b+a})

					if e > oe {
						seeds = append(seeds, Pair{oe,e})
					}

					if s>os {
						seeds = append(seeds, Pair{s,os})
					}
					break
				}
			}
		}
		if flag==0 {
			locs = append(locs, Pair{s,e})
		}
	}

	print(locs)
	

}

type Pair struct {
	start int
	end int
}

func main() {
	file, err := os.Open("./input.txt")

	if err!=nil {
		log.Fatal("Error opening input file")
	}

	defer file.Close()

	scanner:=bufio.NewScanner(file)

	var seeds[] int
	var maps[][][]int

	var tempmap[][]int

	for scanner.Scan() {
		line:=scanner.Text()
		line=manageWhitespace(line)
		if strings.Contains(line,"seeds") {
			seedData:=strings.Fields(strings.Split(line, ":")[1])
			for _,s:=range seedData {
				seed,err:=strconv.Atoi(string(s))
				if err!=nil {
					log.Fatal("Error converting string to int")
				}

				seeds=append(seeds, seed)
			}
			
		} else if strings.Contains(line,"map") {
			if tempmap!=nil {
				maps=append(maps, tempmap)
			}
			tempmap=[][]int {}

		} else if line!="" {
			row:=strings.Fields(line)
			var maprow[] int 

			for _,r:=range row {
				val,err:=strconv.Atoi(r) 
				if err!=nil {
					log.Fatal("Error converting string to int")
				}
				maprow=append(maprow, val)
			}
			tempmap=append(tempmap, maprow)
		}
	}

	if tempmap!=nil {
		maps=append(maps, tempmap)
	}

	// fmt.Println("Seeds: ",seeds)

	


	ans1:=part1(seeds,maps)

	fmt.Printf("Answer to part1 is %v\n",ans1)


	var rangeseeds []Pair

	for i:=0;i<len(seeds);i+=2 {
		rangeseeds = append(rangeseeds, Pair{seeds[i],seeds[i]+seeds[i+1]-1})
	}

	

	

	part2(rangeseeds,maps)


}