package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/qiniu/api.v7/kodo"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodocli"
	// "reflect"
	"time"
	"upload-image/utils"
)

var bm, err = cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":1,"EmbedExpiry":120}`)

//构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

type UploadController struct {
	beego.Controller
}

func (this *UploadController) Get() {
	// ===============  Cache Test
	timeoutDuration := 10 * time.Second
	if err = bm.Put("astaxie", "111", timeoutDuration); err != nil {
		fmt.Println("Save Error")
		fmt.Println(err)
	} else {
		fmt.Println("Save Success")
	}
	if v := bm.Get("astaxie"); v.(string) != "111" {
		fmt.Println("Read Error")
	} else {
		fmt.Println("Read Success")
	}
	fmt.Println(bm.Get("astaxie").(string))
	// ===============  Cache Test
	this.Data["Title"] = "Upload Image 2 QN"
	this.TplName = "upload/index.html"
}

func (this *UploadController) Post() {
	//初始化AK，SK
	conf.ACCESS_KEY = this.GetString("a-k")
	conf.SECRET_KEY = this.GetString("s-k")
	bucket := this.GetString("bucket-name")
	// fmt.Println(utils.GetAppRoot())
	fmt.Println("AK-> " + conf.ACCESS_KEY + " SK-> " + conf.SECRET_KEY)
	fmt.Println("Bucket --> " + bucket)

	// 创建一个Client
	c := kodo.New(0, nil)

	//设置上传的策略
	policy := &kodo.PutPolicy{
		Scope: bucket,
		//设置Token过期时间
		Expires: 3600,
	}
	//生成一个上传token
	token := c.MakeUptoken(policy)

	//构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	// Get the form file
	f, h, err := this.GetFile("upload-image")
	defer f.Close()
	if err != nil {
		fmt.Println("Get File Error")
	} else {
		err = this.SaveToFile("upload-image", beego.AppConfig.String("UploadPath")+h.Filename)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	var ret PutRet
	//设置上传文件的路径
	filepath := utils.GetAppRoot() + "/upload/" + h.Filename
	fmt.Println("FileName --> " + filepath)
	//调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名
	res := uploader.PutFileWithoutKey(nil, &ret, token, filepath, nil)
	//打印返回的信息
	fmt.Println(ret)
	//打印出错信息
	if res != nil {
		fmt.Println("io.Put failed:", res)
		this.Data["res"] = res
		this.TplName = "error.html"
		return
	} else {
		this.Data["tips"] = "Upload Success!"
		this.Data["filepath"] = "FilePath: "
		this.TplName = "upload/index.html"
	}

	// this.TplName = "upload/index.html"
}

// return result = {Ft2K9RTV_kSlX8KM29eLS9YC1SJq Ft2K9RTV_kSlX8KM29eLS9YC1SJq}
