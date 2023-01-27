package db

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

type Adapter struct {
	db *leveldb.DB
}

func NewAdpater(path string) (*Adapter, error) {
	//connect
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		log.Fatalf("DB Connection Failure %v", err)
	}
	return &Adapter{db: db}, nil
}

func (da Adapter) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

func (da Adapter) AddNewGoal(title string, purpose1 string, purpose2 string) error {
	err := da.db.Put([]byte(title), []byte(purpose1+"*"+purpose2), nil)
	if err != nil {
		return err
	}
	return nil
}
