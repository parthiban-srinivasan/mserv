package ingester

import (
	"log"

	"github.com/parthiban-srinivasan/mserv/gpslocation/dao"
	"github.com/parthiban-srinivasan/mserv/gpslocation/domain"
	proto "github.com/parthiban-srinivasan/mserv/proto/location"
	"golang.org/x/net/context"
)

var (
	Topic = "geo.location"
)

type Geo struct{}

func (g *Geo) Handle(ctx context.Context, e *proto.Entity) error {
	log.Printf("Saving entity ID %s", e.Id)
	dao.Save(domain.ProtoToEntity(e))
	return nil
}