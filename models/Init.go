package models

import (
	"log"

	. "log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var X *xorm.Engine

func init() {

	var err error
	X, err = xorm.NewEngine("mssql", "server=.;database=PMTData;user id=sa;password=0306")
	if err != nil {
		log.Fatalf("failt to carete engine...")
	}
	X.SetMapper(core.SameMapper{})
	//日志表
	if err = X.Sync2(new(T_xiugai_log)); err != nil {
		log.Fatalf("failt T_xiugai_log to sync", err.Error())
	}

	//var xiugai=new(T_xiugai_log).Init()

}
