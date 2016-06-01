package dao

import (
	"sync"

	geo "github.com/hailocab/go-geoindex"
	"github.com/parthiban-srinivasan/mserv/gpslocation/domain"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/server"
)

var (
	mtx          sync.RWMutex
	geoIndex = geo.NewPointsIndex(geo.Km(5.0))
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

func Search(typ string, entity *domain.Entity, radius float64, numEntities int) []*domain.Entity {
	mtx.RLock()
	defer mtx.RUnlock()

	points := geoIndex.KNearest(entity, numEntities, geo.Meters(radius), func(p geo.Point) bool {
		e, ok := p.(*domain.Entity)
		if !ok || e.Type != typ {
			return false
		}
		return true
	})

	var entities []*domain.Entity

	for _, point := range points {
		e, ok := point.(*domain.Entity)
		if !ok {
			continue
		}
		entities = append(entities, e)
	}

	return entities
}