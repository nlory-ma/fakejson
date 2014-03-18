// Server to deliver fake JSon data
package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
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

type Tvg struct {
	Profil              string `json:"profil"`
	Montantinitial      int    `json:"montantinitial"`
	Montantdate         int    `json:"montantdate"`
	Apportsretraits     int    `json:"apportsretraits"`
	Apportsretraitsdate int    `json:"apportsretraitsdate"`
}

func createtvg() Tvg {
	p := []string{"Discrétionnaire", "Assurance vie", "Immobilier", "Fonds"}[randInt(0, 4)]
	m1 := randInt(80000, 90000)
	m2 := randInt(90000, 100000)
	a1 := randInt(110000, 120000)
	a2 := randInt(120000, 130000)
	return Tvg{Profil: p, Montantinitial: m1, Montantdate: m2, Apportsretraits: a1, Apportsretraitsdate: a2}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func apis(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UTC().UnixNano())
	var api string = r.URL.Path[len("/json_server/"):]
	var o interface{}
	switch api {
	case "tvg":
		o = createtvg()
	default:
		o = "Error: " + api + " (no api with this name)"
	}
	js, err := json.Marshal(o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/", apis)
	//http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}
