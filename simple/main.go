package main

import (
  "github.com/loz/gocomputation/simple/simple"
  "fmt"
  )

func main() {
  var env simple.Env;

  example1 := simple.Add{(simple.Multiply{(simple.Number{1}), (simple.Number{2})}), (simple.Multiply{(simple.Number{3}), (simple.Number{4})}) }
  fmt.Println(example1.Inspect())

  fmt.Println("\nReduced on a Machine")
  machine := simple.Machine{example1,env}
  machine.Run()

  fmt.Println("\nBooleans")
  example2 := simple.LessThan{simple.Number{5}, simple.Add{simple.Number{2},simple.Number{2}}}
  machine = simple.Machine{example2,env}
  machine.Run()

  fmt.Println("\nVariables")
  env = simple.Env {"x":simple.Number{3}, "y":simple.Number{4}}
  example3 := simple.Add{simple.Variable{"x"},simple.Variable{"y"}}
  machine = simple.Machine{example3,env}
  machine.Run()

  fmt.Println("\nAssignment")
  env = simple.Env {"x":simple.Number{2}}
  example4 := simple.Assign{"x",simple.Add{simple.Variable{"x"},simple.Number{1}}}
  machine = simple.Machine{example4,env}
  machine.Run()

  fmt.Println("\nConditional")
  env = simple.Env {"x":simple.Boolean{true}}
  example5 := simple.If{simple.Variable{"x"},simple.Assign{"y",simple.Number{1}},simple.Assign{"y",simple.Number{2}}}
  machine = simple.Machine{example5,env}
  machine.Run()
  env = simple.Env {"x":simple.Boolean{false}}
  machine = simple.Machine{example5,env}
  machine.Run()

  fmt.Println("\nSequences")
  env = simple.Env{}
  example6 := simple.Sequence{simple.Assign{"x",simple.Add{simple.Number{1},simple.Number{1}}},simple.Assign{"y",simple.Add{simple.Variable{"x"},simple.Number{3}}}}
  machine = simple.Machine{example6,env}
  machine.Run()

  fmt.Println("\nWhile Loop")
  env = simple.Env{"x":simple.Number{1}}
  example7 := simple.While{simple.LessThan{simple.Variable{"x"},simple.Number{5}},simple.Assign{"x",simple.Multiply{simple.Variable{"x"},simple.Number{3}}}}
  machine = simple.Machine{example7,env}
  machine.Run()
}
