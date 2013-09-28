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
  eval, eval_env := example1.Evaluate(env)
  fmt.Println("Evaluated:", eval.Inspect(), eval_env)
  fmt.Printf("Ruby\n----\n%s\n----\n", example1.ToRuby())

  fmt.Println("\nBooleans")
  example2 := simple.LessThan{simple.Number{5}, simple.Add{simple.Number{2},simple.Number{2}}}
  machine = simple.Machine{example2,env}
  machine.Run()
  eval, eval_env = example2.Evaluate(env)
  fmt.Println("Evaluated:", eval.Inspect(), eval_env)
  fmt.Printf("Ruby\n----\n%s\n----\n", example2.ToRuby())

  fmt.Println("\nVariables")
  env = simple.Env {"x":simple.Number{3}, "y":simple.Number{4}}
  example3 := simple.Add{simple.Variable{"x"},simple.Variable{"y"}}
  machine = simple.Machine{example3,env}
  machine.Run()
  env = simple.Env {"x":simple.Number{3}, "y":simple.Number{4}}
  eval, eval_env = example3.Evaluate(env)
  fmt.Println("Evaluated:", eval.Inspect(), eval_env)
  fmt.Printf("Ruby\n----\n%s\n----\n", example3.ToRuby())

  fmt.Println("\nAssignment")
  env = simple.Env {"x":simple.Number{2}}
  example4 := simple.Assign{"x",simple.Add{simple.Variable{"x"},simple.Number{1}}}
  machine = simple.Machine{example4,env}
  machine.Run()
  env = simple.Env {"x":simple.Number{2}}
  eval, eval_env = example4.Evaluate(env)
  fmt.Println("Evaluated:", eval.Inspect(), eval_env)
  fmt.Printf("Ruby\n----\n%s\n----\n", example4.ToRuby())

  fmt.Println("\nConditional")
  env = simple.Env {"x":simple.Boolean{true}}
  example5 := simple.If{simple.Variable{"x"},simple.Assign{"y",simple.Number{1}},simple.Assign{"y",simple.Number{2}}}
  machine = simple.Machine{example5,env}
  machine.Run()
  env = simple.Env {"x":simple.Boolean{true}}
  eval, eval_env = example5.Evaluate(env)
  fmt.Println("Evaluated:", eval.Inspect(), eval_env)

  env = simple.Env {"x":simple.Boolean{false}}
  machine = simple.Machine{example5,env}
  machine.Run()
  env = simple.Env {"x":simple.Boolean{false}}
  eval, eval_env = example5.Evaluate(env)
  fmt.Println("Evaluated:", eval.Inspect(), eval_env)
  fmt.Printf("Ruby\n----\n%s\n----\n", example5.ToRuby())

  fmt.Println("\nSequences")
  env = simple.Env{}
  example6 := simple.Sequence{simple.Assign{"x",simple.Add{simple.Number{1},simple.Number{1}}},simple.Assign{"y",simple.Add{simple.Variable{"x"},simple.Number{3}}}}
  machine = simple.Machine{example6,env}
  machine.Run()
  env = simple.Env{}
  eval, eval_env = example6.Evaluate(env)
  fmt.Println("Evaluated:", eval.Inspect(), eval_env)
  fmt.Printf("Ruby\n----\n%s\n----\n", example6.ToRuby())

  fmt.Println("\nWhile Loop")
  env = simple.Env{"x":simple.Number{1}}
  example7 := simple.While{simple.LessThan{simple.Variable{"x"},simple.Number{5}},simple.Assign{"x",simple.Multiply{simple.Variable{"x"},simple.Number{3}}}}
  machine = simple.Machine{example7,env}
  machine.Run()
  env = simple.Env{"x":simple.Number{1}}
  eval, eval_env = example7.Evaluate(env)
  fmt.Println("Evaluated:", eval.Inspect(), eval_env)
  fmt.Printf("Ruby\n----\n%s\n----\n", example7.ToRuby())
}
