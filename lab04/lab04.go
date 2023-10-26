package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// TODO: Create a struct to hold the data sent to the template

type Answer struct {
	Expression string
	Result     string
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func Calculator(w http.ResponseWriter, r *http.Request) {
	// TODO: Finish this function
	op, num1, num2 := r.URL.Query().Get("op"), r.URL.Query().Get("num1"), r.URL.Query().Get("num2")

	number1, err := strconv.Atoi(num1)

	if err != nil {
		http.ServeFile(w, r, "./error.html")
		return
	}

	number2, err := strconv.Atoi(num2)

	if err != nil {
		http.ServeFile(w, r, "./error.html")
		return
	}

	var str string
	var op_chr string
	switch op {
	case "add":
		str = fmt.Sprintf("%d", number1+number2)
		op_chr = " + "

	case "sub":
		str = fmt.Sprintf("%d", number1-number2)
		op_chr = " - "

	case "mul":
		str = fmt.Sprintf("%d", number1*number2)
		op_chr = " * "

	case "div":
		if number2 == 0 {

			http.ServeFile(w, r, "./error.html")
			return

		} else {
			str = fmt.Sprintf("%d", number1/number2)
		}
		op_chr = " / "
	case "gcd":
		str = fmt.Sprintf("%d", gcd(number1, number2))
		op_chr = ", "
	case "lcm":
		str = fmt.Sprintf("%d", number1*number2/gcd(number1, number2))
		op_chr = ","

	default:
		http.ServeFile(w, r, "./error.html")
		return
	}
	exp := fmt.Sprintf("%d%s%d", number1, op_chr, number2)

	if op == "lcm" {
		exp = fmt.Sprintf("LCM(%d, %d)", number1, number2)
	} else if op == "gcd" {
		exp = fmt.Sprintf("GCD(%d, %d)", number1, number2)
	}

	obj := Answer{Expression: exp, Result: str}

	err = template.Must(template.ParseFiles("./index.html")).Execute(w, obj)

	if err != nil {
		http.ServeFile(w, r, "./error.html")
	}
}

func main() {
	http.HandleFunc("/", Calculator)
	log.Fatal(http.ListenAndServe(":8084", nil))
}
