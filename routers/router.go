package routers

import (
	"github.com/astaxie/beego"
	"upload-image/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/upload", &controllers.UploadController{})
}
