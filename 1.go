// id 1
package main

import "fmt"

func main() {
  var result int
  for i := 0; i < 1000; i++ {
    if i % 3 == 0 || i % 5 == 0 {
      result += i
    }
  }

  fmt.Printf("%v\n", result)
}