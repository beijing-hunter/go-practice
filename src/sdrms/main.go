package main

import (
	_ "sdrms/routers"
	_ "sdrms/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run("192.168.1.106:8080")
}
