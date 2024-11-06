package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var age int
	fmt.Println("What is your age?")
	fmt.Scan(&age)

	fmt.Println("Hello World", age)

	// reading a string
	reader := bufio.NewReader(os.Stdin)
	var fullName string = ""
	fmt.Println("What is your full  name?")
	fullName, _ = reader.ReadString('\n')

	// reading a string
	reader = bufio.NewReader(os.Stdin)
	var country string = ""
	fmt.Println("What country are you in?")
	country, _ = reader.ReadString('\n')

	fmt.Println("Your name is ", fullName, " and you are age ", age)

	if checkNDISAge(age) {
		fmt.Println("You are eligible for NDIS")
	} else {
		fmt.Println("You are not eligible for NDIS")
	}

	if checkDrivingAge(age, country) {
		fmt.Println("You are eligible to drive")
	} else {
		fmt.Println("You are not eligible to drive")
	}
}

func checkNDISAge(age int) bool {
	if age >= 7 && age <= 65 {
		return true
	}
	return false
}

func checkDrivingAge(age int, country string) bool {
	if country == "Australia" && age >= 18 {
		return true
	}
	if country == "USA" && age >= 16 {
		return true
	}
	return false
}
