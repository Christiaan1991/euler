package main

import "fmt"

func main(){
	var i = 0
	var sum = 0
	for {
		value := fib(i)
		if value >= 4000000 {
			break
		}
		if value % 2 == 0 {
			sum += value
		}
		i++
	}
	fmt.Println(sum)
}

func fib(n int) int{
	if n == 0 {
		return 1
	} else if n == 1 {
		return 2
	} else {
		return fib(n-2) + fib(n-1)
	}
}
