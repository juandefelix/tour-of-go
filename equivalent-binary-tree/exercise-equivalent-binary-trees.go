epackage main

/*
STEPS FOLLOWED TO RUN THIS EXERCISE WITH REMOTE DEPENDENCIES:

1. Initially this file was in the parent directory
2. Created a folder exercise-equivalent-binary-tree. Change directory to it
3. Initialize a go module: `go mod init example/exercixe-binary-tree`
4. Fetch the remote libraries: `go mod tidy`
5. Now you should be able to run the program: `go run exercse-equivalent-binary-trees.go`

*/
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

