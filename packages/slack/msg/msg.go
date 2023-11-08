package main

import (
	"context"
	"fmt"
)

type Event struct {
	Message string `json:"message"`
}

func Main(ctx context.Context, event Event) map[string]interface{} {
	fmt.Printf("Received event: %+v\n", event)

	name := event.Message
	if name == "" {
		name = "World"
	}

	msg := make(map[string]interface{})
	msg["body"] = "Hello " + name + "!"
	return msg
}
