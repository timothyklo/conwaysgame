package main

import (
"fmt"
"math/rand"
"time"
"html/template"
"log"
"net/http"
"strconv"
)

var world = [][]int{}
var neighbors = [][]int{}
var intSized int = 0

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

// get the size from the user
func getsize(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting size:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("getsize.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// changing slice to string
		var stringSize string = r.FormValue("size")
		// changing string to integer
		intSized, err := strconv.Atoi(stringSize)
		if err != nil {
			log.Fatal(err)
		}
		// create world
		world = createBoard(intSized)
		fmt.Println("world: \n", world)
		fmt.Println("appears: \n")
		for i := 0; i < intSized; i++ {
			fmt.Println(world[i], "\n")
		}
		fmt.Println("intsize: \n", intSized)
		// create neighbors
		neighbors = createBoard(intSized)
		fmt.Println("neighbors: \n", neighbors)
		// show on template
		t, _ := template.ParseFiles("conway.gtpl")
		t.Execute(w, world)
	}
}

func runlife(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("running life \n")
	// fmt.Println("world: \n", world)

	// logic to populate neighbors
	for i := 0; i < len(world); i++ {
		for j :=0; j < len(world[i]); j++ {
			// fmt.Println(world[i][j])
			var count = 0
			
			neighbors[i][j] = count
		}
	  fmt.Println(world[i])
	}
	fmt.Println("neighbors: \n", neighbors)

	t, _ := template.ParseFiles("conway.gtpl")
	t.Execute(w, neighbors)
}

func main() {
	http.HandleFunc("/getsize", getsize)
	http.HandleFunc("/runlife", runlife)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
