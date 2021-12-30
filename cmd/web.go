package main

import (
	"net/http"
	"os"
	"log"
	"io/ioutil"
)

func WebServer(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", showMain)
	mux.HandleFunc("/id", showOrder)
	http.ListenAndServe("127.0.0.1:2000", mux)
}

func showOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	answer, ok := Cache[id]
	if ok{
		w.Write([]byte(answer))
	} else{
		bad_ans := "{\"order_uid\": \"" + id + " is nil\"}"
		w.Write([]byte(bad_ans))
	}
}

func showMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	htmlMain, err := os.Open("html/index.html")
	if err != nil{
		log.Fatal(err)
	}
	defer htmlMain.Close()
	htmlMainBytes, err := ioutil.ReadAll(htmlMain)
	if err != nil{
		log.Fatal(err)
	}
	w.Write([]byte(htmlMainBytes))
}