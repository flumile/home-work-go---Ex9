package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func small() {
enterAgain:
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter number : ")
	scanner.Scan()
	numsStr := strings.Split(scanner.Text()," ")
	nums, wrongTxt := eachNum(numsStr)
	if wrongTxt {
		fmt.Println("Please Enter number!!!")
		goto enterAgain
	}
	smallest := getSmallest(nums)
	fmt.Println("The Smallest value is :",smallest)
}

func eachNum(num []string) (val []int , wrongTxt bool){
	wrongTxt = false
	for i := range num{
		reNum, err := strconv.Atoi(num[i])
		if err != nil {
			wrongTxt = true
			break
		}
		val = append(val,reNum)
	}
	return
}

func getSmallest(num []int) (min int) {
	min=num[0]
	for i:=1; i<len(num); i++{
		if num[i]<min {
			min=num[i]
		}
	}
	return
}