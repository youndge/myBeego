package routers

import (
	"myBeego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.MainController{})
	/*Call designated Get func*/
	beego.Router("/login", &controllers.MainController{}, "get:ShowLogin;post:HandleLogin")

}
