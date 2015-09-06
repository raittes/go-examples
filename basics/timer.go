package main

import "fmt"
import "time"
import "os"

func main() {
	start := time.Now()
	counter := 0
	tick := time.Tick(100 * time.Millisecond)
	tick2 := time.Tick(5 * time.Second)
	tick3 := time.Tick(10 * time.Second)
	timeout := time.After(time.Second * 30)

	fmt.Println("Started at ", start)
	for {
		select {
		case <-tick:
			fmt.Print(".")
			counter += 1

		case <-tick2:
			fmt.Println("")
			fmt.Println("Time elapsed: ", time.Since(start))

		case <-tick3:
			fmt.Println("\nCounter:", counter)

		case <-timeout:
			fmt.Println("Timeout")
			os.Exit(0)
		}
	}

}
