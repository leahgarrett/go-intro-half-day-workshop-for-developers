

At the command line

Make a new folder
`mkdir hello`

Move into the folder
`cd mkdir`

Create a go file
`touch main.go`

Create the package
`go mod init gophercon/hello`

Add the following code to `main.go`

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello World")
}
```

`fmt` is part of the Go standard library.

The `fmt` package includes `Println`


## String Variables

```go
package main

import "fmt"

func main() {
  name := "Leah"
  fmt.Println("Hello", name)
}
```


<br>

## Numeric Variables

Arithmetic Operators: +, -, *, /, %

```go
package main

import "fmt"

func main() {
  x := 20
  y := 30
  fmt.Println("x =", x, "y =", y)
  fmt.Println("x + y =", x + y)
}
```

Integer calculations have no decimal 

If we want to use a decimal we can declare variables with decimals i.e. as float 


```go
package main

import "fmt"

func main() {
  x := 20.0
  y := 30.0
  fmt.Println("x =", x, "y =", y)
  fmt.Println("x + y =", x + y)
  fmt.Println("x / y =", x / y)
}
```

Output
```
x = 20 y = 30
x + y = 50
x / y = 0.6666666666666666
```

We have been creating variables inferring the type
`x := 10`

We can specify the type  
`var x float64 = 10`


We can check what type variables are using `%T`
`fmt.Printf("%v %T", x, x)`



## Type safety
You cannot change the type of a variable one it has been declared.

```go
package main

import "fmt"

func main() {
  x := "Leah"
  y := 30.0
  fmt.Println("x =", x, "y =", y)
  fmt.Println("x + y =", x + y)
  fmt.Println("x / y =", x / y)
}
```

## Comments
You can add notes to your code using comments. 

You can use a comment to hide code and test things. 

```go
package main

import "fmt"

func main() {
  x := 7 // this variable is created as an integer
  y := 3

  // this is the modulus operator
  fmt.Println("x % y =", x % y) 
}
```

<br />  

## Reading user input


```go
package main

import "fmt"

func main() {
    justRight := 88
	var temperature int
	fmt.Println("What temperature is the soup?")
	fmt.Scan(&temperature)
}
```

<br />

## Conditional statements

Comparison Operators: ==, !=, <, >, <=, >=


Checking the result


```go
package main

import "fmt"

func main() {
	justRight := 88
	var temperature int
	fmt.Println("What temperature is the soup?")
	fmt.Scan(&temperature)

	if temperature == justRight {
		fmt.Println("The soup is just right")
	}
}
```

<br />

The scan function returns multiple values. We can check for errors.

```go
package main

import "fmt"

func main() {
	justRight := 88
	var temperature int
	fmt.Println("What temperature is the soup?")
	result, err := fmt.Scan(&temperature)

	if err != nil {
		fmt.Println("Error reading temperature:", err)
	}

	fmt.Println("Result of scan: ", result)

	if temperature == justRight {
		fmt.Println("The soup is just right")
	}
}
```

Can choose to panic when there is an error
```go
	panic(err)
```

Or add the 'Log" package and log a fatal error
```go
    log.Fatal(err)
```

If we do not need to use the value of result we can ignore it by using `_`
```go
_, err := fmt.Scan(&temperature)
```


## What is the `mod` file?

We used `go mod init` to create a mod file.

### What is it?
```
Each Go module is defined by a go.mod file that describes the module's properties, including its dependencies on other modules and on versions of Go.  
```
https://go.dev/doc/modules/gomod-ref

```
mod file lists the specific versions of the dependencies that your project uses. The go. sum file provides checksums for the exact contents of each dependency at the time it is added to your module. The dependencies are stored in the Go module cache, which is shared across all projects on your system.
```
https://www.freecodecamp.org/news/golang-environment-gopath-vs-go-mod/#:~:text=mod%20file%20lists%20the%20specific,all%20projects%20on%20your%20system.




<hr>

<br>

## Practice

### Calculate area

`math.Pi` is a pre-defined constant in Go's `math` package

What does the following code output?
What does the '\n' do?
What would '\t' do?

```go
package main

import "fmt"
import "math"

func main() {
	fmt.Printf("%.48f\n", math.Pi)
	fmt.Printf("%v\n", math.Pi)
	fmt.Printf("%T", math.Pi)
}
```

The area of a circle is Pi times radius squared.
Add code to ask the user to enter the radius and display the area of the circle.

Display the answer with two decimal places.

How did you declare the radius variable?
Can you declare it as an integer?



<br />

<hr />

<br />  


[Next: 02_loops_and_slices](02_loops_and_slices.md)




