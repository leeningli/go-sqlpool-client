package sqlpool_client

import (
	"sync"

	"github.com/jinzhu/gorm"
)

type MysqlPoolClient struct {
}

var (
	instance *MysqlPoolClient
	once     sync.Once
	db       *gorm.DB
	err_db   error
)

func GetInstance() *MysqlPoolClient {
	once.Do(func() {
		instance = &MysqlPoolClient{}
	})
	return instance
}

func (m *MysqlPoolClient) InitPoolClient() (issucc bool, err error) {
	db, err_db = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/nanili?charset=utf8&parseTime=True&loc=Local")
	if err_db != nil {
		return false, err_db
	}
	return true, nil
}

func (m *MysqlPoolClient) GetMysqlPoolClient() (db_con *gorm.DB) {
	return db
}
