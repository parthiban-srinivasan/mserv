package handler

import (
	"log"

	"github.com/parthiban-srinivasan/mserv/gpslocation/dao"
	"github.com/parthiban-srinivasan/mserv/gpslocation/domain"
	proto "github.com/parthiban-srinivasan/mserv/proto/location"

	"golang.org/x/net/context"
)

func (loc *GpsLocation) Search(ctx context.Context, req *proto.SearchRequest, rsp *proto.SearchResponse) error {
	log.Print("Received Location.Search request")

	entity := &domain.Entity{
		Latitude:  req.Center.Latitude,
		Longitude: req.Center.Longitude,
	}

	entities := dao.Search(req.Type, entity, req.Radius, int(req.NumEntities))

	for _, e := range entities {
		rsp.Entities = append(rsp.Entities, e.ToProto())
	}

	return nil
}