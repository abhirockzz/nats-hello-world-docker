package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/nats-io/go-nats"
)

func main() {
	subject := "foo"
	natsURL := "nats://server:4222"

	opts := nats.Options{
		AllowReconnect: true,
		MaxReconnect:   5,
		ReconnectWait:  5 * time.Second,
		Timeout:        3 * time.Second,
		Url:            natsURL,
	}

	conn, _ := opts.Connect()
	//defer conn.Close()
	fmt.Println("Subscriber connected to NATS server")

	fmt.Printf("Subscribing to subject %s\n", subject)
	conn.Subscribe(subject, func(msg *nats.Msg) {
		fmt.Printf("Got message '%s\n", string(msg.Data)+"'")
	})

	runtime.Goexit()
}
