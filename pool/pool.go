package main

import "fmt"

func main (){


  total:= 0.3
  flag := 1

  for total < 50 {

    total = total*2
    flag= flag+1
    fmt.Println( "En el día ",flag, "  la piscina va llenandose por ",total)
  }
}
