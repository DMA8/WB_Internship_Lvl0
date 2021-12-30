package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"github.com/nats-io/stan"
)

func main(){
	sc, err := stan.Connect(
		"test-cluster",
		"client-2",
		stan.Pings(1, 3),
		stan.NatsURL(strings.Join(os.Args[1:], ",")),
	)
	if err != nil{
		log.Fatalln(err)
	}
	defer sc.Close()

	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer sub.Unsubscribe()
	jsonFile, err := os.Open("message_sender/validMessages/model.json")
	if err != nil{
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValJson, err := ioutil.ReadAll(jsonFile)
	if err != nil{
		log.Fatal(err)
	}
	if err := sc.Publish("foo", byteValJson); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Message is sent!")
}