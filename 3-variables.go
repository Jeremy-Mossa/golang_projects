package main

import "fmt"

func main() {
    var a = "some string"
    fmt.Println(a)

    var b, c = true, false
    fmt.Println(b, c)

    var d int = 3
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "only double-quotes inside the Println method"
    fmt.Println(f)
}
