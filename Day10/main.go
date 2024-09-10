package main

import (
	"bufio"
	"fmt"
	//"fmt"
	"log"
	"os"
	"slices"
	"strings"
	//"strings"
)

type Pair struct {
	row int 
	col int
}


func part1(matrix []string, sr int,sc int) int{
	var queue []Pair
	//fmt.Print(string(matrix[sr][sc]))

	queue = append(queue, Pair{sr,sc})

	var seen []Pair

	seen = append(seen, Pair{sr,sc})

	for len(queue)>0 {
		row:=queue[0].row
		col:=queue[0].col

		queue=queue[1:]

		curr:=string(matrix[row][col])

		//fmt.Print(slices.Contains(seen,Pair{-2,0}))

		//fmt.Print(row," ",col,"\n")

		if row>0 && strings.ContainsAny(curr,"S|JL") && strings.ContainsAny(string(matrix[row-1][col]),"|7F") && !slices.Contains(seen,Pair{row-1,col}){
			//fmt.Print("Here1\n")
			queue = append(queue, Pair{row-1,col})
			seen = append(seen, Pair{row-1,col})
		}

		if row<len(matrix)-1 && strings.ContainsAny(curr,"S|F7") && strings.ContainsAny(string(matrix[row+1][col]),"|LJ") && !slices.Contains(seen,Pair{row+1,col}){
			//fmt.Print("Here2\n")
			queue = append(queue, Pair{row+1,col})
			seen = append(seen, Pair{row+1,col})
		}

		if col>0 && strings.ContainsAny(curr,"S-J7") && strings.ContainsAny(string(matrix[row][col-1]),"FL-") && !slices.Contains(seen,Pair{row,col-1}) {
			//fmt.Print("Here3\n")
			queue = append(queue, Pair{row,col-1})
			seen = append(seen, Pair{row,col-1})
		}

		if col<len(matrix[0])-1 && strings.ContainsAny(curr,"FL-S") && strings.ContainsAny(string(matrix[row][col+1]),"-7J") && !slices.Contains(seen,Pair{row,col+1}){
			//fmt.Print("Here4\n")
			queue = append(queue, Pair{row,col+1})
			seen = append(seen, Pair{row,col+1})
		}
	}

	// The path is guaranteed to be one continuous loop.
	// So the tile farthest from the Start is halfway across the map. What this means is since the start and the end are the same, the farthest you can get from the Start tile is the tile numbered len(seen)/2

	//fmt.Print(seen)

	return len(seen)/2
}


func findStart(matrix []string) (int,int){

	//fmt.Print(matrix[0])
	r:=0
	c:=0
	for i:=0;i<len(matrix);i++ {
		for j:=0;j<len(matrix[0]);j++ {
			if matrix[i][j]=='S' {
				r=i
				c=j
			}
		}
	}
	return r,c
}
func main() {
	file, err := os.Open("./input.txt")
	if err!=nil {
		log.Fatal("Error opening file")
	}

	defer file.Close()

	var matrix []string

	scanner:=bufio.NewScanner(file)

	for scanner.Scan() {
		line:=scanner.Text()
		//lineSlice:=strings.Fields(line)
		matrix = append(matrix, line)
	}

	// for _,ele:=range matrix {
	// 	fmt.Println(ele)
	// }

	sr,sc:=findStart(matrix)

//	fmt.Print(sr,sc)

	ans1:=part1(matrix,sr,sc)

	fmt.Printf("Answer to part1 is %v\n",ans1)

}