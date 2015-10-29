package main

import (
	"fmt"
	"gopkg.in/redis.v3"
	"os"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("System Problem But it worked!")
		os.Exit(0)
	}

	fmt.Println("Entering Program and attempting to connect to %s", os.Args[1])
	client := redis.NewClient(&redis.Options{
		Addr:     os.Args[1],
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()

	fmt.Println("Attempting Read!")
	fmt.Println("Setting The PingRedis key")
	err := client.Set("PingRedis", time.Now().Format(time.RFC1123), 0).Err()

	if err != nil {
		fmt.Println("Could not read")
		return
		//panic(err)
	}

	val, errPing := client.Get("PingRedis").Result()
	fmt.Println("Getting The PingRedis key that we set")
	fmt.Println(val)

	if errPing != nil {
		panic(err)
	}

	fmt.Println("Attempting delete of PingRedis key")
	errDel := client.Del("PingRedis").Err()

	if errDel != nil {
		panic(err)
	}

	fmt.Println("Attempting Ping Pong")
	pong, err := client.Ping().Result()

	fmt.Println(pong, err)

}
