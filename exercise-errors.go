package main

import (
  "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}
 
func Sqrt(x float64) (float64, error) {
  if (x < 0) {
    return 0, ErrNegativeSqrt(x)
  } else {
    z := float64(1)
    for iteration := 0; iteration < 10; iteration ++ {
      z -= (z*z - x) / (2*z)
    }
    return z, nil
  }
}

func main() {
  value, error := Sqrt(-2)
  
  if error != nil {
    fmt.Println("Error:", error)    
  } else {
    fmt.Println("Value: ", value)
  }
}
