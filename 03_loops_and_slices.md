
# Loops

Increment and Decrement Operators: ++, --


```go
package main

import (
	"fmt"
)

func main() {
  index := 0

  fmt.Println("Index:", index)

  index++
  fmt.Println("Index:", index)

  index++
  fmt.Println("Index:", index)

  index++
  fmt.Println("Index:", index)
}
```

We can use a loop to repeat a set number of times.
```go
package main

import (
	"fmt"
)

func main() {
  index := 0
  max := 10

  fmt.Println("Before loop:", index)

  for index < max {
    index++
    fmt.Println("Index:", index)
  }

}
```

We can use loop shorthand 
```go
package main

import (
	"fmt"
)

func main() {
  max := 10

  fmt.Println("Before loop:", index)

  for index := 0; index < max; index++ {
    fmt.Println("Index:", index)
  }
}
```

Now `index` is only accessible inside the loop. Use the loop short hand when you need to reference the loop index inside the loop only.

How would we change this to count down?

<br>
<hr>
<br>  



# Slices

Go provides slices to store a collection of data

```go
package main

import "fmt"

func main() {
  letters := []string{"a", "b", "c", "d"}
	
  fmt.Println("First in the slice",letters[0])
}
```


<br />

```
A slice has both a length and a capacity.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
```
https://go.dev/tour/moretypes/11  

<br />


We can use a for loop to look at each element in the slice.

```go
package main

import "fmt"

func main() {
  letters := []string{"a", "b", "c", "d"}
	
  fmt.Println("First in the slice",letters[0])
  fmt.Println("Number of elements in the letters slice",len(letters))
  fmt.Println("The capacity of the slice", cap(letters))
  
  for index :=0; index < len(letters); index++ {
      fmt.Println(index, "in the slice",letters[index])
  }
}
```

We can change or update elements in the slice

```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  letters := []string{"a", "b", "c", "d"}
	
  fmt.Println("First in the slice",letters[0])
  fmt.Println("Number of elements in the letters slice",len(letters))
  fmt.Println("The capacity of the slice", cap(letters))

  for index :=0; index < len(letters); index++ {
      letters[index] = strings.ToUpper(letters[index])
      fmt.Println(index, "in the slice",letters[index])
  }
}
```

## Appending to a slice
Go will make a new slice from our existing slice and the new element  
`letters = append(letters, "e")`


## Make a function for repeated code

```go
package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c", "d"}
	printSlice(letters)

	letters = append(letters, "e")
	printSlice(letters)

	letters = append(letters, "f")
	printSlice(letters)
}

func printSlice(s []string) {
	fmt.Println("Number of elements in the letters slice", len(s))
	fmt.Println("The capacity of the slice", cap(s))

	for index := 0; index < len(s); index++ {
		fmt.Println(index, "in the slice", s[index])
	}
}
```


<br />

<br />




## Practice  

### Calculate area
Extend the calculate example to loop 5 times asking for the radius and displaying the area 5 times.

Refactor to add a function for the calculation.

How could you change the loop to continually ask for the radius?

Create another version that uses a slice of values for radius.



### Fruit
Create a new version of our `letters` example above using the following slice 
`fruit := []string{"apple", "Banana", "cherry", "orange", "watermelon", "Mango", "Papaya", "Blueberry", "plum", "Peach", "pear", "Pineapple", "Raspberry", "Strawberry", "lemon", "lime"}`

Use append to add three more fruit.

Create another loop that displays the values using `Title` instead of `ToUpper`
And another using `ToLower`

Create another loop to display the fruit on one line using `Print` instead of `Println`


<br />

<hr />

<br />  


[Next: 04_tests](04_tests.md)

