package gormdb

import (
	"testing"
)

func TestNewDB(t *testing.T) {
	dbConfig := &DBConfig{
		Server:    "192.168.7.7",
		UserName:  "feibor",
		Password:  "feibor",
		DBName:    "feibor",
		DebugMode: true,
	}
	db, err := NewDB(dbConfig)
	if err != nil {
		panic(err)
	}
	db.Get("feifei")
}
