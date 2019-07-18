package main

import (
	"github.com/hhhhp52/go_design_pattern/factory"
)

func main() {

	// design pattern : factory
	// GetProduct{"product", "state"},product can choose {potatofries or sweetpotatofries}
	product := factory.GetProduct("sweetpotatofries", "")
	factory.DescribeProduct(product)
}
