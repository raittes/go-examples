package main

import "fmt"
import "time"
import "math/rand"

func random_array() {
  a := [10]int{}
  for i := 0; i < len(a); i++ {
    a[i] = rand.Intn(100)
    fmt.Print(a[i]," ")
  }
  fmt.Println()
}

func main(){
  fmt.Print("Not so random: ")
  random_array()

  fmt.Print("Really random: ")
  rand.Seed( time.Now().UTC().UnixNano() )
  random_array()
}
