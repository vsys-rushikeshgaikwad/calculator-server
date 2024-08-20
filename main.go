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

type operation struct {
	a, b float64
}

func (c *operation) Add() (float64, error) {
	if c.b == 0 || c.a == 0 {
		return 0, errors.New("Cannot add by zero")
	}
	return c.a + c.b, nil
}

func (c *operation) Subtract() (float64, error) {
	if c.b == 0 || c.a == 0 {
		return 0, errors.New("Cannot divide by by zero")
	}
	return c.a - c.b, nil
}

func (c *operation) Multiply() (float64, error) {
	if c.b == 0 || c.a == 0 {
		return 0, errors.New("Cannot Multiply by zero")
	}
	return c.a * c.b, nil
}

func (c *operation) Divide() (float64, error) {
	if c.b == 0 || c.a == 0 {
		return 0, errors.New("division by zero")
	}
	return c.a / c.b, nil
}

func main() {
	for {
		resultMap := make(map[string]interface{})
		var result float64
		var err error
		fmt.Println("Enter a number between 1 to 5")
		fmt.Println("Enter 1 for Addition")
		fmt.Println("Enter 2 for Subtraction")
		fmt.Println("Enter 3 for Multiplication")
		fmt.Println("Enter 4 for Division")
		fmt.Println("Enter 5 for Results")
		fmt.Println("Enter 0 for Exit")
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
			fmt.Println("Enter 5 to see the result:")
			var c int
			fmt.Scanln(&c)
			cal := operation{a, b}
			if c == 5 {
				result, err = cal.Add()
				fmt.Println("The result is:", result)
			} else {
				fmt.Println("Invalid choice")
			}

		case 2:
			fmt.Println("You have Selected 2:")
			fmt.Println("Please Enter the value for a")
			var a float64
			fmt.Scanln(&a)
			fmt.Println("Please Enter the value for b")
			var b float64
			fmt.Scanln(&b)
			fmt.Println("Enter 5 to see the result:")
			var c int
			fmt.Scanln(&c)
			cal := operation{a, b}
			if c == 5 {
				result, err = cal.Subtract()
				fmt.Println("The result is:", result)
			} else {
				fmt.Println("Invalid choice")
			}
		case 3:
			fmt.Println("You have Selected 3:")
			fmt.Println("Please Enter the value for a")
			var a float64
			fmt.Scanln(&a)
			fmt.Println("Please Enter the value for b")
			var b float64
			fmt.Scanln(&b)
			fmt.Println("Enter 5 to see the result:")
			var c int
			fmt.Scanln(&c)
			cal := operation{a, b}
			if c == 5 {
				result, err = cal.Multiply()
				fmt.Println("The result is:", result)
			} else {
				fmt.Println("Invalid choice")
			}

		case 4:
			fmt.Println("You have Selected 4:")
			fmt.Println("Please Enter the value for a")
			var a float64
			fmt.Scanln(&a)
			fmt.Println("Please Enter the value for b")
			var b float64
			fmt.Scanln(&b)
			fmt.Println("Enter 5 to see the result:")
			var c int
			fmt.Scanln(&c)
			cal := operation{a, b}
			if c == 5 {
				result, err = cal.Divide()
				fmt.Println("The result is:", result)
			} else {
				fmt.Println("Invalid choice")
			}

		case 5:
			fmt.Println("Previous results:")
			//resultMap["result"] = result
			fmt.Println("Result in map:", resultMap)

		case 0:
			fmt.Println("You have Selected 0 for exit")
			return

		default:
			fmt.Println("Invalid Number")
			fmt.Println("Please choose a valid number")
			main()
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		//resultMap := make(map[string]interface{})
		resultMap["result"] = result
		fmt.Println("Result in map:", resultMap)
	}
}
