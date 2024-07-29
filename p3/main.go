package main

import "fmt"

func main(){
	var num = 600851475143
	var largestFact = 0
	var counter int = 2

	for counter*counter <= num {
		if num % counter == 0 {
			num /= counter
			largestFact = counter
		} else {
			counter++
		}
	}

	if num > largestFact {
		fmt.Println(num)
	}
}

