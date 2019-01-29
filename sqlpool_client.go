package sqlpool_client

import (
	"sync"

	"github.com/jinzhu/gorm"
)

type MysqlPoolClient struct {
	Ip   string
	Port string
	User string
	Pwd  string
	Db   string
}

var (
	instance *MysqlPoolClient
	once     sync.Once
	db       *gorm.DB
	err_db   error
)

func GetInstance(ip, port, user, pwd, db string) *MysqlPoolClient {
	once.Do(func() {
		instance = &MysqlPoolClient{
			Ip:   ip,
			Port: port,
			User: user,
			Pwd:  pwd,
			Db:   db,
		}
	})
	return instance
}

func (m *MysqlPoolClient) InitPoolClient() (err error) {
	url := m.User + ":" + m.Pwd + "@tcp(" + m.Ip + ":" + m.Port + ")/" + m.Db + "?charset=utf8&parseTime=True&loc=Local"
	db, err_db = gorm.Open("mysql", url)
	if err_db != nil {
		return err_db
	}
	return nil
}

func (m *MysqlPoolClient) GetMysqlPoolClient() (db_con *gorm.DB) {
	return db
}
