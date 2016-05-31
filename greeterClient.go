package main

import (
    "fmt"

    micro "github.com/micro/go-micro"
	proto "github.com/parthiban-srinivasan/mserv/proto"
    "golang.org/x/net/context"
)


func main() {
    // Create a new service. Optionally include some options here.
    service := micro.NewService(
            micro.Name("greeter.client"),
    )

    // Create new greeter client
    greeter := proto.NewGreeterClient("greeter", service.Client())

    // Call the greeter
    rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
    if err != nil {
        fmt.Println(err)
    }

    // Print response
    fmt.Println(rsp.Greeting)
}