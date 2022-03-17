package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ReplaceText struct {
	OldText string //需要替换的文本
	NewText string //新的文本
}

type ReplaceHelper struct {
	Root        string        //根目录
	ReplaceList []ReplaceText //替换数组
}

func main() {
	app := gin.Default()
	// http://loclahost:8080/common-java-deploy
	app.Handle("GET", "/common-java-deploy", func(context *gin.Context) {

		fmt.Println(context.FullPath())
		filename := context.DefaultQuery("file", "common-java-deployment.yaml-template")
		image := context.DefaultQuery("image", "hello")
		Project := context.DefaultQuery("Project", "hello")
		port := context.DefaultQuery("port", "80")
		fmt.Println(filename)
		fmt.Println(image)
		fmt.Println(Project)
		fmt.Println(port)
		//下载文件
		downloadyaml(filename)
		//替换文件
		replacelist := []ReplaceText{{OldText: "##IMAGE##", NewText: image},
			{OldText: "##PROJECT##", NewText: Project},
			{OldText: "##CPORT##", NewText: port}}
		helper := ReplaceHelper{
			Root:        filename,
			ReplaceList: replacelist,
		}
		err := helper.DoWrok()
		if err == nil {
			fmt.Println("done!")
		} else {
			fmt.Println("error:", err.Error())
		}
		//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		content, err := ioutil.ReadAll(file)
		fileNames := url.QueryEscape(filename) // 防止中文乱码
		context.Writer.Header().Add("Content-Type", "application/octet-stream")
		context.Writer.Header().Add("Content-Disposition", "attachment; filename=\""+fileNames+"\"")
		if err != nil {
			fmt.Println("Read File Err:", err.Error())
		} else {
			context.Writer.Write(content)
		}
	})

	app.Run()
}

func downloadyaml(file string) {
	//定义要下载的文件
	var durl = "https://git.thpyun.com/api/v4/projects/197/repository/files/k8s-yaml%2F" + file + "/raw?ref=master"
	//解析url
	_, err := url.ParseRequestURI(durl)
	if err != nil {
		panic("地址错误")
	}

	client := http.DefaultClient
	client.Timeout = time.Second * 60 //设置超时时间

	body, err := GetHttpsSkip(durl, "vCop9iKiN-59xQUTtUGk")
	//覆盖写入
	ioutil.WriteFile(file, body, 0664)

}

func GetHttpsSkip(url, token string) ([]byte, error) {

	// 创建各类对象
	var client *http.Client
	var request *http.Request
	var resp *http.Response
	var body []byte
	var err error

	//`这里请注意，使用 InsecureSkipVerify: true 来跳过证书验证`
	client = &http.Client{}
	// 获取 request请求
	request, err = http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("GetHttpSkip Request Error:", err)
		return nil, nil
	}
	// 加入 token
	request.Header.Add("PRIVATE-TOKEN", token)
	resp, err = client.Do(request)
	if err != nil {

		log.Println("GetHttpSkip Response Error:", err)
		return nil, nil
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	defer client.CloseIdleConnections()
	return body, nil
}

func (h *ReplaceHelper) DoWrok() error {

	return filepath.Walk(h.Root, h.walkCallback)

}

func (h ReplaceHelper) walkCallback(path string, f os.FileInfo, err error) error {

	if err != nil {
		return err
	}
	if f == nil {
		return nil
	}
	if f.IsDir() {
		//fmt.Pringln("DIR:",path)
		return nil
	}

	//文件类型需要进行过滤

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		//err
		return err
	}
	content := string(buf)
	newContent := strings.Replace(content, "", "", -1)
	//数组替换
	for index, replace := range h.ReplaceList {
		newContent = strings.Replace(newContent, replace.OldText, replace.NewText, -1)
		log.Println(index)
	}
	//重新写入
	ioutil.WriteFile(path, []byte(newContent), 0)

	return err
}
