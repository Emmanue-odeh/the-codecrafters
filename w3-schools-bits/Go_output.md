## Go Output Functions

Go has three functions to output text:

    Print()
    Println()
    Printf()

## The Print() Function

The Print() function prints its arguments with their default format.

Print the values of i and j:
``` go

package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i)
  fmt.Print(j)
}
```

## The Println() Function

The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end:
```go
package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Println(i,j)
}
```

## The Printf() Function

The Printf() function first formats its argument based on the given formatting verb and then prints them.

Here we will use two formatting verbs:

    %v is used to print the value of the arguments
    %T is used to print the type of the arguments
```go
package main
import ("fmt")

func main() {
  var i string = "Hello"
  var j int = 15

  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T", j, j)
}
```