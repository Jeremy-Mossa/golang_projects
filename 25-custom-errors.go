package main

import (
  "fmt"
  "errors"
)

var pln = fmt.Println

type argError struct {
  arg int
  message string
}

func (e *argError) Error() string {
  return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
  if arg == 42 {
    return -1, &argError{arg, "too old, no jobs for you"}
  }
  return arg + 3, nil
}
func main() {
  _, err := f(42)
  var ae *argError
  if errors.As(err, &ae) {
      pln(ae.arg)
      pln(ae.message)
  } else {
      pln("err failed to match argError")
  }
}
