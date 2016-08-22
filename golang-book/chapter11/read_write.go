package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
  //Here starts the read method
  //____________________________________________________________

    bs, err := ioutil.ReadFile("textos/test.txt")

    if err != nil {
    return
    }

    str := string(bs)
    fmt.Println(str)
    //Here ends the read method
    //____________________________________________________________

    //Here starts the create method
    //____________________________________________________________

    file, err := os.Create("test.txt")
    if err != nil {
    // handle the error here
    return
    }
    defer file.Close()
    file.WriteString("This is a test to how to create a file")
//Here ends the create method
//____________________________________________________________

    bs1, err1 := ioutil.ReadFile("test.txt")

    if err1 != nil {
    return
    }

    str2 := string(bs1)
    fmt.Println(str2)

}
