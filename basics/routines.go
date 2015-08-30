package main

import "fmt"
import "sync"
import "time"

func print_dot(){
  fmt.Print(".")
  time.Sleep(300 * time.Millisecond)
}

func print_comma(){
  fmt.Print(",")
  time.Sleep(300 * time.Millisecond)
}

// ================================

func run_serial(){
  size := 10
  for i:=0; i < size; i++{
    print_dot()
  }

  for i:=0; i < size; i++{
    print_comma()
  }
  fmt.Println("Done!")
}

func run_parallel(){
  var wg sync.WaitGroup
  size := 10

  wg.Add(size)
  go func(){
    for i:=0; i < size; i++{
      print_dot()
      //defer wg.Done()
      wg.Done()
    }
  }()

  wg.Add(size)
  go func(){
    for i:=0; i < size; i++{
      print_comma()
      //defer wg.Done()
      wg.Done()
    }
  }()

  wg.Wait()
  fmt.Println("Done!")
}


func main(){
  fmt.Println("Run serial:")
  run_serial()

  fmt.Print("\n")

  fmt.Println("Run in parallel:")
  run_parallel()


  // ref
  // https://gobyexample.com/goroutines
  // http://nathanleclaire.com/blog/2014/02/21/how-to-wait-for-all-goroutines-to-finish-executing-before-continuing-part-two-fixing-my-ooops/
}
