package main

import (
	"log"

	"github.com/micro/cli"
	micro "github.com/micro/go-micro"

	"github.com/parthiban-srinivasan/mserv/xentity/dao"
	"github.com/parthiban-srinivasan/mserv/xentity/handler"
	proto "github.com/parthiban-srinivasan/mserv/proto/xentity"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.xentity"),
		micro.Flags(
			cli.StringFlag{
				Name:   "api_key",
				EnvVar: "API_KEY",
				Usage:  "Vendor API key",
			},
			cli.StringFlag{
				Name:   "db_id",
				EnvVar: "DB_ID",
				Usage:  "Database name - identifier",
			},
			cli.StringFlag{
				Name:   "db_type",
				EnvVar: "DB_TYPE",
				Usage:  "DASD vs Inmemory storage",
			},
		),
		micro.Action(func(ctx *cli.Context) {
			xentitydao.Key = ctx.String("api_key")
			xentitydao.DbID = ctx.String("db_id")
			xentitydao.DbType = ctx.String("db_type")
		}),
	)

	service.Init()

	proto.RegisterXentityHandler(service.Server(), new(handler.Xentity))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
