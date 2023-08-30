package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var (
	steps      int
	stepsMutex sync.Mutex
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/count", countHandler)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	stepsMutex.Lock()
	defer stepsMutex.Unlock()

	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<link rel="stylesheet" href="styles.css">
		<title>Steps Counter</title>
	</head>
	<body>
		<h1>Steps Counter</h1>
		<p>Total Steps: <span id="steps">%d</span></p>
		<button id="add-steps">Add Steps</button>
		<script src="script.js"></script>
	</body>
	</html>
	`, steps)
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		stepsStr := r.FormValue("steps")
		newSteps, err := strconv.Atoi(stepsStr)
		if err != nil {
			http.Error(w, "Invalid steps input", http.StatusBadRequest)
			return
		}

		stepsMutex.Lock()
		defer stepsMutex.Unlock()

		steps += newSteps
		fmt.Fprintf(w, "%d", steps)
	}
}
