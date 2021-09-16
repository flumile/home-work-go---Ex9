package main

import (
	"fmt"
)

func main() {
start:
	fmt.Println("What's do you want to do ?")
	fmt.Println("Enter 1 : Square of Asterisks")
	fmt.Println("Enter 2 : Check Palindromes")
	fmt.Println("Enter 3 : Find the Smallest Value")
	fmt.Println("Enter 0 : To Exit")
	fmt.Print("☻ Please Select : ")
	numCase:=0
	fmt.Scanln(&numCase)
	fmt.Println("----------------------------")
	if numCase == 1 {
		fmt.Println("♠ Square of Asterisks ♠")
		square()
	} else if numCase == 2 {
		fmt.Println("♥ Palindromes ♥")
		choosePal()
	} else if numCase == 3 {
		fmt.Println("♦ Find the Smallest Value ♦")
		small()
	} else if numCase == 0 {
		return
	} else {
		fmt.Println("Please Select again!!!")
		goto start
	}
exit:
	fmt.Print("Do you want to exit? (Y/N) ")
	var ans string
	fmt.Scan(&ans)
	if ans == "Y" || ans == "y" {
		return
	} else if ans == "N" || ans == "n" {
		fmt.Println("----------------------------")
		goto start
	} else {
		fmt.Println("Please Enter Y/N")
		goto exit
	}
}