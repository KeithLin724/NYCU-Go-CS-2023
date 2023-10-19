package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Calculator(w http.ResponseWriter, r *http.Request) {
	// TODO: implement a calculator

	strCmd := strings.Split(r.URL.Path, "/")
	strCmd = strCmd[1:]
	// fmt.Println(strCmd)
	if len(strCmd) != 3 {
		w.Write([]byte("Error!"))
		return
	}
	str := ""

	number1, err := strconv.Atoi(strCmd[1])
	if err != nil {
		w.Write([]byte("Error!"))
		return
	}
	number2, err := strconv.Atoi(strCmd[2])

	if err != nil {
		w.Write([]byte("Error!"))
		return
	}

	switch strCmd[0] {
	case "add":
		str = fmt.Sprintf("%d + %d = %d", number1, number2, number1+number2)

	case "sub":
		str = fmt.Sprintf("%d - %d = %d", number1, number2, number1-number2)

	case "mul":
		str = fmt.Sprintf("%d * %d = %d", number1, number2, number1*number2)

	case "div":
		if number2 == 0 {
			str = "Error!"
		} else {
			str = fmt.Sprintf("%d / %d = %d, reminder = %d", number1, number2, number1/number2, number1%number2)
		}

	default:
		str = "Error!"
	}
	fmt.Println(str)
	w.Write([]byte(str))

}

func main() {
	http.HandleFunc("/", Calculator)
	log.Fatal(http.ListenAndServe(":8083", nil))
}
