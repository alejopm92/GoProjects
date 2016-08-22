package main

import (
  "fmt"
  "time"
)

//Using channels to store ping in channel c

func pinger(c chan string) {
  for i := 0; ; i++ {
  c <- "ping"
  }
}

//Using channels to store pong in channel c

func ponger(c chan string) {
  for i := 0; ; i++ {
    c <- "pong"
  }
}

//Using print the info of the channels

func printer(c chan string) {
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}

func main() {
  var c chan string = make(chan string, 2)
  var c2 chan string = make(chan string, 1)
  go pinger(c)
  go ponger(c2)
  go printer(c)
  go printer(c2)
  var input string
  fmt.Scanln(&input)
}
