package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	
)

func printMatrix(matrix [][]int) {
	for _,row:=range matrix {
		fmt.Print(row,"\n")
	}
}

func inspect(matrix [][]int, i int, j int,j_start int) bool { // 0 2 0
	dirs:=[]int {-1,0,1}
	
	for _,x:=range dirs {
		for _,y:=range dirs {
			for z:=j_start;z<j;z++ {
				index1:=i+x
				index2:=z+y

				if index1>=0 && index2>=0 && index1<len(matrix) && index2<len(matrix[0]) {
					if matrix[index1][index2] != 46 && (matrix[index1][index2] < 48 || matrix[index1][index2] > 57) { //adjacent to symbol
						return true
					} 
				}
			}


		}
	}

	return false

}

func part1(matrix [][]int) int{
	count:=0
	for i:=0;i<len(matrix);i++ {
		j:=0
		j_start:=0
		
		num:=0
		for j<len(matrix[0]) {
			if matrix[i][j]<48 || matrix[i][j]>57 {
				//fmt.Print("Here ",num,"\n")
				if num!=0 {
					canAdd:=inspect(matrix,i,j,j_start) // 0 2 0

					if canAdd {
						count+=num
					}
				}
				num=0
				j+=1
				j_start=j
				continue

			} else {
				num=num*10+(matrix[i][j]-48)
			}

			j+=1

		}

		if num!=0 { //If the number is not 0 by the end of the row, i.e the number is the last element of the row
			canAdd:=inspect(matrix,i,j,j_start) // 0 2 0

			if canAdd {
				count+=num
			}
		}
		

	}
	return count
}

func getStars(matrix[][] int) map[int][]int {
	m:=make(map[int][]int)

	//fmt.Printf("Number of columns is %d\n",len(matrix[0]))
	for i:=0;i<len(matrix);i++ {
		for j:=0;j<len(matrix[i]);j++ {
			if matrix[i][j]==42 {
				m[i]=append(m[i], j)
			}
		}
	}

	return m
}

type Pair struct {
	left int
	right int
}

func contains(s []Pair, p Pair) bool {
	for _,v:=range s {
		if v.left==p.left && v.right==p.right {
			return true
		}
	}
	return false
}
func find(matrix[][]int,key int,value int) int{

	//fmt.Printf("Called with key %d and value %d\n",key,value)
	dirs:=[]int {-1,0,1}
	visited:=make(map[int][]Pair)
	for _,x:=range dirs {
		for _,y:=range dirs {
			i:=key+x
			j:=value+y

			if i>=0 && j>=0 && i<len(matrix) && j<len(matrix[0]) {
				if matrix[i][j] >=48 && matrix[i][j]<=57 {
					left:=j
					right:=j

					for left>=0 && (matrix[i][left] >=48 && matrix[i][left]<=57 ) {
						left-=1
					}
					left+=1

					for right<len(matrix[0]) && (matrix[i][right] >=48 && matrix[i][right]<=57) {
						right+=1
					}

					if !contains(visited[i],Pair{left: left,right: right}) {
						visited[i]=append(visited[i], Pair{left: left,right: right})
					}
					//fmt.Print("Row ",i,": ",left," ",right,"\n")

				}
			}
		}
		//fmt.Print("\n")
	}

	//fmt.Print(visited," ",len(visited),"\n")
	res:=1
	if len(visited)==2 { //We need two adjacent elements from different rows
		for key,values:=range visited {
			if len(values) !=1 { // Ensure one row has only one adjacent element, else we will have exceeded 2
				return 0
			}

			tempnum:=0

			start:=values[0].left
			end:=values[0].right

			row:=key

			//fmt.Printf("Row %d Start %d End %d\n",row,start,end)

			for s:=start;s<end;s++ {
				tempnum=tempnum*10 + (matrix[row][s]-48)
			}
			res*=tempnum
			//fmt.Print("Tempnum is ",tempnum,"\n")
		}
	} else if len(visited)==1{ //In case both the adjecent elements are on the same row
	
			for key,values:= range visited {
				if len(values) ==2 { //We need to have exactly 2
					//fmt.Print("Im here\n")
					tempnum:=0

					start:=values[0].left
					end:=values[0].right

					row:=key

					//fmt.Printf("Row %d Start %d End %d\n",row,start,end)

					for s:=start;s<end;s++ {
						tempnum=tempnum*10 + (matrix[row][s]-48)
					}
					//fmt.Print("tempnum ",tempnum,"\n")
					res*=tempnum

					tempnum=0
					start=values[1].left
					end=values[1].right

					row=key

					//fmt.Printf("Row %d Start %d End %d\n",row,start,end)

					for s:=start;s<end;s++ {
						tempnum=tempnum*10 + (matrix[row][s]-48)
					}
					res*=tempnum
					
				} else {
					return 0
				}

				
			}
	} else {
		return 0
	}

	//fmt.Print("Result is ",res,"\n")

	return res
}

func part2(matrix[][]int) int{
	stars:=getStars(matrix)

	//fmt.Print(stars,"\n")
	ratio:=0
	for key,values:= range stars {
		for _,v:=range values {
			//fmt.Print("Ignore ",key,v)
			ratio+=find(matrix,key,v)
		}
		//return 
	}
	return ratio
}

func main() {
	file,err:=os.Open("./symbols.txt")
	if err!=nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner:=bufio.NewScanner(file)
	var matrix [][]int
	for scanner.Scan() {
		lineBytes := scanner.Bytes()
		lineInts := make([]int, len(lineBytes)) 

		for i, b := range lineBytes {
			lineInts[i] = int(b) 
		}

		matrix = append(matrix, lineInts)
	}

	//printMatrix(matrix)

	ans1:=part1(matrix)

	fmt.Print("ans for part1 is ",ans1,"\n")

	ans2:=part2(matrix)

	fmt.Print("Answer to part2 is ",ans2,"\n")
	
}