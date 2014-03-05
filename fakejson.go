// Server to deliver fake JSon data
package main

import (
	"encoding/json"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string
}

type Linechart struct {
}

type Piechardata struct {
	Titre       string
	Actions     int
	Fonds       int
	Obligations int
	Liquidites  int
}

type Areachart struct {
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
