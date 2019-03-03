package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	fmt.Println("HELLO MUTEX")
	store := CounterStore{counters: map[string]int{"i": 0, "j": 0}}
	http.HandleFunc("/get", store.get)
	http.HandleFunc("/set", store.set)
	http.HandleFunc("/inc", store.inc)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

type CounterStore struct {
	sync.Mutex
	counters map[string]int
}

func (cs *CounterStore) get(w http.ResponseWriter, r *http.Request) {
	log.Printf("get %v", r)

	cs.Lock()
	defer cs.Unlock()

	name := r.URL.Query().Get("name")
	if val, ok := cs.counters[name]; ok {
		fmt.Fprintf(w, "%s: %d/n", name, val)
	} else {
		fmt.Fprintf(w, "%s not found/n", name)
	}
}

func (cs *CounterStore) set(w http.ResponseWriter, r *http.Request) {
	log.Printf("get %v\n\n", r)

	cs.Lock()
	defer cs.Unlock()

	name := r.URL.Query().Get("name")
	val := r.URL.Query().Get("val")
	intval, err := strconv.Atoi(val)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
	} else if _, ok := cs.counters[name]; ok {
		cs.counters[name] = intval
		fmt.Fprintf(w, "ok\n")
	} else {
		fmt.Fprintf(w, "%s not found/n", name)
	}
}

func (cs *CounterStore) inc(w http.ResponseWriter, r *http.Request) {
	log.Printf("get %v", r)

	name := r.URL.Query().Get("name")
	if _, ok := cs.counters[name]; ok {
		cs.counters[name]++
		fmt.Fprintf(w, "ok\n", cs.counters[name])
	} else {
		fmt.Fprintf(w, "%s not found/n", name)
	}
}
