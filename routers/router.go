package routers

import (
	"github.com/hymen81/JpInputToChineseConverterGo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
