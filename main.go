package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main()  {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	if err := rdb.Ping(timeout).Err(); err != nil {
		log.Fatalln(err)
	}

	log.Println("done.")

	rdb.Set(context.Background(), "wallet1", "100.00", 10 * time.Minute)


	subscribe := rdb.PSubscribe(context.Background(), "__key*__:*")
	// sub key modified only
	for msg := range subscribe.Channel() {
		fmt.Println(msg)
	}

}