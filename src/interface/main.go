package main

import (
	"fmt"
)

type BinaryOperator interface {
	Operate(int, int) int
}

type Appender interface {
	Append(int)
}

// implementation
type Calculator int

func (c Calculator) Operate(a, b int) int {
	return a + b
}

func (c *Calculator) Append(n int) {
  *c += Calculator(n)
}

func main() {
	var c Calculator = Calculator(1)
	fmt.Println(c)

  var operator BinaryOperator = c // cast to interface
  n := operator.Operate(11, 22)
	fmt.Println(n)

  var appender Appender = &c
  appender.Append(33)
  fmt.Println(c)
}
