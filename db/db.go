package db

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// import your used driver
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var Cfg = beego.AppConfig

func init() {

	var dbLink string

	dbMode := Cfg.String("db_mode")
	dbHost := Cfg.String("db_host")
	dbPort := Cfg.String("db_port")
	dbUser := Cfg.String("db_user")
	dbPassword := Cfg.String("db_password")
	dbName := Cfg.String("db_name")

	dbMaxIdleConn, err := Cfg.Int("db_max_idle_conn")
	if err != nil {
		dbMaxIdleConn = 30
	}

	dbMaxOpenConn, err := Cfg.Int("db_max_open_conn")
	if err != nil {
		dbMaxOpenConn = 200
	}

	switch dbMode {
	case "mysql":
		//register mysql driver
		orm.RegisterDriver("mysql", orm.DR_MySQL)
		dbLink = dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&allowOldPasswords=1"
		orm.RegisterDataBase("default", "mysql", dbLink, dbMaxIdleConn, dbMaxOpenConn, 30)
		name := "default"
		force := false
		verbose := false
		err := orm.RunSyncdb(name, force, verbose)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("Registered DB:", dbMode, dbLink)
	case "pgsql":
		//register postgres driver
		orm.RegisterDriver("postgres", orm.DR_Postgres)
		dbLink = dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?sslmode=disable"
		orm.RegisterDataBase("default", "postgres", dbLink, dbMaxIdleConn, dbMaxOpenConn)
		name := "default"
		force := false
		verbose := false
		err := orm.RunSyncdb(name, force, verbose)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("Registered DB:", dbMode, dbLink)
	case "sqlite":
		//register sqlite driver
		orm.RegisterDriver("sqlite", orm.DR_Sqlite)
		dbLink = dbName
		orm.RegisterDataBase("default", "sqlite3", dbLink)
		name := "default"
		force := false
		verbose := false
		err := orm.RunSyncdb(name, force, verbose)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("Registered DB:", dbMode, dbLink)
	default:
		beego.Error("DB undefined")
	}

	// orm.RegisterDataBase("default", "mysql", "root:mysql@/speakLT?charset=utf8&allowOldPasswords=1", 30)

	RunMode := Cfg.String("runmode")
	if RunMode == "dev" {
		fmt.Println("SQL debug mode ON")
		orm.Debug = true
	}
}
