package main

import "fmt"

func square() {
	fmt.Print("Please select your square size : ")
enterAgain:
	b:=0
	fmt.Scanln(&b)
	if b<1 || b>20 {
		fmt.Print("Please select your square size again : ")
		goto enterAgain
	}
	for i:=1; i<=b; i++{
		if i==1 || i==b {
			for j:=0; j<b; j++ {
				fmt.Print("* ")
			}
		} else{
			fmt.Print("* ")
			for j:=1; j<b-1; j++ {
				fmt.Print("  ")
			}
			fmt.Print("*")
		}
		fmt.Println()
	}
}