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
	// sets size of world
	var size int = 39

	// create world
	world := createBoard(size)

	// create neighbors array to be used to count neighbors
	neighbors := createBoard(size)

	// shows world
	fmt.Println(world)
	fmt.Println(neighbors)
}