package main

import(
	"os"
	"bufio"
	"fmt"
	"unicode"
	"strconv"
	"regexp"
)

func letter(s string) (int,int) {

	//numbers:=[]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	//two
	start:=0
	//fmt.Printf("String is %s",s)
	//fmt.Print("\n")
	for start<len(s){
		j:=start
		firstChar:=string(s[j])
		start_regex:=regexp.MustCompile(`(?i)[otfsen]`)

		if start_regex.MatchString(firstChar) {
			
			//fmt.Print("Starts well")

			switch firstChar {
				case "o":
					if j+2 >=len(s) {
						start+=1
						break
					}
					if string(s[j+1])=="n" && string(s[j+2])=="e"{
						return j,1
					}
					start+=1

				case "t":
					if j+2>=len(s) {
						start+=1
						break
					}

					if string(s[j+1])=="w" && string(s[j+2])=="o" {
						return j,2
					}
					if j+4>=len(s) {
						start+=1
						break
					}

					sliced:=s[j+1:j+5]
					if sliced=="hree" {
						return j,3
					}
					start+=1
				
				case "f":
					if j+3>=len(s) {
						start+=1
						break
					}

					sliced:=s[j+1:j+4]
					if sliced=="our" {
						return j,4
					} else if sliced=="ive" {
						return j,5
					}
					start+=1

				case "s":
					if j+2>=len(s) {
						start+=1
						break
					}
					sliced:=s[j+1:j+3]
					if sliced=="ix" {
						return j,6
					} else{
						if j+4>=len(s) {
							start+=1
							break
						}
						sliced=s[j+1:j+5]
						if sliced=="even" {
							return j,7
						}
					}
					start+=1
				case "e":
					if j+4>=len(s) {
						start+=1
						break
					}
					sliced:=s[j+1:j+5]
					if sliced=="ight" {
						return j,8
					}
					start+=1
				case "n":
					if j+3>=len(s) {
						start+=1
						break
					}
					sliced:=s[j+1:j+4]
					if sliced=="ine" {
						return j,9
					}
					start+=1
				default:
					start+=1

			}
			
		} else {
			//fmt.Print("Nope")
			start+=1
		}
		
	}

	return -1,-1
}

func reverse_letter(s string) (int,int) {
	start:=len(s)-1
	//numbers:=[]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for start>=0 {
		j:=start
		firstChar:=string(s[j])

		start_regex:=regexp.MustCompile(`(?i)[eorxnt]`)

		if start_regex.MatchString(firstChar) {
			switch firstChar {
				case "e":
					if j-2<0 {
						start-=1
						break
					} 
					sliced:=s[j-2:j]
					if sliced =="on" {
						//fmt.Print("This is 1")
						return j,1
					}
					
					if j-3<0 {
						start-=1
						break
					}
					sliced=s[j-3:j]
					if sliced == "fiv" {
						//fmt.Print("This is 5")
						return j,5
					} else if sliced == "nin" {
						//fmt.Print("This is 9")
						return j,9
					}

					if j-4<0 {
						start-=1
						break
					}
					sliced=s[j-4:j]
					if sliced == "thre" {
						//fmt.Print("This is 3")
						return j,3
					}

					start-=1

				case "o":
					if j-2<0 {
						start-=1
						break
					}
					sliced:=s[j-2:j] 
					if sliced=="tw" {
						//fmt.Print("This is 2")
						return j,2
					}
					start-=1
				
				case "r":
					if j-3<0 {
						start-=1
						break
					}
					sliced:=s[j-3:j]
					if sliced=="fou" {
						//fmt.Print("This is 4")
						return j,4
					}
					start-=1

				case "x":
					if j-2 < 0 {
						start-=1
						break
					}

					sliced:=s[j-2:j]
					if sliced=="si" {
						//fmt.Print("this is 6")
						return j,6
					}
					start-=1

				case "n":
					if j-4<0 {
						start-=1
						break
					}
					sliced:=s[j-4:j]
					if sliced=="seve" {
						//fmt.Print("this is 7")
						return j,7
					}
					start-=1
				
				case "t":
					if j-4<0 {
						start-=1
						break
					}
					sliced:=s[j-4:j]
					if sliced=="eigh" {
						//fmt.Print("This is 8")
						return j,8
					}
					start-=1

				default:
					start-=1


			}
		} else {
			start-=1
		}

	}
	return -1,-1
}

func main() {
	file,err:=os.Open("./calib.txt")
	if err!=nil{
		fmt.Print(err)
		return
	}

	defer file.Close()

	scanner:=bufio.NewScanner(file)

	count:=0

	

	for scanner.Scan(){
		s:=scanner.Text()
		
		num1:=0
		num2:=0

		idxnum1:=0
		idxnum2:=0
		for i:=0;i<len(s);i++ {
			if unicode.IsDigit(rune(s[i])) {
				idxnum1=i
				break
			}
		}
		letteridx,letternum:=letter(s)


		

		//fmt.Printf("The forward letter number is %d and letter index is %d\n",letternum,letteridx)

		if letteridx!=-1 && letteridx <= idxnum1 {
			num1=letternum
		} else {
			num1,_=strconv.Atoi(string(s[idxnum1]))
		}

		reverseidx,reversenum:=reverse_letter(s)

		//fmt.Printf("The reverse number is %d and reverse index is %d\n",reversenum,reverseidx)

		for j:=len(s)-1;j>=0;j--{
			if unicode.IsDigit(rune(s[j])) {
				idxnum2=j
				//num2,_=strconv.Atoi(string(s[j]))
				break
			}
		}	

		if reverseidx!=-1 && reverseidx>idxnum2 {
			num2=reversenum
		} else {
			num2,_=strconv.Atoi(string(s[idxnum2]))
		}

		//fmt.Printf("Num1 is %d and num2 is %d\n",num1,num2)
		
		//fmt.Printf("The number is %d\n",num1*10+num2)
		count+=num1*10+num2

		

	}

	fmt.Printf("Result is %d",count)
}