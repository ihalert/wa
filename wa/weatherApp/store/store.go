package store

import (
	"database/sql"
	"fmt"
	"github.com/mivan/weatherApp/entities"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//
//

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func InsertData() {

}

func (t *Store) GetSqlDataInMap() entities.DBStruct {
	// treba i error vraćat
	dbstruct := entities.DBStruct{}

	return dbstruct
}

func DeleteData() {

}
