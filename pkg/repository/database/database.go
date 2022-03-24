package database

import (
	"database/sql"
	"errors"
	"time"

	// mysql "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 数据库单例
var _db *db

var (
	ErrNotFound = errors.New("根据条件未找到值！")
)

type iDb interface {
	dbInit
	// dbError
}

type dbInit interface {
	// 初始化数据库连接池
	initPool(c *gorm.Config) error
}

type dbError interface {
	GetErr() error
}

type db struct {
	Db      *gorm.DB
	Records []string
}

// 初始化并返回一个Db操作实例
func GetDb(c *gorm.Config) (*db, error) {
	// 返回数据库类
	var err error
	if _db == nil {
		_db = &db{}
		if c == nil {
			err = _db.initPool(&gorm.Config{
				PrepareStmt: false,
				Logger:      Default.LogMode(logger.Error),
			})
		} else {
			err = _db.initPool(c)
		}

	}
	return _db, err
}

// 初始化数据库连接池
func (d *db) initPool(c *gorm.Config) error {
	dsn := "root:123456@tcp(127.0.0.1:3306)/tudou_list?parseTime=true&loc=Asia%2FShanghai"
	sqlDb, err := sql.Open("mysql", dsn)

	if err != nil {
		return err
	}

	sqlDb.SetConnMaxLifetime(time.Minute * 3)
	sqlDb.SetMaxOpenConns(10) // 设置空闲连接池中连接的最大数量
	sqlDb.SetMaxIdleConns(10) // 设置打开数据库连接的最大数量

	// 初始化gorm连接池
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), c)

	if err != nil {
		return err
	} else {
		d.Db = db
	}
	return nil
}

// 获取错误
// func (d *db) GetErr() error {
// 	return d.err
// }
