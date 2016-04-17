package controllers

import (
	"github.com/astaxie/beego"
	// . "github.com/qiniu/api.v7/conf"
	"fmt"
	"log"
)

type UploadController struct {
	beego.Controller
}

func (c *UploadController) Get() {
	c.Data["Title"] = "Upload Image 2 QN"
	c.TplName = "upload/index.html"
}

func (c *UploadController) Post() {
	AK := c.GetString("a-k")
	SK := c.GetString("s-k")
	log.Println("AK-> " + AK + " SK-> " + SK)
	f, h, err := c.GetFile("upload-image")
	defer f.Close()
	if err != nil {
		fmt.Println("Get File Error")
	} else {
		err = c.SaveToFile("upload-image", beego.AppConfig.String("UploadPath")+h.Filename)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	c.TplName = "upload/index.html"
}

func Upload_Image() {

}
