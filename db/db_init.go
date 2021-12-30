package main

import (
	"context"
	"fmt"
	"log"
	"../repository"
	"github.com/jackc/pgx/pgxpool"
)

func main(){
	db, err := repository.ConnectToDB()
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	err = CreateOrders(db)
	if err != nil{
		log.Fatal(err)
	}
	err = cleanAllTables(db)
	if err != nil{
		log.Fatal(err)
	}
}
func CreateOrders(db *pgxpool.Pool) error{
	ctx := context.Background()
	query := `CREATE TABLE IF NOT EXISTS Orders (
		order_uid	varchar(45) NOT NULL,
		data		text NOT NULL,
		PRIMARY KEY (order_uid)
	  )`
	 tag, err := db.Exec(ctx, query)
	 fmt.Println(tag)
	 return err
}

func cleanAllTables(db *pgxpool.Pool) error{
	ctx := context.Background()
	query := `TRUNCATE TABLE Orders`
	tag, err := db.Exec(ctx, query)


	fmt.Println(tag)
	return err
}