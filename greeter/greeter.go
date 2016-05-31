package main

import (
	"log"

	"github.com/micro/go-micro"
	proto "github.com/parthiban-srinivasan/mserv/proto/greeter"

	"golang.org/x/net/context"
)

type Greeter struct{}

func (gr *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("0.0"),
		micro.Metadata(map[string]string{
			"type": "demo",
		}),
	)

	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
