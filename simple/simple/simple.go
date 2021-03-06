package simple

import (
  "fmt"
)

type Env map[string]Node

type Node interface {
  Reduceable() bool
  Reduce(Env) (Node, Env)
  Inspect() string
  Evaluate(Env) (Node, Env)
  ToRuby() string
}

/* Machine */
type Machine struct {
  Expression Node
  Environment Env
}

func (self *Machine) Step() {
  self.Expression, self.Environment = self.Expression.Reduce(self.Environment)
}

func (self *Machine) Run() {
  for self.Expression.Reduceable() {
    fmt.Printf("%v, %v\n", self.Expression, self.Environment)
    self.Step()
  }
  fmt.Printf("%v, %v\n", self.Expression, self.Environment)
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

func (self Number) Reduce(e Env) (Node, Env) {
  return self, e
}

func (self Number) Evaluate(e Env) (Node, Env) {
  return self, e
}

func (self Number) ToRuby() string {
  return fmt.Sprintf("-> e { %v }", self.Value)
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

func (self Boolean) Reduce(e Env) (Node, Env) {
  return self, e
}

func (self Boolean) Evaluate(e Env) (Node, Env) {
  return self, e
}

func (self Boolean) ToRuby() string {
  return fmt.Sprintf("-> e { %v }", self.Value)
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

func (self Add) Reduce(e Env) (Node, Env) {
  if self.Left.Reduceable() {
    l, _ := self.Left.Reduce(e)
    return Add{l, self.Right}, e
  } else if self.Right.Reduceable() {
    r, _ := self.Right.Reduce(e)
    return Add{self.Left, r}, e
  } else {
    return Number{(self.Left.(Number).Value + self.Right.(Number).Value)}, e
  }
}

func (self Add) Evaluate(e Env) (Node, Env) {
  le, _ := self.Left.Evaluate(e)
  re, _ := self.Right.Evaluate(e)
  return Number{(le.(Number).Value + re.(Number).Value)}, e
}

func (self Add) ToRuby() string {
  return fmt.Sprintf("-> e { (%s).call(e) + (%s).call(e) }", self.Left.ToRuby(), self.Right.ToRuby())
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

func (self Multiply) Reduce(e Env) (Node, Env) {
  if self.Left.Reduceable() {
    l, _ := self.Left.Reduce(e)
    return Multiply{l, self.Right}, e
  } else if self.Right.Reduceable() {
    r, _ := self.Right.Reduce(e)
    return Multiply{self.Left, r}, e
  } else {
    return Number{(self.Left.(Number).Value * self.Right.(Number).Value)}, e
  }
}

func (self Multiply) Evaluate(e Env) (Node, Env) {
  le, _ := self.Left.Evaluate(e)
  re, _ := self.Right.Evaluate(e)
  return Number{(le.(Number).Value * re.(Number).Value)}, e
}

func (self Multiply) ToRuby() string {
  return fmt.Sprintf("-> e { (%s).call(e) * (%s).call(e) }", self.Left.ToRuby(), self.Right.ToRuby())
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

func (self LessThan) Reduce(e Env) (Node, Env) {
  if self.Left.Reduceable() {
    l, _ := self.Left.Reduce(e)
    return LessThan{l, self.Right}, e
  } else if self.Right.Reduceable() {
    r, _ := self.Right.Reduce(e)
    return LessThan{self.Left, r}, e
  } else {
    return Boolean{(self.Left.(Number).Value < self.Right.(Number).Value)}, e
  }
}

func (self LessThan) Evaluate(e Env) (Node, Env) {
  le, _ := self.Left.Evaluate(e)
  re, _ := self.Right.Evaluate(e)
  return Boolean{(le.(Number).Value < re.(Number).Value)}, e
}

func (self LessThan) ToRuby() string {
  return fmt.Sprintf("-> e { (%s).call(e) < (%s).call(e) }", self.Left.ToRuby(), self.Right.ToRuby())
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

func (self Variable) Reduce(environment Env) (Node, Env) {
  return environment[self.Name], environment
}

func (self Variable) Evaluate(e Env) (Node, Env) {
  return e[self.Name], e
}

func (self Variable) ToRuby() string {
  return fmt.Sprintf("-> e { e['%s'] }", self.Name)
}

/* Statements */
/******************/

/* DoNothing */
type DoNothing struct {}

func (self DoNothing) String() string {
  return "do-nothing"
}

func (self DoNothing) Inspect() string {
  return fmt.Sprintf("≪%v≫", self)
}

func (self DoNothing) Reduceable() bool {
  return false
}

func (self DoNothing) Reduce(e Env) (Node, Env) {
  return self, e
}

func (self DoNothing) Evaluate(e Env) (Node, Env) {
  return self, e
}

func (self DoNothing) ToRuby() string {
  return "-> e { e }"
}

/* Assign */
type Assign struct {
  Name string
  Expression Node
}

func (self Assign) String() string {
  return fmt.Sprintf("%v = %v", self.Name, self.Expression)
}

func (self Assign) Inspect() string {
  return fmt.Sprintf("≪%v≫", self)
}

func (self Assign) Reduceable() bool {
  return true
}

func (self Assign) Reduce(e Env) (Node, Env) {
  if self.Expression.Reduceable() {
    exp, _ := self.Expression.Reduce(e)
    return Assign{self.Name, exp}, e
  } else {
    e[self.Name] = self.Expression
    return DoNothing{}, e
  }
}

func (self Assign) Evaluate(e Env) (Node, Env) {
  e[self.Name], _ = self.Expression.Evaluate(e)
  return self, e
}

func (self Assign) ToRuby() string {
  return fmt.Sprintf("-> e { e.merge({'%s' => (%s).call(e)}) }", self.Name, self.Expression.ToRuby())
}

/* If */
type If struct {
  Condition Node
  Consequence Node
  Alternative Node
}

func (self If) String() string {
  return fmt.Sprintf("if (%v) { %v } else { %v }",
                      self.Condition, self.Consequence,
                      self.Alternative)
}

func (self If) Inspect() string {
  return fmt.Sprintf("≪%v≫", self)
}

func (self If) Reduceable() bool {
  return true
}

func (self If) Reduce(e Env) (Node, Env) {
  if self.Condition.Reduceable() {
    cond, _ := self.Condition.Reduce(e)
    return If{cond, self.Consequence, self.Alternative}, e
  } else {
    if self.Condition == (Boolean{true}) {
      return self.Consequence, e
    } else {
      return self.Alternative, e
    }
  }
}

func (self If) Evaluate(e Env) (Node, Env) {
  con, _ := self.Condition.Evaluate(e)
  if con == (Boolean{true}) {
    return self.Consequence.Evaluate(e)
  } else {
    return self.Alternative.Evaluate(e)
  }
}

func (self If) ToRuby() string {
  return fmt.Sprintf("-> e { if (%s).call(e) then (%s).call(e) else (%s).call(e) end }",
                          self.Condition.ToRuby(),
                          self.Consequence.ToRuby(),
                          self.Alternative.ToRuby())
}

/* Sequence */
type Sequence struct {
  First Node
  Second Node
}

func (self Sequence) String() string {
  return fmt.Sprintf("%v; %v", self.First, self.Second)
}

func (self Sequence) Inspect() string {
  return fmt.Sprintf("≪%v≫", self)
}

func (self Sequence) Reduceable() bool {
  return true
}

func (self Sequence) Reduce(e Env) (Node, Env) {
  if self.First == (DoNothing{}) {
    return self.Second, e
  } else {
    r_first, r_env := self.First.Reduce(e)
    return Sequence{r_first, self.Second}, r_env
  }
}

func (self Sequence) Evaluate(e Env) (Node, Env) {
  _, new_env := self.First.Evaluate(e)
  return self.Second.Evaluate(new_env)
}

func (self Sequence) ToRuby() string {
  return fmt.Sprintf("-> e { (%s).call((%s).call(e)) }",
                          self.Second.ToRuby(),
                          self.First.ToRuby())
}

/* While */
type While struct {
  Condition Node
  Body Node
}

func (self While) String() string {
  return fmt.Sprintf("while (%v) { %v }", self.Condition, self.Body)
}

func (self While) Inspect() string {
  return fmt.Sprintf("≪%v≫", self)
}

func (self While) Reduceable() bool {
  return true
}

func (self While) Reduce(e Env) (Node, Env) {
  return If{self.Condition,Sequence{self.Body,self},DoNothing{}}, e
}

func (self While) Evaluate(e Env) (Node, Env) {
  con, _ := self.Condition.Evaluate(e)
  if con == (Boolean{true}) {
    _, new_env := self.Body.Evaluate(e)
    return self.Evaluate(new_env)
  } else {
    return self, e
  }
}

func (self While) ToRuby() string {
  return fmt.Sprintf("-> e { while (%s).call(e); e = (%s).call(e); end; e }",
                          self.Condition.ToRuby(),
                          self.Body.ToRuby())
}
