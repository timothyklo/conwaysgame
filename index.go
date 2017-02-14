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
		// create temporary neighbors
		neighbors = createBoard(intSized)
		// show on template
		t, _ := template.ParseFiles("conway.gtpl")
		t.Execute(w, world)
	}
}

func runlife(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("running life")
	// logic to populate neighbors
	for i := 0; i < len(world); i++ {
		for j :=0; j < len(world[i]); j++ {
			var count = 0
			if i>0 && j > 0 {
				if world[i-1][j-1] == 1 {
					count++
				}
			}
			if i>0 {
				if world[i-1][j] == 1 {
					count++
				}
			}
			if j > 0 {
				if world[i][j-1] == 1 {
					count++
				}
			}
			if i > 0 && j < len(world[i])-1 {
				if world[i-1][j+1] == 1 {
					count++
				}
			}
			if i < len(world[i])-1 && j > 0 {
				if world[i+1][j-1] == 1 {
					count++
				}
			}
			if j < len(world[i])-1 {
				if world[i][j+1] == 1 {
					count++
				}
			}
			if i < len(world[i])-1 {
				if world[i+1][j] == 1 {
					count++
				}
			}
			if i < len(world[i])-1 && j < len(world[i])-1 {
				if world[i+1][j+1] == 1 {
					count++
				}
			}
			neighbors[i][j] = count
		}
	}
	// go through neighbors and reset world
	for i := 0; i < len(world); i++ {
		for j :=0; j < len(world[i]); j++ {
			// Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
			if world[i][j] == 1 && neighbors[i][j] < 2 {
				world[i][j] = 0
			}
			// Any live cell with two or three live neighbours lives on to the next generation.
			// Any live cell with more than three live neighbours dies, as if by overpopulation.
			if world[i][j] == 1 && neighbors[i][j] > 3 {
				world[i][j] = 0
			}
			// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
			if world[i][j] == 0 && neighbors[i][j] == 3 {
				world[i][j] = 1
			}
		}
	}
	t, _ := template.ParseFiles("conway.gtpl")
	t.Execute(w, world)
}

func main() {
	http.HandleFunc("/getsize", getsize)
	http.HandleFunc("/runlife", runlife)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
