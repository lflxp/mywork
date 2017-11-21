package routers

import (
	"github.com/lflxp/mywork/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
