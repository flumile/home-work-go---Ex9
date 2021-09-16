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

func ex1() {
	a:=0
	b:=1
	fmt.Print(a," ")
	fmt.Print(b," ")
	for i:=0; i<20; i++ {
		c:=a+b
		a=b
		b=c
		fmt.Print(b," ")
	}
	fmt.Println()
}
func ex2() {
	for a:=1; a<=30; a++ {
		fmt.Print(a," ")
		if a%3==0 && a%5==0 {
			fmt.Println("FizzBuzz")
		}else if a%3==0 {
			fmt.Println("Fizz")
		} else if a%5==0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println()
		}
	}
}

func ex3() {
	for i:=0; i<=5; i++ {
		for j:=0; j<i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func ex4() {
	for i:=0; i<=5; i++ {
		for j:=0; j<=5-i; j++ {
			fmt.Print(" ")
		}
		for k:=0; k<i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func ex5() {
	for i:=0; i<5; i++ {
		for j:=0; j<=5-i; j++ {
			fmt.Print(" ")
		}
		for k:=0; k<i; k++ {
			fmt.Print("*")
		}
		for k:=1; k<i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for b:=5; b>=0; b-- {
		for j:=5; j>=b; j-- {
			fmt.Print(" ")
		}
		for k:=b; k>0; k-- {
			fmt.Print("*")
		}
		for k:=b; k>1; k-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}