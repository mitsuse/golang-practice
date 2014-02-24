package main

import "fmt"

func fibonacci() func() int {
    x, y := 0, 1
    next := func() int {
        z := x
        x, y = y, x + y
        return z
    }
    return next
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
