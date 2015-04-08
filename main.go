package main

import (
	_ "github.com/bakanis/simpleauth"

	_ "github.com/bakanis/simpleauth-sample/db"

	_ "github.com/bakanis/simpleauth-sample/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
