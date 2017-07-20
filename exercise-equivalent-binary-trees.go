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
  ch1 := make(chan int)
  ch2 := make(chan int)

  go Walk(t1, ch1)
  go Walk(t2, ch2)

  for {
    n1, b1 := <-ch1
    n2, b2 := <-ch2

    if b1 != b2 {
      return false
    }
    if b1 {
      if n1 != n2 {
        return false
      }
    } else {
      return true
    }
  }
}

func main() {
  t1 := tree.New(10)
  t2 := tree.New(10)
  fmt.Println(Same(t1, t2))
}
