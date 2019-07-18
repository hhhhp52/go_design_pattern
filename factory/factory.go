package factory

import (
	"fmt"
	"log"
)

// Product interface
type Product interface {
	Describe()
}

// PotatoFries struct
type PotatoFries struct {
	state string
}

// SweetPotatoFries struct
type SweetPotatoFries struct {
	state string
}

// Describe potato fries
func (potatofries PotatoFries) Describe() {
	if potatofries.state != "" {
		fmt.Printf("This is %s potato fries\n", potatofries.state)
	} else {
		fmt.Printf("This is original potato fries\n")
	}
}

// Describe sweet potato fries
func (sweetpotatofries SweetPotatoFries) Describe() {
	if sweetpotatofries.state != "" {
		fmt.Printf("This is %s sweet potato fries\n", sweetpotatofries.state)
	} else {
		fmt.Printf("This is original sweet potato fries\n")
	}
}

// DescribeProduct to describe product
func DescribeProduct(product Product) {
	if product != nil {
		product.Describe()
	}
}

// GetProduct to get the product
func GetProduct(product string, state string) Product {
	switch product {
	case "potatofries":
		return &PotatoFries{state: state}
	case "sweetpotatofries":
		return &SweetPotatoFries{state: state}
	default:
		log.Println("Product not found")
		return nil
	}
}
