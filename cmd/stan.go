package main

import (
	"../valid"
	"../repository"
	"github.com/nats-io/stan"
	"encoding/json"
	"log"
	"fmt"
)

func FooMsgHandler(m *stan.Msg) {
	var MyModel repository.Order
	isMessageValidJSON := valid.ValidateMyJSON(m.Data)
	fmt.Println("New message is valid ", isMessageValidJSON==1)
	if isMessageValidJSON == 1{
		json.Unmarshal(m.Data, &MyModel)
		MyModel.Data = string(m.Data)
		if _, ok := Cache[MyModel.Order_uid]; !ok { 
			db, err := repository.ConnectToDB()
			if err != nil{
				log.Fatal(err)
			}
			defer db.Close()
			repository.AddOrderTx(db, MyModel)
			Cache[MyModel.Order_uid] = string(m.Data)
			fmt.Println("Added new order with id: ", MyModel.Order_uid)
		} else{
			fmt.Println("\"" + MyModel.Order_uid + "\"" + "  order_uid is already in DB")
		}
	}
}

func StanConnect(cluster, client, url string) stan.Conn{
	
	sc, err := stan.Connect(
		cluster,
		client,
		stan.Pings(1, 3),
		stan.NatsURL(""),
	)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Printf("Connected to cluster \"%s\" as client \"%s\"...\n", cluster, client)
	return sc
}