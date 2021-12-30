package main

import (
	"fmt"
	"log"

	"../repository"
)

var Cache map[string]string

func main(){
	Cache = make(map[string]string)
	repository.FillCache(Cache)
	natsConn := StanConnect("test-cluster", "client-1", "")
	defer natsConn.Close()
	subcr, err := natsConn.Subscribe("foo", FooMsgHandler)
	if err != nil {
		log.Fatalln(err)
	} else{
		fmt.Println("Subscription is succeeded...")
	}
	defer subcr.Unsubscribe()
	WebServer()
}


