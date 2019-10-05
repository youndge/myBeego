package main

import (
	_ "myBeego/models"
	_ "myBeego/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
