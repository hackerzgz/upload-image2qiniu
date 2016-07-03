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
	"upload-image2qiniu/utils"
)

var bm, err = cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}`)

//构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

// 页面返回值变量
var ReturnValue struct {
	Flag     bool
	Title    string
	Tips     string
	FilePath string
}

type UploadController struct {
	beego.Controller
}

func (this *UploadController) Get() {
	// ===============  Cache Test
	this.Data["AKEY"], this.Data["SKEY"] = ReadCache()
	// ===============  Cache Test
	if !ReturnValue.Flag {
		setReturnValue("Upload Image 2 QiNiu", "一旦上传成功，会将你上传成功的AK以及SK进行加密缓存10min，这时候之后只要你不重新刷新页面，你依然不需要重新CV你的AK以及SK", "")
	}
	this.Data["Title"] = ReturnValue.Title
	this.Data["Tips"] = ReturnValue.Tips
	this.Data["Filepath"] = ReturnValue.FilePath
	this.TplName = "upload/index.html"
}

func (this *UploadController) Post() {
	//初始化AK，SK
	conf.ACCESS_KEY = this.GetString("a-k")
	conf.SECRET_KEY = this.GetString("s-k")
	bucket := this.GetString("bucket-name")

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
	// 简易上传
	ret, res := SimpleUploadFile(bucket, h.Filename)
	//打印返回的Hash,Key信息
	fmt.Println(ret)
	//打印出错信息
	if res != nil {
		Jump2ErrPage(this, res)
	} else {
		fmt.Println("I am In!")
		// 将AK以及SK缓存到Cache模块中
		setCache(this.GetString("a-k"), this.GetString("s-k"))
		setReturnValue("Upload Image 2 QiNiu", "Upload Success!", GetFilePath()+ret.Key)
		ReturnValue.Flag = true
		go setReturnStatus()()
		// 303防止POST提交之后刷新出现重新提交操作
		this.Redirect("/upload", 303)
	}
}

// 设置页面返回值提示
func setReturnValue(title, tips, filepath string) {
	ReturnValue.Title, ReturnValue.Tips, ReturnValue.FilePath = title, tips, filepath
}

// 设置过时页面返回信息，5分钟后取消返回上传图片的公网地址
func setReturnStatus() {
	time.Sleep(5 * time.Minute)
	ReturnValue.Flag = false
}

// 简易上传
func SimpleUploadFile(bucket, fileName string) (putRet PutRet, err error) {
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

	var ret PutRet
	//设置上传文件的路径
	filepath := utils.GetAppRoot() + "/upload/" + fileName
	fmt.Println("FileName --> " + filepath)
	//调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名
	res := uploader.PutFile(nil, &ret, token, fileName, filepath, nil)
	return ret, res
}

// 读取AK/SK缓存
func ReadCache() (ak, sk string) {
	return bm.Get("AK").(string), bm.Get("SK").(string)
}

// 设置AK/SK缓存
func setCache(ak, sk string) {
	if err := bm.Put("AK", ak, 10*time.Minute); err != nil {
		fmt.Println("AK Cache Faile!")
	}
	if err := bm.Put("SK", sk, 10*time.Minute); err != nil {
		fmt.Println("SK Cache Faile!")
	}
}

// 跳转至错误页面
func Jump2ErrPage(this *UploadController, err error) {
	this.Data["res"] = "文件上传出错了：" + err.Error()
	this.TplName = "error.html"
	return
}

// 获取设置中的公网地址
func GetFilePath() (filePath string) {
	fmt.Println(beego.AppConfig.String("PublicPath"))
	if filePath = beego.AppConfig.String("PublicPath"); filePath == "" {
		filePath = "FilePath: <公网地址>"
	}
	filePath += "/"
	return
}
