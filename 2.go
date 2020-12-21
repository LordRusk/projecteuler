// id 2
package main

import "fmt"

func main() {
  fibNums := []int{1, 2}
  result := 2
  
  for i := 2; result < 4000000; i++ {
    fibNums = append(fibNums, fibNums[i-1] + fibNums[i-2])
    if fibNums[i] % 2 == 0 {
      result += fibNums[i]
    }
  }

  fmt.Printf("%v\n", result)
}