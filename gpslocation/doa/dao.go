package dao

import (
	"sync"

	geo "github.com/hailocab/go-geoindex"
	"github.com/parthiban-srinivasan/msrv/gplslocation/domain"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/server"
)

var (
	mtx          sync.RWMutex
	geoIndex = geo.NewPointsIndex(geo.Km(1.0))
)

func Get(id string) (*domain.Entity, error) {
	mtx.RLock()
	defer mtx.RUnlock()

	pt := geoIndex.Get(id)
	if pt == nil {
		return nil, errors.NotFound(server.DefaultOptions().Name+".Get", "Not found")
	}

	entity, ok := pt.(*domain.Entity)
	if !ok {
		return nil, errors.InternalServerError(server.DefaultOptions().Name+".Get", "Error reading entity")
	}

	return entity, nil
}

func Save(e *domain.Entity) {
	mtx.Lock()
	geoIndex.Add(e)
	mtx.Unlock()
}