## Go Slices

Slices are similar to arrays, but are more powerful and flexible.

Like arrays, slices are also used to store multiple values of the same type in a single variable.

However, unlike arrays, the length of a slice can grow and shrink as you see fit.

In Go, there are several ways to create a slice:

    Using the []datatype{values} format
    Create a slice from an array
    Using the make() function

## Create a Slice With []datatype{values}
Syntax
slice_name := []datatype{values}

A common way of declaring a slice is like this:
myslice := []int{}


## Go Access, Change, Append and Copy Slices
Access Elements of a Slice

You can access a specific slice element by referring to the index number.

In Go, indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.

## Change Elements of a Slice

You can also change a specific slice element by referring to the index number.

## Append Elements To a Slice

You can append elements to the end of a slice using the append()function:
Syntax
slice_name = append(slice_name, element1, element2, ...)

