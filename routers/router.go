package routers

import (
	"github.com/hymen81/JpInputToChineseConverterGo/controllers"
	//"../controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/cars/:carId", &controllers.CarController{},"*:Get")
}
