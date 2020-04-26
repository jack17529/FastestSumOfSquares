package main

import (
	"bufio"
	"fmt"
	"os"
)

// Stores the final result
var res []int
var res2 []int

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
func sumOfSquares(r2 *bufio.Reader, n2 int) int {
	// Base case of recursion
	if n2 == 0 {
		return 0
	}

	num2:=scan(r2)
    
	squareOfNum2:=0

	if num2 > 0 {
		squareOfNum2 = num2 * num2
	}

	return squareOfNum2 + sumOfSquares(r2, n2-1)
}


// Used concurrency to calculate the sum of squares.
func concurrent_sumOfSquares(r *bufio.Reader, n int, sum int, c chan int) {
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
    go concurrent_sumOfSquares(r, n-1, sum, c)
}

func solveEachTestCase(r2 *bufio.Reader) int {
	n2 := scan(r2)

	// Get the sum of the squares
	ans2 := sumOfSquares(r2, n2)

	return ans2
}

func concurrent_solveEachTestCase(r *bufio.Reader) int {
	n := scan(r)

    // Declaring a channel which will recieve sum of squares for a test case
    c := make(chan int)
    
    // Using goroutine to calculate sum of squares
	go concurrent_sumOfSquares(r, n, 0, c)
	
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
	file, err := os.Open("file.txt")
	if err != nil {
    	log.Fatal(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)

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
