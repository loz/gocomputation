package simple

import (
  "fmt"
)

type Env map[string]Node

type Node interface {
  Reduceable() bool
  Reduce(Env) Node
  Inspect() string
}

/* Machine */
type Machine struct {
  Expression Node
  Environment Env
}

func (self *Machine) Step() {
  self.Expression = self.Expression.Reduce(self.Environment)
}

func (self *Machine) Run() {
  for self.Expression.Reduceable() {
    fmt.Println(self.Expression)
    self.Step()
  }
  fmt.Println(self.Expression)
}

/* Number */
type Number struct {
  Value int
}

func (self Number) String() string {
  return fmt.Sprintf("%v", self.Value)
}

func (self Number) Inspect() string {
  return fmt.Sprintf("≪%v≫", self)
}

func (self Number) Reduceable() bool {
  return false
}

func (self Number) Reduce(e Env) Node {
  return self
}

/* Boolean */
type Boolean struct {
  Value bool
}

func (self Boolean) String() string {
  return fmt.Sprintf("%v", self.Value)
}

func (self Boolean) Inspect() string {
  return fmt.Sprintf("≪%v≫", self)
}

func (self Boolean) Reduceable() bool {
  return false
}

func (self Boolean) Reduce(e Env) Node {
  return self
}

/* Add */
type Add struct {
  Left Node
  Right Node
}

func (self Add) String() string {
  return fmt.Sprintf("%v + %v", self.Left, self.Right)
}

func (self Add) Inspect() string {
  return fmt.Sprintf("≪%v≫", self)
}

func (self Add) Reduceable() bool {
  return true
}

func (self Add) Reduce(e Env) Node {
  if self.Left.Reduceable() {
    return Add{(self.Left.Reduce(e)), self.Right}
  } else if self.Right.Reduceable() {
    return Add{self.Left, self.Right.Reduce(e)}
  } else {
    return Number{(self.Left.(Number).Value + self.Right.(Number).Value)}
  }
}

/* Multiply */
type Multiply struct {
  Left Node
  Right Node
}

func (self Multiply) String() string {
  return fmt.Sprintf("%v * %v", self.Left, self.Right)
}

func (self Multiply) Inspect() string {
  return fmt.Sprintf("≪%v≫", self)
}

func (self Multiply) Reduceable() bool {
  return true
}

func (self Multiply) Reduce(e Env) Node {
  if self.Left.Reduceable() {
    return Multiply{(self.Left.Reduce(e)), self.Right}
  } else if self.Right.Reduceable() {
    return Multiply{self.Left, self.Right.Reduce(e)}
  } else {
    return Number{(self.Left.(Number).Value * self.Right.(Number).Value)}
  }
}

/* LessThan */
type LessThan struct {
  Left Node
  Right Node
}

func (self LessThan) String() string {
  return fmt.Sprintf("%v < %v", self.Left, self.Right)
}

func (self LessThan) Inspect() string {
  return fmt.Sprintf("≪%v≫", self)
}

func (self LessThan) Reduceable() bool {
  return true
}

func (self LessThan) Reduce(e Env) Node {
  if self.Left.Reduceable() {
    return LessThan{(self.Left.Reduce(e)), self.Right}
  } else if self.Right.Reduceable() {
    return LessThan{self.Left, self.Right.Reduce(e)}
  } else {
    return Boolean{(self.Left.(Number).Value < self.Right.(Number).Value)}
  }
}

/* Variable */
type Variable struct {
  Name string
}

func (self Variable) String() string {
  return self.Name
}

func (self Variable) Inspect() string {
  return fmt.Sprintf("≪%v≫", self)
}

func (self Variable) Reduceable() bool {
  return true
}

func (self Variable) Reduce(environment Env) Node {
  return environment[self.Name]
}
