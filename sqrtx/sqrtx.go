package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mySqrtPocketCalucatorAlgo(100))
	fmt.Println(mySqrtBinarySearch(6))
	fmt.Println(mySqrtUsingRecurson(65))
	fmt.Println(mySqrtUsingNewtonsMethod(65))	
}

//Best method: Newton's method
//time: O(log N)
//space: O(1)
func mySqrtUsingNewtonsMethod(x int) int {
	if x < 2 {
			return x
	}
	
	x0 := float64(x)
	x1 := (x0 + float64(x) / x0) / 2.0
	
	for math.Abs(x0-x1) >= 1 {
			x0 = x1
			x1 = (x0 + float64(x) / x0) / 2.0
	}
	return int(x1)
}

//Binary search
//time complextiy : O(log N)
//space complexity: O(1)
func mySqrtBinarySearch(x int) int {
	if x < 2 {
		return x
	}

	left := 2
	right := x / 2
	var num int64
	var pivot int

	for left <= right {
		pivot = left + (right-left)/2
		num = int64(pivot * pivot)

		if num > int64(x) {
			right = pivot - 1
		} else if num < int64(x) {
			left = pivot + 1
		} else {
			return pivot
		}
	}

	return right

}

//Pocket Calculator Algorithm
//squre root of x = e power of ((1/2) * log x)
// time complexity: O(1)
// space complexity: O(1)
func mySqrtPocketCalucatorAlgo(x int) int {
	if x < 2 {
		return x
	}

	left := int(math.Pow(math.E, 0.5*math.Log(float64(x))))
	right := left + 1

	if right*right > x {
		return left
	} else {
		return right
	}
}

//recursion + bit shift
//time: O(log N)
//space: O(log N)
func mySqrtUsingRecurson(x int) int {
	if x < 2 {
			return x
	}
	
	left := mySqrtUsingRecurson(x >> 2) << 1
	right := left + 1    
	
	if int64(right*right) > int64(x)  {
			return left
	} else {
			return right
	}     
}

