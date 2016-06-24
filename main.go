package main

import (
	"github.com/astaxie/beego"
	_ "upload-image2qiniu/routers"
)

func main() {
	beego.Run()
}
