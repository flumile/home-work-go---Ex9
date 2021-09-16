package main

import (
	"fmt"
	"strconv"
	"math"
	"time"
	"math/rand"
)

func choosePal(){
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	i:=r1.Intn(10)
	if i%2 == 0{
		fmt.Println("• Palindrome Method 1 •")
		palindrome()
	} else {
		fmt.Println("• Palindrome Method 2 •")
		palindrome2()
	}
}

//method 1

func palindrome() {
	num:=0
enterNum:
	fmt.Print("Enter Number : ")
	fmt.Scanln(&num)
	if num<10000||num>100000{
		fmt.Println("Enter Number with 5 digits!!!")
		goto enterNum
	}
	num2 := strconv.Itoa(num)
	if reverse(num2) == num2 {
		fmt.Println("It's palindrome")
	} else {
		fmt.Println("It isn't palindrome")
	}
}

func reverse(num string) (rvNum string) {
	for _, v := range num {
		rvNum = string(v) + rvNum
	}
	return
}

//method 2

func palindrome2(){
	num:=0
enterNum:
	fmt.Print("Enter Number : ")
	fmt.Scanln(&num)
	a:=countDigit(num)
	if a!=5 {
		fmt.Println("Please Enter number with 5 digits!!")
		goto enterNum
	}
	if checkPal(num){
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}
}

func countDigit(num int) (cnt int) {
	if num < 10 {
		return 1
	} else {
		return 1 + countDigit( num/10 )
	}
}

func checkPal(num int) (pal bool) {
	var reNum [5]int
	num2:=num
	for i:=0; i<5; i++ {
		reNum[i] = num2%10
		num2=num2/10
	}
	rvNum:=0
	for i:=0; i<5; i++ {
		echn:=reNum[i] * int(math.Pow10(4-i))
		rvNum+=echn
	}
	fmt.Println("reverse ",rvNum)
	if num==rvNum {
		return true
	} else {
		return false
	}
}