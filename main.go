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
			fmt.Println("Please Enter the value for a")
			var a float64
			fmt.Scanln(&a)
			fmt.Println("Please Enter the Value for b")
			var b float64
			fmt.Scanln(&b)
			cal := operation{a, b}
			result, err := cal.Add()
			if err != nil {
				fmt.Println(err)
				continue
			}
			resultMap["Addition"] = result
			fmt.Println("The result is:", result)

		case 2:
			fmt.Println("You have Selected Subtraction:")
			fmt.Println("Please Enter the value for a")
			var a float64
			fmt.Scanln(&a)
			fmt.Println("Please Enter the value for b")
			var b float64
			fmt.Scanln(&b)
			cal := operation{a, b}
			result, err := cal.Subtract()
			if err != nil {
				fmt.Println(err)
				continue
			}
			resultMap["Subtraction"] = result
			fmt.Println("The result is:", result)

		case 3:
			fmt.Println("You have Selected Multiplication:")
			fmt.Println("Please Enter the value for a")
			var a float64
			fmt.Scanln(&a)
			fmt.Println("Please Enter the value for b")
			var b float64
			fmt.Scanln(&b)
			cal := operation{a, b}
			result, err := cal.Multiply()
			if err != nil {
				fmt.Println(err)
				continue
			}
			resultMap["Multiplication"] = result
			fmt.Println("The result is:", result)

		case 4:
			fmt.Println("You have Selected Division:")
			fmt.Println("Please Enter the value for a")
			var a float64
			fmt.Scanln(&a)
			fmt.Println("Please Enter the value for b")
			var b float64
			fmt.Scanln(&b)
			cal := operation{a, b}
			result, err := cal.Divide()
			if err != nil {
				fmt.Println(err)
				continue
			}
			resultMap["Division"] = result
			fmt.Println("The result is:", result)

		case 5:
			fmt.Println("Previous results:")
			for i, value := range resultMap {
				if v, ok := value.(float64); ok {
					if v == float64(int(v)) {
						fmt.Printf("%s: %d\n", i, int(v))
					} else {
						fmt.Printf("%s: %f\n", i, v)
					}
				} else {
					fmt.Println("Error: unexpected type")
				}
			}

		case 0:
			fmt.Println("You have selected 0 for exit")
			return

		default:
			fmt.Println("Invalid choice. Please choose a valid number.")
		}
	}
}
