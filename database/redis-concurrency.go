package main

import "fmt"
import "time"
import "gopkg.in/redis.v3"

var concurrency int = 4
var hit, miss int

func setup_redis() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
        PoolSize: concurrency,
    })
    return client
}

func get(key string, client *redis.Client) string {
  value, err := client.Get(key).Result()

  if err == nil {
    hit += 1
    //fmt.Println(key, "=", value)
  } else {
    miss += 1
    fmt.Println("Miss =>", miss, "/ Hit =>", hit)
    value = update_cache(key, client)
    //panic(err)
  }
  return value
}

func update_cache(key string, client *redis.Client) string {
  value := "doesnt_really_matter"
  ttl := 10 * time.Second
  err := client.Set(key, value, ttl).Err()

  if err != nil {
    fmt.Println("Failed to set:", err)
    panic(err)
  }
  return value
}

func benchmark(){
  last := 0
  for {
    total := hit + miss
    ops := total - last
    last = total
    fmt.Println("Operations per sec:", ops)
    time.Sleep(1 * time.Second)
  }
}

func main(){
  sem := make(chan bool, concurrency)
  client := setup_redis()
  key := "test-key"

  go benchmark()

  for i := 0; i < 5000000; i++ {
    //get(key, client)  // serial

    sem <- true
    go func (){
      get(key, client)
      <-sem
    }()
  }

  fmt.Println("Done!")
}
