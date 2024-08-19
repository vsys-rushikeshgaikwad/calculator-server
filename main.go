package main

import (
	"errors"
	"fmt"
)

type Calculator interface {
	Add() float64
	Subtract() float64
	Multiply() float64
	Divide() (float64, error)
}

type calculator struct {
	a, b float64
}

func (c *calculator) Add() float64 {
	return c.a + c.b
}

func (c *calculator) Subtract() float64 {
	return c.a - c.b
}

func (c *calculator) Multiply() float64 {
	return c.a * c.b
}

func (c *calculator) Divide() (float64, error) {
	if c.b == 0 {
		return 0, errors.New("division by zero")
	}
	return c.a / c.b, nil
}

func main() {
	var result float64
	var err error
	fmt.Println("Enter a number between 1 to 4")
	fmt.Println("Enter 1 for Addition")
	fmt.Println("Enter 2 for Subtraction")
	fmt.Println("Enter 3 for Multiplication")
	fmt.Println("Enter 4 for Division")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("You have Selected 1:")
		fmt.Println("Please Enter the value for a")
		var a float64
		fmt.Scanln(&a)
		fmt.Println("Please Enter the Value for b")
		var b float64
		fmt.Scanln(&b)
		cal := calculator{a, b}
		result = cal.Add()
		fmt.Println("The result is:", result)

	case 2:
		fmt.Println("You have Selected 2:")
		fmt.Println("Please Enter the value for a")
		var a float64
		fmt.Scanln(&a)
		fmt.Println("Please Enter the value for b")
		var b float64
		fmt.Scanln(&b)
		cal := calculator{a, b}
		result = cal.Subtract()
		fmt.Println("The result is:", result)

	case 3:
		fmt.Println("You have Selected 3:")
		fmt.Println("Please Enter the value for a")
		var a float64
		fmt.Scanln(&a)
		fmt.Println("Please Enter the value for b")
		var b float64
		fmt.Scanln(&b)
		cal := calculator{a, b}
		result = cal.Multiply()
		fmt.Println("The result is:", result)

	case 4:
		fmt.Println("You have Selected 4:")
		fmt.Println("Please Enter the value for a")
		var a float64
		fmt.Scanln(&a)
		fmt.Println("Please Enter the value for b")
		var b float64
		fmt.Scanln(&b)
		cal := calculator{a, b}
		result, err = cal.Divide()
		fmt.Println("The result is:", result)

	default:
		fmt.Println("Invalid Number")
		fmt.Println("Please choose a valid number")
		main()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println("Result:", result)

	resultMap := map[string]float64{
		"result": result,
	}

	fmt.Println("Result in map:", resultMap)
}
