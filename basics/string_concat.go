package main

import "fmt"
import "strconv"

func main() {
  a := "my_string"

  // concat with int
  i := 2015
  b := a + strconv.Itoa(i)
  fmt.Println(b)

  // concat with float
  var d float64 = 3.1416
  c := a + strconv.FormatFloat(d, 'f', 4, 64)
  fmt.Println(c)

  // ref
  // http://golang.org/pkg/strconv/
  // https://blog.golang.org/strings
}
