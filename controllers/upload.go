package controllers

import (
	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

func (c *UploadController) Get() {
	c.Data["Title"] = "Upload Image 2 QN"
	c.TplName = "upload/index.html"
}

func (c *UploadController) Post() {
	// var String AK = "m_xn4k8rHFoXmjYC0z02X-nm5yFaIAy5dw-S4Vcu"
	// var String SK = "IYKDGVTahKPIxe4kJpmxQZ48FIpKrZHEUc7crLco"
}

func Upload_Image() {

}
