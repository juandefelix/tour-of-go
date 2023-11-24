package main

import ( 
  "golang.org/x/tour/tree"
  "fmt"
)


// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  if t == nil { return }

  Walk(t.Left, ch)
  ch <- t.Value
  Walk(t.Right, ch)
}


func Same(t1, t2 *tree.Tree) bool {
  ch1 := make(chan int, 10)
  ch2 := make(chan int, 10)

  go Walk(t2, ch2)
  go Walk(t1, ch1)

  for i := 0; i < 10; i++ {
    if <-ch1 != <-ch2 { return false }
  }
  return true
}

func main() {
  areTheSame := Same(tree.New(2), tree.New(2))
  fmt.Println(areTheSame)

  areTheSame = Same(tree.New(1), tree.New(2))
  fmt.Println(areTheSame)
}

// Import local modules: https://stackoverflow.com/questions/35480623/how-to-import-local-packages-in-go
