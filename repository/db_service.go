package repository

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"os"
	"github.com/jackc/pgx/pgxpool"
)

func ConnectToDB() (*pgxpool.Pool, error) {
	connstr := os.Getenv("CONNSTR")
	if connstr == "" {
		log.Fatal("не указан адрес подключения к БД")
	 }
	db, err := pgxpool.Connect(context.Background(), connstr)
	return db, err
}

func GetStat(db *pgxpool.Pool){
	stats := db.Stat()
	fmt.Println("Статистика подключений к БД:")
	fmt.Println("Максимум подключений к БД:", stats.MaxConns())
	fmt.Println("Всего задействовано подключений к БД:", stats.TotalConns())
	fmt.Println("Занято подключений к БД:", stats.AcquiredConns())
	fmt.Println("Свободно подключений к БД:", stats.IdleConns())
	fmt.Println("Всего выделено подключений за все время:", stats.AcquireCount())
}

func ConnectWithMoreConn(connsToDBPerCore int) (*pgxpool.Pool, error){
	// to set n simultanious connection to db per CPU core
	connstr := os.Getenv("CONNSTR")
	config, err := pgxpool.ParseConfig(connstr)
	if err != nil{
		return nil, err
	}
	config.MaxConns = int32(runtime.NumCPU() * connsToDBPerCore)
	db, err := pgxpool.ConnectConfig(context.Background(), config)
	
	if err != nil{
		log.Fatal("Connection to DB is not succeed %w", err)
	}
	return db, err
}