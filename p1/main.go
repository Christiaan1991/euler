package main

import "fmt"

func main(){
	var i = 0
	var count = 0
	for i < 1000 {
		if i % 3 == 0 && i % 5 == 0 {
			count += i
		}
		i++
	}
	fmt.Println(count)
}