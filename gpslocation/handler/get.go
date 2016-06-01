package handler

import (
	"log"

	"github.com/parthiban-srinivasan/mserv/gpslocation/dao"
	proto "github.com/parthiban-srinivasan/mserv/proto/location"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/server"

	"golang.org/x/net/context"
)

type GpsLocation struct{}

func (loc *GpsLocation) Get(ctx context.Context, req *proto.GetRequest, rsp *proto.GetResponse) error {
	log.Print("Received GpsLocation.Get request")

	id := req.Id

	if len(id) == 0 {
		return errors.BadRequest(server.DefaultOptions().Name+".Get", "Require Id")
	}

	log.Print("Received GpsLocation.Get valid ID %q", id)
	entity, err := dao.Get(id)
	if err != nil {
		return err
	}

	log.Print("Received GpsLocation.Get after DAO")
	rsp.Entity = entity.ToProto()

	return nil
}