package main

import (
	"github.com/hhhhp52/go_design_pattern/factory"
	"github.com/hhhhp52/go_design_pattern/strategy"
)

func main() {

	// design pattern : factory
	// GetProduct{"product", "state"},product can choose {potatofries or sweetpotatofries}
	product := factory.GetProduct("sweetpotatofries", "")
	factory.DescribeProduct(product)

	// design pattern : strategy
	opeartion := strategy.GetCaculator("DIVIDE", 1, 2)
	strategy.Caculator(opeartion)

}
