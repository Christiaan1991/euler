package main

import "fmt"

func main(){
	var highestPalindrome = 0
	for i:=1000; i > 0; i-- {
		for j:=1000; j > 0; j-- {
			if isPalindrome(i*j) {
				if i*j > highestPalindrome {
					highestPalindrome = i*j
					fmt.Printf("%d x %d = %d\n", i, j, highestPalindrome)
				}
			}
		}
	}

}

func isPalindrome(num int) bool {
	copyNum := num
	var newnum int = 0
	for copyNum != 0 {
		digit := copyNum % 10
		newnum = newnum*10 + digit
		copyNum /= 10
	}
	if newnum == num {
		return true
	}
	return false

}
