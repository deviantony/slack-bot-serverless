package msg

import (
	"context"
	"fmt"
)

type Event struct {
	Message string `json:"message"`
}

func Main(ctx context.Context, args map[string]interface{}) map[string]interface{} {
	// fmt.Printf("Received event: %+v\n", event)

	name := args["message"].(string)
	if name == "" {
		name = "World"
	}

	msg := make(map[string]interface{})
	msg["body"] = fmt.Sprintf("Event: %+v\n", args)
	return msg
}
