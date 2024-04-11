package main

import (
  "fmt"
  "time"
)

func main() {
  i := 2
  fmt.Println("Write ", i, " as ")
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("three")
  }

  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend")
  default:
    fmt.Println("It's a weekday")
  }

  t := time.Now()
  switch {
  case t.Hour() > 17:
    fmt.Println("It's evenin'")
  default:
    fmt.Println("It ain't evenin'")
  }

  whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
      fmt.Println("boolean - bool")
    case int:
      fmt.Println("integer - int")
    case string:
      fmt.Println("string")
    default:
      fmt.Println("I am of an unknown type", t)
    }
  }

  whatAmI("hey there")
  whatAmI(17)
  whatAmI(false)
}
