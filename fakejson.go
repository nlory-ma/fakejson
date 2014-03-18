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

type Tvgp struct {
	Unmoisperf    int `json:"unmoisperf"`
	Unmoispmv     int `json:"unmoispmv"`
	Troismoisperf int `json:"troismoisperf"`
	Troismoispmv  int `json:"troismoispmv"`
	Unanperf      int `json:"unanperf"`
	Unanpmv       int `json:"unanpmv"`
	Troisansperf  int `json:"troisansperf"`
	Troisanspmv   int `json:"troisanspmv"`
	Creationperf  int `json:"creationperf"`
	Creationpmv   int `json:"creationpmv"`
}

func createtvgp() Tvgp {
	u1 := randInt(5, 15)
	u2 := randInt(500, 1500)
	t1 := randInt(5, 15)
	t2 := randInt(500, 1500)
	a1 := randInt(50, 150)
	a2 := randInt(5000, 15000)
	r1 := randInt(50, 150)
	r2 := randInt(5000, 15000)
	c1 := randInt(500, 1500)
	c2 := randInt(50000, 150000)
	return Tvgp{Unmoisperf: u1, Unmoispmv: u2, Troismoisperf: t1, Troismoispmv: t2, Unanperf: a1, Unanpmv: a2, Troisansperf: r1, Troisanspmv: r2, Creationperf: c1, Creationpmv: c2}
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
	case "tvgp":
		o = createtvgp()
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
