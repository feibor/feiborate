package gormdb

import (
  "fmt"
  // 添加mysql driver
  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
  "github.com/sirupsen/logrus"
  "time"
)

// DBConfig DB的配置文件
type DBConfig struct {
  Server    string // 链接地址及端口，如 cdb-8ytuyrra.cd.tencentcdb.com:10023
  UserName  string
  Password  string
  DBName    string // database名称
  DebugMode bool   // 如果为true，则打印sql语句
}

// GetDBConnectionURL 获取连接信息
func (d *DBConfig) GetDBConnectionURL() string {
  return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True&loc=Local", d.UserName, d.Password, d.Server, d.DBName)
}

// NewDB 新增初始化数据库链接
func NewDB(config *DBConfig) (db *gorm.DB, err error) {
  mysqlURL := config.GetDBConnectionURL()
  contextLogger := logrus.WithFields(logrus.Fields{"method": "newPool",})
  contextLogger.Debug(mysqlURL)
  if db, err = gorm.Open("mysql", mysqlURL); err == nil {
    // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
    db.DB().SetMaxIdleConns(10)
    // SetMaxOpenConns sets the maximum number of open connections to the database.
    db.DB().SetMaxOpenConns(100)
    // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
    db.DB().SetConnMaxLifetime(time.Hour)
  }
  // 设置DB的日志模式
  if config.DebugMode {
    db.LogMode(true)
  }
  return
}
