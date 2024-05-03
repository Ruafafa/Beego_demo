package controllers

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type UploadController struct {
	web.Controller
}

func (c *UploadController) Post() {
	// 获取上传的文件
	f, h, err := c.GetFile("file")
	if err != nil {
		c.Ctx.WriteString("文件上传失败" + err.Error())
		return
	}
	defer f.Close()

	// 从配置中读取阿里云OSS的参数
	endpoint, _ := web.AppConfig.String("aliyun.oss.endpoint")
	accessKeyId, _ := web.AppConfig.String("aliyun.oss_access.key.id")
	accessKeySecret, _ := web.AppConfig.String("aliyun.oss.access.key.secret")
	bucketName, _ := web.AppConfig.String("aliyun.oss.bucket.name")

	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		logs.Informational("创建OSSClient实例失败: ", err)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		logs.Informational("获取存储空间失败: ", err)
	}

	// 上传文件。
	err = bucket.PutObject(h.Filename, f)
	if err != nil {
		logs.Informational("上传文件失败: ", err)
	}

	// 获取文件的外链
	fileUrl := "https://" + bucketName + "." + endpoint + "/" + h.Filename

	// 创建一个包含状态码、消息和文件外链的map
	response := map[string]interface{}{
		"status":  200,
		"message": "文件上传成功",
		"url":     fileUrl,
	}
	logs.Informational("文件上传成功: ", fileUrl)
	// 设置JSON响应
	c.Data["json"] = response
	err = c.ServeJSON()
}
