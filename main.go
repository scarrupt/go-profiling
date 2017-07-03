package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func main() {
	log.Printf("Starting on port 8080")

	http.HandleFunc("/register", handleRegister)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRegister(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	yearStr := r.FormValue("year")

	reg := regexp.MustCompile(`^\d*$`)
	if !reg.MatchString(yearStr) {
		http.Error(w, "Year is invalid", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		http.Error(w, "Year is invalid", http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprintf("%s %d years old", name, time.Now().Year()-year)))
}
