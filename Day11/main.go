package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)


type Pair struct {
	x int 
	y int
}

func emptyRows(grid []string) []int{
	var emptyrows []int
	for i:=0;i<len(grid);i++ {
		flag:=true
		for j:=0;j<len(grid[i]);j++ {
			if grid[i][j]!='.' {
				flag=false 
				break
			}
			
		}
		if flag {
			emptyrows = append(emptyrows, i)
		}
	}
	return emptyrows
}

func emptyCols(grid []string) []int {
	var emptycols []int
	for j:=0;j<len(grid[0]);j++ {
		flag:=true
		for i:=0;i<len(grid);i++ {
			if grid[i][j]!='.' {
				flag=false
				break
			}	
		}
		if flag {
			emptycols = append(emptycols, j)
		}
	}
	return emptycols
}

func findGalaxies(grid []string) []Pair{
	var galx []Pair 

	for i:=0;i<len(grid);i++ {
		for j:=0;j<len(grid[0]);j++ {
			if grid[i][j]=='#' {
				galx = append(galx, Pair{i,j})
			}
		}
	}

	return galx

}


func min (a int, b int) int{
	if a<b {
		return a
	}
	return b
}

func max(a int,b int) int {
	if a>b {
		return a
	} 
	return b
}
func part1(grid []string,scale int) int{
	// for _,ele:=range grid {
	// 	fmt.Println(ele)
	// }	

	emptyrows:=emptyRows(grid)
	emptycols:=emptyCols(grid)
	

	//fmt.Println(emptyrows,"\n",emptycols)

	galaxies:=findGalaxies(grid)

//	fmt.Println(galaxies)


	path:=0
	for i:=0;i<len(galaxies);i++ {
		for j:=i+1;j<len(galaxies);j++ {
			start:=galaxies[i]
			end:=galaxies[j]

			//oldpath:=path

			for m:=min(start.x,end.x); m<max(start.x,end.x); m++ {
				if slices.Contains(emptyrows,m) {
					path+=scale
				} else {
					path+=1
				}
			}

			for n:=min(start.y,end.y); n<max(start.y,end.y); n++ {
				if slices.Contains(emptycols,n) {
					path+=scale
				} else {
					path+=1
				}
			}

			//fmt.Printf("Ind path between %v and %v is %v\n",i+1,j+1,path-oldpath)
		}
	}

	return path

}
func main() {
	file,err:=os.Open("./input.txt")
	if err!=nil {
		log.Fatal("Error opening file")
	}

	defer file.Close()

	scanner:=bufio.NewScanner(file)

	var grid []string

	for scanner.Scan() {
		line:=scanner.Text()
		grid = append(grid, line)
	}

	ans1:=part1(grid,2)

	fmt.Printf("Answer to part1 is %v\n",ans1)

	ans2:=part1(grid,1000000)

	fmt.Printf("Answer to part2 is %v\n",ans2)
}