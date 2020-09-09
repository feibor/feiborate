package gormdb

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestNewDB(t *testing.T) {
	// dbConfig := &DBConfig{
	// 	Server:    "192.168.7.7",
	// 	UserName:  "feiborate",
	// 	Password:  "feiborate",
	// 	DBName:    "feiborate",
	// 	DebugMode: true,
	// }
	// db, err := NewDB(dbConfig)
	// if err != nil {
	// 	panic(err)
	// }
	// db.Get("feifei")
	// 第一种调用方法
	sum := sha256.Sum256([]byte("hello world\n"))
	fmt.Printf("%x\n", sum)
}
