package main

import (
	"log"
	"time"

  _ "net/http/pprof"
	"github.com/parthiban-srinivasan/mserv/gpslocation/handler"
	"github.com/parthiban-srinivasan/mserv/gpslocation/ingester"
	"github.com/micro/go-micro"
	proto "github.com/parthiban-srinivasan/mserv/proto/location"
	//	"golang.org/x/net/context"
)

func main() {

    // Register server def
	service := micro.NewService(
		micro.Name("go.micro.srv.gpsloc"),
		micro.Version("0.1"),
		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(time.Second*30),
	)

	// Initialize Server
	service.Init()

	// Register Handlers
	proto.RegisterGpsLocationHandler(service.Server(), new(handler.GpsLocation))

	// Register Subscriber
	service.Server().Subscribe(
		service.Server().NewSubscriber(
			ingester.Topic,
			new(ingester.Geo),
		),
	)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

	//End of Program
}