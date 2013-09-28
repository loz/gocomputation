package main

import (
  "github.com/loz/gocomputation/simple/simple"
  "fmt"
  )

func main() {
  example1 := simple.Add{(simple.Multiply{(simple.Number{1}), (simple.Number{2})}), (simple.Multiply{(simple.Number{3}), (simple.Number{4})}) }
  fmt.Println(example1.Inspect())

  fmt.Println("Reduced on a Machine")
  machine := simple.Machine{example1}
  machine.Run()
}
