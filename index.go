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
	fmt.Println("method:", r.Method)
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
		var world = createBoard(intSized)
		fmt.Println("whole: \n", world)
		fmt.Println("world: \n")
		for i := 0; i < intSized; i++ {
			fmt.Println("\n", world[i])
		}
		// create neighbors
		var neighbors = make([][]int, intSized)
		fmt.Println("neighbors: \n", neighbors)
		// show on template
		t, _ := template.ParseFiles("conway.gtpl")
		t.Execute(w, world)
	}
}

func main() {
	http.HandleFunc("/", getsize)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

