package main

import (
	"log"

	"github.com/micro/cli"
	micro "github.com/micro/go-micro"

	"github.com/parthiban-srinivasan/mserv/geocode/googlemap"
	"github.com/parthiban-srinivasan/mserv/geocode/handler"
	proto "github.com/parthiban-srinivasan/mserv/proto/geomap"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.geocode"),
		micro.Flags(
			cli.StringFlag{
				Name:   "api_key",
				EnvVar: "API_KEY",
				Usage:  "Vendor maps API key",
			},
			cli.StringFlag{
				Name:   "client_id",
				EnvVar: "CLIENT_ID",
				Usage:  "Vendor client id",
			},
			cli.StringFlag{
				Name:   "google_signature",
				EnvVar: "SIGNATURE",
				Usage:  "Vendor signature",
			},
		),
		micro.Action(func(ctx *cli.Context) {
			googlemap.Key = ctx.String("api_key")
			googlemap.ClientID = ctx.String("client_id")
			googlemap.Signature = ctx.String("signature")
		}),
	)

	service.Init()

	proto.RegisterGeomapHandler(service.Server(), new(handler.Geomap))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}