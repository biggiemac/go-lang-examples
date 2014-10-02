package main

import "fmt"

func main() {

    for i := 1; i <= 100; i++ {
        switch i {
        case i % 3 == 0 : fmt.Print("Fizz")
        case i % 5 == 0 : fmt.Print("Buzz")
        default: fmt.Print(i)
        }

        fmt.Println()
    }
}
