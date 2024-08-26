package main

import (
	"errors"
	"fmt"
)

type Calculator interface {
	Add() (float64, error)
	Subtract() (float64, error)
	Multiply() (float64, error)
	Divide() (float64, error)
}

type operation struct {
	a, b float64
}

func (c *operation) Add() (float64, error) {
	return c.a + c.b, nil
}

func (c *operation) Subtract() (float64, error) {
	return c.a - c.b, nil
}

func (c *operation) Multiply() (float64, error) {
	return c.a * c.b, nil
}

func (c *operation) Divide() (float64, error) {
	if c.b == 0 {
		return 0, errors.New("division by zero")
	}
	return c.a / c.b, nil
}

func userInput() (float64, float64) {
	fmt.Println("Enter the value of a")
	var a float64
	fmt.Scanln(&a)
	fmt.Println("Enter the value of b")
	var b float64
	fmt.Scanln(&b)
	return a, b
}

func main() {
	// Map to store results of operations
	resultMap := make(map[string]interface{})

	for {
		fmt.Println("Enter a number between 1 to 5")
		fmt.Println("Enter 1 for Addition")
		fmt.Println("Enter 2 for Subtraction")
		fmt.Println("Enter 3 for Multiplication")
		fmt.Println("Enter 4 for Division")
		fmt.Println("Enter 5 to View Results")
		fmt.Println("Enter 0 to Exit")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("You have Selected Addition:")
			a, b := userInput()
			cal := operation{a, b}
			result, err := cal.Add()
			if err != nil {
				fmt.Println(err)
				continue
			}
			resultMap["Addition"] = result

		case 2:
			fmt.Println("You have Selected Subtraction:")
			a, b := userInput()
			cal := operation{a, b}
			result, err := cal.Subtract()
			if err != nil {
				fmt.Println(err)
				continue
			}
			resultMap["Subtraction"] = result

		case 3:
			fmt.Println("You have Selected Multiplication:")
			a, b := userInput()
			cal := operation{a, b}
			result, err := cal.Multiply()
			if err != nil {
				fmt.Println(err)
				continue
			}
			resultMap["Multiplication"] = result

		case 4:
			fmt.Println("You have Selected Division:")
			a, b := userInput()
			cal := operation{a, b}
			result, err := cal.Divide()
			if err != nil {
				fmt.Println(err)
				continue
			}
			resultMap["Division"] = result

		case 5:
			fmt.Println("Previous results:")
			for i, value := range resultMap {
				fmt.Printf("%s: %v\n", i, value)
			}
		case 0:
			fmt.Println("You have selected 0 for exit")
			return

		default:
			fmt.Println("Invalid choice. Please choose a valid number.")
		}
	}
}
