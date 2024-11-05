
# Automate testing 


Prepare code for testing. Add logic code in a separate function to make it easier to test.

```go

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

	fmt.Println("Your name is ", fullName, " and you are age ", age)

	if checkNDISAge(age) {
		fmt.Println("You are eligible for NDIS")
	} else {
		fmt.Println("You are not eligible for NDIS")
	}
}

func checkNDISAge(age int) bool {
	if age < 65 {
		return true
	}
	return false
}

```




Create a file called `main_test.go` and add the following code:

```go
import (
	"testing"
)

func TestNDISAgeCheck(t *testing.T) {

	var agetests = []struct {
		in    int
		label string
		out   bool
	}{
		{in: 64, label: "64", out: true},
		{in: 65, label: "65", out: false},
	}

	for _, tt := range agetests {
		t.Run(tt.label, func(t *testing.T) {
			s := checkNDISAge(tt.in)

			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}
```



<br>
<hr>
<br>  

## Practice  


## Extend NDIS check

Add minimum age of 7 to be eligible for NDIS.
Add tests to verify.
How much test data should you have?


## Driving Age Check

Extend this example to also ask people for the country they live in. 
Create a function that checks driving age. Create a test for this function using the following test data.
```go
	var agetests = []struct {
		age     int
		country string
		out     bool
	}{
		{age: 16, country: "Australia", out: false},
		{age: 18, country: "Australia", out: true},
        {age: 15, country: "USA", out: false},
		{age: 16, country: "USA", out: true},
	}
```


<br />

<hr />

<br />  


[Next: 05_json](05_json.md)

