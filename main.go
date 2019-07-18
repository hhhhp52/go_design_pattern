package main

import (
	"github.com/hhhhp52/go_design_pattern/factory"
)

func main() {

	// design pattern : factory
	product := factory.GetProduct("sweetpotatofries", "salt")
	factory.DescribeProduct(product)
}
