package service

import (
	"errors"
	"fmt"
	"log"

	"ch01/v2/model"

	"github.com/go-xorm/xorm"
)

var DbEngin *xorm.Engine

func init() {
	drivename := "mysql"
	DsName := "root:914082247@QQ.com@(localhost:3306)/chat?charset=utf8"
	err := errors.New("")
	DbEngin, err = xorm.NewEngine(drivename, DsName)
	if nil != err && "" != err.Error() {
		log.Fatal(err.Error())
	}
	//是否显示SQL语句
	DbEngin.ShowSQL(true)
	//数据库最大打开的连接数
	DbEngin.SetMaxOpenConns(2)

	//自动User
	DbEngin.Sync2(new(model.User))
	//DbEngin = dbengin
	fmt.Println("init data base ok")
}
