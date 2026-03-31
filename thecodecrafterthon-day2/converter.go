package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func hexTodecimal(hex string) {

    n, err := strconv.ParseInt(hex, 16, 64)
    if err != nil {
        fmt.Println("invalid input")
    }
    fmt.Println("Decimal: ", n)
}

func binTodecimal(bin string) (int64, error) {
    n, err := strconv.ParseInt(bin, 2, 64)
    if err != nil {
        fmt.Println("invalid input")
    }
    fmt.Println("Decimal:", n)
    return strconv.ParseInt(bin, 2, 64)
}

func convertTodecimal(num string) (int64, error) {
    n, err := strconv.ParseInt(num, 10, 64)
    if err != nil {
        fmt.Println("invalid input")
    }
    fmt.Printf("binary: %b \n", n)
    fmt.Printf("hex: %X \n", n)
    return strconv.ParseInt(num, 10, 64)
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Enter convertion ")

        fmt.Print("> ")
        if !scanner.Scan() {

            break
        }
        input := strings.TrimSpace(scanner.Text())
        if input == "" {
            continue
        }
        if input == "quit" {
            fmt.Println("Good bye!!")
            break
        }
        conv := strings.Fields(input)
        if len(conv) != 3 || conv[0] != "convert" {
            fmt.Println("Invalid operation use convert value, base")
            continue
        }
        value := conv[1]
        Base := conv[2]
        if Base == "hex" {
            hexTodecimal(value)
            continue
        }
        if Base == "bin" {
            binTodecimal(value)
            continue
        }
        if Base == "dec" {
            convertTodecimal(value)
            continue
        }

    }
}



