package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter text: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	angka, err := strconv.Atoi(text)

	if err == nil {
		fibo := fibonaci(angka)
		fmt.Println(fibo)

		fak := faktorial(angka)
		fmt.Println(fak)

		fibofak := fibo + fak
		fmt.Println(fibofak)
	}

func fibonaci(x int) int {
	a := 1
	b := 1
	c := 1
	for i := 0; i < x; i++ {
		if i > 1 {
			c = a + b
			a = b
			b = c
		}
	}
	return c
}

func faktorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * faktorial(x-1)
}


