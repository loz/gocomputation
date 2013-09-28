package simple

import (
  "fmt"
)

type Node interface {}

type Number struct {
  Value int
}

func (self Number) String() string {
  return fmt.Sprintf("%v", self.Value)
}

func (self Number) Inspect() string {
  return fmt.Sprintf("≪%v≫", self)
}


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

