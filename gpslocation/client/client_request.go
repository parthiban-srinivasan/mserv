package main

import (
	"fmt"
	"time"

	proto "github.com/parthiban-srinivasan/mserv/proto/location"
	//loc "github.com/micro/geo-srv/proto/location"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"

	"golang.org/x/net/context"
)

var (
	cl proto.PpsLocationClient
)

func saveEntity() {
	entity := &proto.Entity{
		Id:   "id123",
		Type: "Hotel",
		Name: "California",
		Location: &proto.Point{
			Latitude:  51.516509,
			Longitude: 0.124615,
			Timestamp: time.Now().Unix(),
		},
	}

	_, err := cl.Post(context.Background(), &proto.PostRequest{
		Entity: entity,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Saved entity: %+v\n", entity)
}

func readEntity() {
	rsp, err := cl.Get(context.Background(), &proto.GetRequest{
		Id: "id123",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Read entity: %+v\n", rsp.Entity)
}

func searchForEntities() {
	rsp, err := cl.Search(context.Background(), &proto.SearchRequest{
		Center: &proto.Point{
			Latitude:  51.516509,
			Longitude: 0.124615,
			Timestamp: time.Now().Unix(),
		},
		Radius:      500.0,
		Type:        "Hotel",
		NumEntities: 5,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Search results: %+v\n", rsp.Entities)

}

func main() {
	// init flags
	cmd.Init()

	// use client stub
	cl = proto.NewGpsLocationClient("go.micro.srv.geo", client.DefaultClient)

	// do requests
	saveEntity()
	readEntity()
	searchForEntities()
}