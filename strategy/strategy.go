package strategy

import (
	"fmt"
	"strings"
)

//Strategy have four opeartions
type Strategy interface {
	Opeartions()
}

type numArray struct {
	num1          int
	num2          int
	result        int
	opeartionType string
}

func (numarray numArray) Opeartions() {
	var opeartionType = numarray.opeartionType
	if strings.EqualFold(opeartionType, "ADD") == true {
		fmt.Printf("%d + %d = %d\n", numarray.num1, numarray.num2, result)
	} else if strings.EqualFold(opeartionType, "DIVIDE") == true {
		fmt.Printf("%d - %d = %d\n", numarray.num1, numarray.num2, result)
	} else if strings.EqualFold(opeartionType, "MUTYPLY") == true {
		fmt.Printf("%d / %d = %d\n", numarray.num1, numarray.num2, result)
	} else {
		fmt.Printf("%d * %d = %d\n", numarray.num1, numarray.num2, result)
	}
}

var result int

// AddCaculator is a + b
func AddCaculator(num1, num2 int) {
	result = num1 + num2
}

// DivideCaculator is a - b
func DivideCaculator(num1, num2 int) {
	result = num1 - num2
}

// MinusCaculator is a / b
func MinusCaculator(num1, num2 int) {
	result = num1 / num2
}

// MultyplyCaculator is a * b
func MultyplyCaculator(num1, num2 int) {
	result = num1 * num2
}

// Caculator to use the opeartion and ruturn the ans
func Caculator(strategy Strategy) {
	if strategy != nil {
		strategy.Opeartions()
	}
}

// GetCaculator to choose the opeartion
func GetCaculator(opeartionType string, num1, num2 int) Strategy {
	switch opeartionType {
	case "ADD":
		AddCaculator(num1, num2)
		return &numArray{num1: num1, num2: num2, result: result, opeartionType: opeartionType}
	case "DIVIDE":
		DivideCaculator(num1, num2)
		return &numArray{num1: num1, num2: num2, result: result, opeartionType: opeartionType}
	case "MUTYPLY":
		MultyplyCaculator(num1, num2)
		return &numArray{num1: num1, num2: num2, result: result, opeartionType: opeartionType}
	case "MINUS":
		MinusCaculator(num1, num2)
		return &numArray{num1: num1, num2: num2, result: result, opeartionType: opeartionType}
	default:
		fmt.Printf("No this opeartion")
		return nil
	}
}
