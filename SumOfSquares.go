package sos

import (
	"bufio"
	"fmt"
	"os"
)

// Stores the final result
var res []int

// Custom scan function to read from Stdin
func scan(r *bufio.Reader) int {
	var input int
	_, err := fmt.Fscan(r, &input)
	
	if err!=nil{
	    fmt.Println("[Fscan ERRO]Items scanned less then number of arguments.")
	}
	return input
}

// Scan the next given number of integers and calculate the sum of their square values.
func sumOfSquares(r *bufio.Reader, n int) int {
	// Base case of recursion
	if n == 0 {
		return 0
	}

	num:=scan(r)
    
	squareOfNum:=0

	if num > 0 {
		squareOfNum = num * num
	}

	return squareOfNum + sumOfSquares(r, n-1)
}

func solveEachTestCase(r *bufio.Reader) int {
	n := scan(r)

	// Get the sum of the squares
	ans := sumOfSquares(r, n)

	return ans
}

// Solves all test cases
func solve(r *bufio.Reader, t int) {
	// Base case
	if t == 0 {
		return
	}

	// Solve each test case
	val:=solveEachTestCase(r)
	
	res = append(res,val)
	solve(r, t-1)
}

func main() {
	r := bufio.NewReader(os.Stdin)

	t := scan(r)
    i:=0
	solve(r, t)
	
	// Making a loop using goto
	LOOP: if i < t {
	    fmt.Println(res[i])
	    i+=1
	    goto LOOP
	}
}
