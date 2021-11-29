package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var db = map[string]interface{}{
	"test": "1234",
}

var RWMutex = &sync.RWMutex{}

type pair struct {
	Key   string
	Value interface{}
}

func getValue(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	key := req.URL.Path[len("/db/"):]

	RWMutex.RLock()
	if val, ok := db[key]; ok {
		fmt.Fprint(w, val)
	} else {
		http.Error(w, "key not exist", http.StatusNotFound)
	}
	RWMutex.RUnlock()
}

func postValue(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	buf := &pair{}
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(buf); err != nil {
		http.Error(w, "json format incorrect", http.StatusBadRequest)
	}

	RWMutex.Lock()
	db[buf.Key] = buf.Value
	RWMutex.Unlock()
}

func dbHanlder(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		getValue(w, req)
	case "POST":
		postValue(w, req)
	default:
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/db/", dbHanlder)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
