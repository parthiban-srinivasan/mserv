package handler

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	dao "github.com/parthiban-srinivasan/mserv/xentity/dao"
	"github.com/parthiban-srinivasan/mserv/xentity/domain"
	proto "github.com/parthiban-srinivasan/mserv/proto/xentity"
	"github.com/micro/go-micro/errors"

	"golang.org/x/net/context"
)

func (g *Xentity) Get(ctx context.Context, req *proto.GetRequest, rsp *proto.GetResponse) error {

	id := req.Id

	if len(id) == 0 {
		return errors.BadRequest(server.DefaultOptions().Name+".Get", "Require Entity Id")
	}

    db := InitDB(dao.DefaultDbPath)
    defer db.Close()
    
    create
	b, err := googlemap.Do("geocode", u)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.geocode.Google.Geocode", err.Error())
	}
	if err := json.Unmarshal(b, &rsp); err != nil {
		return errors.InternalServerError("go.micro.srv.xentity", err.Error())
	}
	return nil
} 
