## Go Constants
The const keyword declares the variable as "constant", which means that it is unchangeable and read-only. If a variable should have a fixed value that cannot be changed, you can use the const keyword.

# Declaring a Constant
Example of declaring a constant in Go:

```GO
package main

import (
	"fmt"
)
const PI = 5.6

func main() {
	fmt.Println(PI)
}
```
## Constant Rules

    Constant names follow the same naming rules as variables
    Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
    Constants can be declared both inside and outside of a function

## Constant Types

There are two types of constants:

    Typed constants
    Untyped constants

## Typed Constants

Example of Typed constants declared with a defined type:

```go

package main
import ("fmt")

const A int = 1

func main() {
  fmt.Println(A)
}
```

## Untyped Constants

Example of Untyped constants declared without a type:
```go

package main
import ("fmt")

const A = 1

func main() {
  fmt.Println(A)
}
```