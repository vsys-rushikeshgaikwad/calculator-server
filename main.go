package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
)

var resultMap = make(map[string]interface{})

type Calculator interface {
	Add() (float64, error)
	Subtract() (float64, error)
	Multiply() (float64, error)
	Divide() (float64, error)
}

type value struct {
	A string
	B string
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

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path == "/results" {
			page := &ResultsPage{}
			w.Write([]byte(page.Build()))
		} else {
			page := &CalculatorPage{}
			w.Write([]byte(page.Build()))
		}
	} else if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var data map[string]string
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println(err)
			return
		}

		valueA, err := strconv.ParseFloat(data["a"], 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		valueB, err := strconv.ParseFloat(data["b"], 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		operationType := data["op"]
		fmt.Printf("Operation: %s\n", operationType)
		fmt.Printf("Value A: %v\n", valueA)
		fmt.Printf("Value B: %v\n", valueB)

		var calc Calculator
		op := operation{valueA, valueB}
		switch operationType {
		case "add":
			calc = &op
			result, err := calc.Add()
			if err != nil {
				fmt.Println(err)
				return
			}
			resultMap["Addition"] = result
			fmt.Printf("Result: %v\n", result)
			w.Write([]byte(fmt.Sprintf("%v", result)))
		case "subtract":
			calc = &op
			result, err := calc.Subtract()
			if err != nil {
				fmt.Println(err)
				return
			}
			resultMap["Subtraction"] = result
			fmt.Printf("Result: %v\n", result)
			w.Write([]byte(fmt.Sprintf("%v", result)))
		case "multiply":
			calc = &op
			result, err := calc.Multiply()
			if err != nil {
				fmt.Println(err)
				return
			}
			resultMap["Multiplication"] = result
			fmt.Printf("Result: %v\n", result)
			w.Write([]byte(fmt.Sprintf("%v", result)))

		case "divide":
			calc = &op
			result, err := calc.Divide()
			if err != nil {
				fmt.Println(err)
				return
			}
			resultMap["Division"] = result
			fmt.Printf("Result: %v\n", result)
			w.Write([]byte(fmt.Sprintf("%v", result)))

		default:
			fmt.Println("Invalid operation")
		}

	}
}

type CalculatorPage struct {
	Html string
}

func (c *CalculatorPage) Build() string {
	htmlCode := `<!DOCTYPE html>
<html>
<head>
	<title>Calculator</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			text-align: center;
		}
		.container {
			width: 50%;
			margin: 40px auto;
			padding: 20px;
			border: 1px solid #ccc;
			border-radius: 10px;
			box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
		}
		.dropdown {
			width: 94%;
			padding: 10px;
			border: none;
			border-radius: 5px;
			background-color: #f9f9f9;
            margin-bottom: 18px;
		}
		.input-field {
			width: 92%;
			padding: 10px;
			margin-bottom: 20px;
			border: 1px solid #ccc;
		}
		.button {
			background-color: #4CAF50;
			color: #fff;
			padding: 10px 20px;
			border: none;
			border-radius: 5px;
			cursor: pointer;
		}
		.button:hover {
			background-color: #3e8e41;
		}
		.result {
			margin-top: 20px;
			font-size: 24px;
			font-weight: bold;
			color: #666;
		}
		#result-field{
			margin-top: -7rem;
		}
		
	</style>
</head>
<body>
	<div class="container">
		<h2>Calculator</h2>
		<select class="dropdown" id="operation">
			<option value="">Select Operation</option>
			<option value="add">Addition</option>
			<option value="subtract">Subtraction</option>
			<option value="multiply">Multiplication</option>
			<option value="divide">Division</option>
		</select>
		<input type="number" class="input-field" id="a" placeholder="Enter value for A">
		<input type="number" class="input-field" id="b" placeholder="Enter value for B">
		<div class="result" id="result"></div>
		<input type="text" class="input-field" id="result-field" placeholder="Result">
		<button class="button" id="submit" name="submit">Submit</button>
		<button class="button" id="result-btn">Result</button>
		<button class="button" id="clear-btn">Clear</button>
	</div>

	<script>
  	document.addEventListener("DOMContentLoaded", () => {
    const SubmitBtn = document.getElementById("submit");
	const clearBtn = document.getElementById("clear-btn");
    const resultBtn = document.getElementById("result-btn");

    SubmitBtn.addEventListener("click", (event) => {
      event.preventDefault();

      const valueA = document.getElementById("a").value;
      const valueB = document.getElementById("b").value;
      const operation = document.getElementById("operation").value;
      const url = "http://localhost:8086/";
      fetch(url, {
        method: 'POST',
        body: JSON.stringify({ a: valueA, b: valueB, op: operation }),
        headers: {
          'Content-Type': 'application/json'
        }
      })
      .then(response => response.text())
      .then(result => {
        document.getElementById("result-field").value = result;
      });
    });

    resultBtn.addEventListener("click", (event) => {
      event.preventDefault();
      window.location.href = "http://localhost:8086/results";
    });

    clearBtn.addEventListener("click", (event) => {
      event.preventDefault();
      document.getElementById("a").value = "";
      document.getElementById("b").value = "";
      document.getElementById("operation").value = "";
      document.getElementById("result-field").value = "";
    });
  });
</script>
</body>
</html>`
	return htmlCode
}

type ResultsPage struct {
	Html string
}

func (r *ResultsPage) Build() string {
	var resultsHtml strings.Builder
	resultsHtml.WriteString(`<html><body><h2>Operation Results</h2><ul>`)
	for operation, result := range resultMap {
		resultsHtml.WriteString(fmt.Sprintf("%s: %v", operation, result))
	}
	resultsHtml.WriteString(`<form action="/" method="get">
	<button class="button" id="back-btn">Back to Calculator</button>
</form></body></html>`)
	return resultsHtml.String()
}

func main() {

	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8086", nil)

	ln, err := net.Listen("tcp", ":8086")
	if err != nil {
		if strings.Contains(err.Error(), "address already in use") {
			return
		}
		fmt.Println(err)
		return
	}
	defer ln.Close()

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
