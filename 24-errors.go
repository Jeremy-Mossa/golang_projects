package main

import (
  "fmt"
  "errors"
)

var pln = fmt.Println

func f(arg int) (int, error) {
  if arg == 42 {
    return -1, errors.New("can't work with 42")
  }
  return arg + 3, nil
}

var ErrOutofTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
  if arg == 2 {
    return ErrOutofTea
  } else if arg == 4 {
    return fmt.Errorf("making tea: %w, ErrPower")
  }
  return nil
}
func main() {
  for _, i := range []int{7, 42} {
    if r, e := f(i); e != nil {
      pln("f failed:", e)
    } else {
        pln("f worked:", r)
    }
  }

  for i := range 5 {
    if err := makeTea(i); err != nil {
      if errors.Is(err, ErrOutofTea) {
        pln("probably buy new tea or suffer caffeine withdrawal ... up2u")
      } else if errors.Is(err, ErrPower) {
          pln("electrical failure. evacuate")
      } else {
        fmt.Printf("unknown error: %s\n", err)
      }
    }
    pln("tea is seeped")
  }
}
