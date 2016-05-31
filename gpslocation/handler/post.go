package handler

import (
	"log"

	"github.com/parthiban-srinivasan/mserv/gpslocation/dao"
	proto "github.com/parthiban-srinivasan/mserv/proto/location"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/server"

	"golang.org/x/net/context"
)

func (loc *GpsLocation) Post(ctx context.Context, req *proto.PostRequest, rsp *proto.PostResponse) error {
	log.Print("Received GpsLocation.Post request")

	entity := req.GetEntity()

	if entity.GetLocation() == nil {
		return errors.BadRequest(server.DefaultOptions().Name+".save", "Require location")
	}

	p := client.NewPublication(ingester.Topic, entity)

	if err := client.Publish(ctx, p); err != nil {
		return errors.InternalServerError(server.DefaultOptions().Name+".save", err.Error())
	}

	return nil
}