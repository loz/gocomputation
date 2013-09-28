package main

import (
  "github.com/loz/gocomputation/simple/simple"
  "fmt"
  )

func main() {
  example1 := simple.Add{(simple.Multiply{(simple.Number{1}), (simple.Number{2})}), (simple.Multiply{(simple.Number{3}), (simple.Number{4})}) }
  fmt.Println(example1.Inspect())

  var next simple.Node;
  next = example1
  fmt.Println(next.Reduceable())
  next = next.Reduce()
  fmt.Println(next.Inspect())
  fmt.Println(next.Reduceable())
  next = next.Reduce()
  fmt.Println(next.Inspect())
  fmt.Println(next.Reduceable())
  next = next.Reduce()
  fmt.Println(next.Inspect())
  fmt.Println(next.Reduceable())
}
