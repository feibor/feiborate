package gormdb

import (
	"testing"
)

func TestNewDB(t *testing.T) {
	dbConfig := &DBConfig{
		Server:    "192.168.7.7",
		UserName:  "feiborate",
		Password:  "feiborate",
		DBName:    "feiborate",
		DebugMode: true,
	}
	db, err := NewDB(dbConfig)
	if err != nil {
		panic(err)
	}
	db.Get("feifei")
}
