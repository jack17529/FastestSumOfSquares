package main

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

// Used concerrency to calculate the sum of squares.
func sumOfSquares(r *bufio.Reader, n int, sum int, c chan int) {
	// Base case of recursion
	if n == 0 {
	    // Writing the total sum into the channel
		c<-sum
		return
	}

	num:=scan(r)
    
	squareOfNum:=0

	if num > 0 {
		squareOfNum = num * num
	}
    sum+=squareOfNum
    
    // Start goroutine
    go sumOfSquares(r, n-1, sum, c)
}


func solveEachTestCase(r *bufio.Reader) int {
	n := scan(r)

    // Declaring a channel which will recieve sum of squares for a test case
    c := make(chan int)
    
    // Using goroutine to calculate sum of squares
	go sumOfSquares(r, n, 0, c)
	
	// Read from the channel when the value is already computed
    ans := <-c
	
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
