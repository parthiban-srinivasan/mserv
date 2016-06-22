package xentitydao

import (
    
	"sync"
	"log"

	"github.com/parthiban-srinivasan/mserv/xentity/domain"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

)

var  (
    Key                string
	DbPath             string
	DbType             string
	DefaultDbPath    = "xentity"
//	DefaultDBType    = ":memory:"
	)
	
func InitDB(filepath string) *sql.DB {
    
    if filepath == nil {
		filepath == DefaultDbPath
    }
    
	db, err := sql.Open("sqlite3", filepath)
    if err != nil {
		log.Fatal(err)
	}
	if db == nil {
	    log.Fatal("db nil")
	}
	return db
}

func CreateXEntity(db *sql.DB) {
    
	// create table if not exists
	create_sql_table := `
	CREATE TABLE IF NOT EXISTS xentity(
		Xid INTEGER NOT NULL PRIMARY KEY,
		Type TEXT,
	//	Class TEXT,
		Value TEXT
	//	InsertedDatetime DATETIME
	);
	`
	_, err := db.Exec(create_sql_table)
	if err != nil { 
	    log.Fatal("create xentity table failed")
	}
}

func StoreXEntity(db *sql.DB, rec domain.XEntity) {
	sql_addrec := `
	INSERT OR REPLACE INTO xentity(
		Xid,
		Type,
//		Class,
		Value,
//		InsertedDatetime
	) values(?, ?, ?, ?, CURRENT_TIMESTAMP)
	`

	add_stmt, err := db.Prepare(sql_addrec)
	if err != nil { log.Fatal("insert preparation failed") }
	defer add_stmt.Close()

	_, err := add_stmt.Exec(rec.Xid, rec.Type, rec.Class, rec.Value)
	if err != nil { log.Fatal("insert  failed - xentity") }
}

func ReadXEntity(db *sql.DB) []domain.XEntity {
	sql_readall := `
	SELECT Xid, Type, Value FROM xentity
	ORDER BY datetime(InsertedDatetime) DESC
	`

	rows, err := db.Query(sql_readall)
	if err != nil { log.Fatal("Readall  Query failed - xentity") }
	defer rows.Close()

	var result []domain.XEntity
	for rows.Next() {
		var rrec domain.XEntity{}
		err2 := rows.Scan(&rrec.Xid, &rrec.Type, &rrec.Value)
		if err2 != nil  { log.Fatal("Readall  scan failed - xentity") }
		result = append(result, rrec)
	}
	return result
}