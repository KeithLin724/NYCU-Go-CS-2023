package main

import "fmt"

func main() {
	var n int64

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	result := Sum(n)
	fmt.Println(result)
}

func Sum(n int64) string {
	// TODO: Finish this function

	ans, str, last := 0, "", n

	if n%7 == 0 {
		last--
	}

	for i := 1; i < int(last); i++ {
		if i%7 == 0 {
			continue
		}
		str += fmt.Sprintf("%d+", i)
		ans += i
	}

	return fmt.Sprintf("%s%d=%d", str, last, ans+int(last))
}
