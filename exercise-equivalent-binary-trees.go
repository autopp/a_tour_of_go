package main

import (
  "golang.org/x/tour/tree"
  "fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  var walkMain func (*tree.Tree, chan int)
  walkMain = func (t *tree.Tree, ch chan int)  {
    if t == nil {
      return
    }
    walkMain(t.Left, ch)
    ch <- t.Value
    walkMain(t.Right, ch)
  }

  walkMain(t, ch)
  close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  return true
}

func main() {
  ch := make(chan int)
  go Walk(tree.New(1), ch)
  for n := range(ch) {
    fmt.Println(n)
  }
}
