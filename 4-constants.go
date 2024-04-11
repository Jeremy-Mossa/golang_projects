package main

import (
    "fmt"
    "math"
)

const s string = "non-regexp"

func main() {
    fmt.Println(s)    

    const denominator = 3000000
    const numerator = 3e17
    const decimated_fraction = numerator/denominator

    fmt.Println(decimated_fraction)

    fmt.Println(int64(decimated_fraction))

    fmt.Println(math.Sin(denominator))
}
