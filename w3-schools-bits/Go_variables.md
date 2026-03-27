# variables
In go variables are containers for storing data values.


## Go Variable Types
## integer
Intergers are used to store numbers

## bool
Bool are used to store value two states: true or false

## float
Float is use to store decimal

## string
String is used to store text



## Declaring (Creating) Variables
1. With the var keyword:
Use the var keyword, followed by variable name and type:
Syntax
var variablename type = Emmy nice

## Variable Declaration With Initial Value
If the value of a variable is known from the start, you can declare the variable and assign a value to it on one line

# Go Variable Naming Rules 

A variable name must start with a letter or an underscore character (_)
A variable name cannot start with a digit
A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
Variable names are case-sensitive (age, Age and AGE are three different variables)
There is no limit on the length of the variable name
A variable name cannot contain spaces
The variable name cannot be any Go keywords

# Multi-Word Variable Names

Variable names with more than one word can be difficult to read.

There are several techniques you can use to make them more readable

# Camel Case

Each word, except the first, starts with a capital letter:
myVariableName = "John"
## Pascal Case

Each word starts with a capital letter:
MyVariableName = "John"
## Snake Case

Each word is separated by an underscore character:
my_variable_name = "John"