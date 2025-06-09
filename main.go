package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i, " is divisible by 3 ")
		}
	}

	fmt.Println("--------------------------------")

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, " Pin Pan!")
		} else if i%3 == 0 {
			fmt.Println(i, " Pin!")
		} else if i%5 == 0 {
			fmt.Println(i, " Pan!")
		}
	}
}
