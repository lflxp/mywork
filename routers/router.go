package routers

import (
	"code.lflxp.cn/life/mywork/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
