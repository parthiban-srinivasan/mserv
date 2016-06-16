package realpropdao

import (
    
	"sync"
	"log"

	"github.com/parthiban-srinivasan/mserv/geocode/dbs/realprop"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

)

func InitDB(filepath string) *sql.DB {
    
	db, err := sql.Open("sqlite3", filepath)
    if err != nil {
		log.Fatal(err)
	}
	if db == nil {
	    log.Fatal("db nil")
	}
	return db
}

func CreateRealprop(db *sql.DB) {
    
	// create table if not exists
	create_sql_table := `
	CREATE TABLE IF NOT EXISTS realprop(
		Guid INTEGER NOT NULL PRIMARY KEY,
		Type TEXT,
	//	Class TEXT,
		Value TEXT
	//	InsertedDatetime DATETIME
	);
	`
	_, err := db.Exec(create_sql_table)
	if err != nil { 
	    log.Fatal("create hot table failed")
	}
}

func StoreRealprop(db *sql.DB, rec *domain.Realprop) {
	sql_addrec := `
	INSERT OR REPLACE INTO realprop(
		Guid,
		Type,
//		Class,
		Value,
//		InsertedDatetime
	) values(?, ?, ?, ?, CURRENT_TIMESTAMP)
	`

	add_stmt, err := db.Prepare(sql_addrec)
	if err != nil { log.Fatal("insert preparation failed") }
	defer add_stmt.Close()

	_, err := add_stmt.Exec(rec.Guid, rec.Type, rec.Class, rec.Value)
	if err != nil { log.Fatal("insert  failed - realprop") }
}

func ReadRealprop(db *sql.DB) []*domain.Realprop {
	sql_readall := `
	SELECT Guid, Type, Value FROM realprop
	ORDER BY datetime(InsertedDatetime) DESC
	`

	rows, err := db.Query(sql_readall)
	if err != nil { log.Fatal("Readall  Query failed - realprop") }
	defer rows.Close()

	var result []*domain.Realprop
	for rows.Next() {
		var rrec *domain.Realprop{}
		err2 := rows.Scan(&rrec.Guid, &rrec.Type, &rrec.Value)
		if err2 != nil  { log.Fatal("Readall  scan failed - realprop") }
		result = append(result, rrec)
	}
	return result
}