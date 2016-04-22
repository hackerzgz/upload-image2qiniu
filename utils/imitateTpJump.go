package utils

import (
// "github.com/astaxie/beego"
)

// /**
//  * 成功跳转
//  */
// func J2Success(this *beego.Controller, msg, url string, second int) {
// 	data := make(map[string]interface{})
// 	data["status"] = true
// 	data["title"] = "提示信息"
// 	data["msg"] = msg
// 	data["sec"] = second
// 	if url == "-1" || url == "-2" {
// 		this.Ctx.Request.Referer()
// 	}
// 	data["url"] = url
// 	this.Data["mes"] = data
// 	this.TplName = "message.html"
// }

// /**
//  * 失败跳转
//  */
// func J2Error(msg, url string, second int) {
// 	data := make(map[string]interface{})
// 	data["status"] = false
// 	data["title"] = "错误提示"
// 	data["msg"] = msg
// 	data["sec"] = second
// 	if url == "-1" || url == "-2" {
// 		this.Ctx.Request.Referer()
// 	}
// 	data["url"] = url
// 	this.Data["mes"] = data
// 	this.TplName = "message.html"
// }

// /**
//  * AJAX跳转
//  */
// func AjaxReturn(status int, msg string, data interface{}) {
// 	json := make(map[string]interface{})
// 	json["status"] = status
// 	json["msg"] = msg
// 	json["data"] = data
// 	this.Data["json"] = json
// 	this.ServeJSON(true)
// 	return
// }
