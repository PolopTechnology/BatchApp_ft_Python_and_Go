package main

import (
     "fmt"
     "io/ioutil"
     "log"
)

func ReadFile(filename string) {
    fmt.Printf("\n\nReading a file in Go lang\n")
      
    data, err := ioutil.ReadFile(filename)
      
    if err != nil {
        log.Panicf("failed reading data from file: %s", err)
    }
    fmt.Printf("\nFile Name: %s", filename)
    fmt.Printf("\nSize: %d bytes", len(data))
    fmt.Printf("\nData: %s", data)
  
}

func main() {
     ReadFile("log.dat")
}





