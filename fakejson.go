// Server to deliver fake JSon data
package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
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
	m1 := randInt(800000, 900000)
	m2 := randInt(900000, 1000000)
	a1 := randInt(110000, 120000)
	a2 := randInt(1200000, 1300000)
	return Tvg{Profil: p, Montantinitial: m1, Montantdate: m2, Apportsretraits: a1, Apportsretraitsdate: a2}
}

type Tvgp struct {
	Unmoisperf    float64 `json:"unmoisperf"`
	Unmoispmv     int     `json:"unmoispmv"`
	Troismoisperf float64 `json:"troismoisperf"`
	Troismoispmv  int     `json:"troismoispmv"`
	Unanperf      float64 `json:"unanperf"`
	Unanpmv       int     `json:"unanpmv"`
	Troisansperf  float64 `json:"troisansperf"`
	Troisanspmv   int     `json:"troisanspmv"`
	Creationperf  float64 `json:"creationperf"`
	Creationpmv   int     `json:"creationpmv"`
}

func createtvgp() Tvgp {
	u1 := 3.49
	u2 := 109763
	t1 := 4.39
	t2 := 136928
	a1 := 7.42
	a2 := 185484
	r1 := -1.83
	r2 := 171668
	c1 := -1.83
	c2 := 171668
	/*
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
	*/
	return Tvgp{Unmoisperf: u1, Unmoispmv: u2, Troismoisperf: t1, Troismoispmv: t2, Unanperf: a1, Unanpmv: a2, Troisansperf: r1, Troisanspmv: r2, Creationperf: c1, Creationpmv: c2}
}

type Lineactif struct {
	N string `json:"n"`
	E string `json:"e"`
	V int    `json:"v"`
	W int    `json:"w"`
	D string `json:"d"`
}

func createlineactif(name string, num bool, eta string) Lineactif {
	n := name
	if num {
		n += " n°" + strconv.Itoa(randInt(1001, 9999))
	}
	e := eta
	v := randInt(10000, 200000)
	w := randInt(10000, 200000)
	d := strconv.Itoa(randInt(10, 28)) + "/" + strconv.Itoa(randInt(10, 12)) + "/2013"
	return Lineactif{N: n, E: e, V: v, W: w, D: d}
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
	case "ccs":
		o = createlineactif("CC", true, "Banque JUMP")
	case "lces":
		o = createlineactif("Livret", true, "Banque JUMP")
	case "cts":
		o = createlineactif("CT", true, "Banque JUMP")
	case "avcs":
		o = createlineactif("Compte VIE", true, "Banque JUMP")
	case "bus":
		o = createlineactif("Emprunt", false, "Direct")
	case "brs":
		o = createlineactif("T2 Sentier Paris", false, "SCI")
	case "pfs":
		o = createlineactif("SCPI", false, "Banque JUMP")
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
	http.ListenAndServe(":8080", nil)
}
