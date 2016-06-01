package domain

import (
	protoEntity "github.com/parthiban-srinivasan/mserv/proto/location"
)

type Entity struct {
	ID        string
	Type      string
	Name      string
	Longitude float64
	Timestamp int64
}

func (e *Entity) Id() string {
	return e.ID
}

func (e *Entity) Lat() float64 {
	return e.Latitude
}

func (e *Entity) Lon() float64 {
	return e.Longitude
}

func (e *Entity) ToProto() *protoEntity.Entity {
	return &protoEntity.Entity{
		Id:   e.ID,
		Type: e.Type,
		Name: e.Name,
		Location: &protoEntity.Point{
			Latitude:  e.Latitude,
			Longitude: e.Longitude,
			Timestamp: e.Timestamp,
		},
	}
}

func ProtoToEntity(e *protoEntity.Entity) *Entity {
	return &Entity{
		ID:        e.Id,
		Type:      e.Type,
		Name:      e.Name,
		Latitude:  e.Location.Latitude,
		Longitude: e.Location.Longitude,
		Timestamp: e.Location.Timestamp,
	}
}