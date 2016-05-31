package main

import (
	"fmt"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	proto "github.com/parthiban-srinivasan/mserv/proto"

	"golang.org/x/net/context"
)

func main() {
    
	cmd.Init()

	// Create new request to service go.micro.srv.greeter, method Say.Hello
	req := client.NewRequest("greeter", "Greeter.Hello", &proto.HelloRequest{
		Name: "John",
	})

	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp := &proto.HelloResponse{}

	// Call service
	if err := client.Call(ctx, req, rsp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}