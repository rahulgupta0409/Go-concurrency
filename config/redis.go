package config

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//ping the database
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Redis")
	return client
}
