package simple

import (
  "fmt"
)

type Node interface {
  Reduceable() bool
  Reduce() Node
  Inspect() string
}

/* Machine */
type Machine struct {
  Expression Node
}

func (self *Machine) Step() {
  self.Expression = self.Expression.Reduce()
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

func (self Number) Reduce() Node {
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

func (self Boolean) Reduce() Node {
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

func (self Add) Reduce() Node {
  if self.Left.Reduceable() {
    return Add{(self.Left.Reduce()), self.Right}
  } else if self.Right.Reduceable() {
    return Add{self.Left, self.Right.Reduce()}
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

func (self Multiply) Reduce() Node {
  if self.Left.Reduceable() {
    return Multiply{(self.Left.Reduce()), self.Right}
  } else if self.Right.Reduceable() {
    return Multiply{self.Left, self.Right.Reduce()}
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

func (self LessThan) Reduce() Node {
  if self.Left.Reduceable() {
    return LessThan{(self.Left.Reduce()), self.Right}
  } else if self.Right.Reduceable() {
    return LessThan{self.Left, self.Right.Reduce()}
  } else {
    return Boolean{(self.Left.(Number).Value < self.Right.(Number).Value)}
  }
}
