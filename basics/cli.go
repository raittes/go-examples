package main

import "fmt"
import "flag"

var (
  name  = flag.String("name", "", "some name")
  float = flag.Float64("float", 1.0, "some float")
  force = flag.Bool("force", false, "force something")
)

func main(){
  flag.Parse()

  fmt.Println("Name:", *name)
  fmt.Println("Float:", *float)
  fmt.Println("Force:", *force)

  // ref: https://golang.org/pkg/flag/
}
