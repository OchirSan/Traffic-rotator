package database

//
//import (
//	"github.com/jmoiron/sqlx"
//	"github.com/rs/zerolog/log"
//	"time"
//)
//
//type collector interface {
//	Set(events *Event)
//	Get() []Event
//}
//
//type database struct {
//	psql            *sqlx.DB
//	eventsCollector collector
//}
//
//func New(db *sqlx.DB, eventsCollector collector) *database {
//	databaseEntity := &database{
//		psql:            db,
//		eventsCollector: eventsCollector,
//	}
//	databaseEntity.startBulkInserter()
//	return databaseEntity
//}
//
//func (d *database) AsyncUpdate(event *Event) {
//	d.eventsCollector.Set(event)
//}
//
//func (d *database) startBulkInserter() {
//	go func() {
//		for range time.NewTicker(10 * time.Minute).C {
//			d.updateEvents()
//		}
//	}()
//}
//
//const batchSize = 100
//
//func (d *database) updateEvents() {
//	events := d.eventsCollector.Get()
//	if len(events) == 0 {
//		return
//	}
//	for i := 0; i < len(events); i += batchSize {
//		var end = minInt(i+batchSize, len(events))
//		_, err := d.psql.NamedExec(
//			"INSERT INTO clicks (event_time, click_url) VALUES (:event_time, :click_url)", events[i:end])
//		if err != nil {
//			log.Error().Err(err).Msg("Can't insert creative statistics")
//		}
//	}
//
//}
//
//func minInt(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}
