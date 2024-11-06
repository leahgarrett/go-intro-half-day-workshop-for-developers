package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c", "d"}
	printSlice(letters)

	letters = append(letters, "e")
	printSlice(letters)

	letters = append(letters, "f")
	printSlice(letters)

	letters = append(letters, "g")
	printSlice(letters)

	letters = append(letters, "h")
	printSlice(letters)

	letters = append(letters, "i")
	printSlice(letters)
}

func printSlice(s []string) {
	fmt.Println("Number of elements in the letters slice", len(s))
	fmt.Println("The capacity of the slice", cap(s))

	for index := 0; index < len(s); index++ {
		fmt.Println(index, "in the slice", s[index])
	}
}
