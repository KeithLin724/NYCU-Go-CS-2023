package main

import (
	"fmt"
	"math/big"
	"strconv"
	"syscall/js"
)

func CheckPrime(this js.Value, args []js.Value) interface{} {
	// Check if the number is prime
	// if len(args) != 1 {
	// 	return js.ValueOf("Invalid number of arguments")
	// }

	str := js.Global().Get("value").Get("value").String()

	number, _ := strconv.ParseUint(str, 10, 64)
	bigN := big.NewInt(int64(number))
	isProbablyPrime := bigN.ProbablyPrime(0)

	// Set the value of the "answer" element by its ID
	answerElement := js.Global().Get("document").Call("getElementById", "answer")
	if isProbablyPrime {
		answerElement.Set("innerText", "It's prime")
	} else {
		answerElement.Set("innerText", "It's not prime")
	}

	return nil
}

func registerCallbacks() {
	// Register the function CheckPrime
	js.Global().Set("CheckPrime", js.FuncOf(CheckPrime))

	// Set the initial innerText of the "answer" element
	// answerElement := js.Global().Get("document").Call("getElementById", "answer")
	// answerElement.Set("innerText", js.FuncOf(CheckPrime))
}

func main() {
	fmt.Println("Golang main function executed")
	registerCallbacks()

	//need block the main thread forever
	select {}
}
