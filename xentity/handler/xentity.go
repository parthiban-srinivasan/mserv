package handler

import (
	//"encoding/json"
	"fmt"
	//"net/url"
	//"strings"

	dao "github.com/parthiban-srinivasan/mserv/xentity/dao"
	proto "github.com/parthiban-srinivasan/mserv/proto/xentity"
	"github.com/micro/go-micro/errors"

	"golang.org/x/net/context"
)

type Xentity struct {}

func (g *Xentity) Get(ctx context.Context, req *proto.GetRequest, rsp *proto.GetResponse) error {

	id := req.Xid

	if len(id) == 0 {
		return errors.BadRequest(server.DefaultOptions().Name+".Get", "Require Entity Id")
	}

    db := InitDB(dao.DefaultDbPath)
    defer db.Close()
    
    var qresp []domain.XEntity
    
    qresp = ReadXEntity(db)
    //CreateDB(db)
 
	if err != nil {
		return errors.InternalServerError("go.micro.srv.geocode.Google.Geocode", err.Error())
	}
	
	if err := json.Unmarshal(b, &rsp); err != nil {
		return errors.InternalServerError("go.micro.srv.xentity", err.Error())
	}
	return nil
}