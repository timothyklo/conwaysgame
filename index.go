package main

import (
"fmt"
"math/rand"
"time"
)

// create array of size "size" with random 0 or 1
func createSet(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		rand.Seed(int64(time.Now().Nanosecond()))
		arr[i] = rand.Intn(2)
	}
	return arr
}

// creates multi-dimensional array of size "size"
func createBoard(size int) [][]int {
	board := make([][]int, size)
	for i := 0; i < size ; i++ {
		board[i] = createSet(size)
	}
	return board
}

func main() {

	fmt.Print("Size of world: ")
	var input int
	fmt.Scanln(&input)

	var world = createBoard(input)
	var neighbors = make([][]int, input)

	fmt.Print(world)
	fmt.Print("\n")
	fmt.Print(neighbors)
	fmt.Print("\n")

}