package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/bndr/gojenkins"
)

func main() {
	ctx := context.Background()
	jenkins := gojenkins.CreateJenkins(nil, "http://192.168.48.37:8080/", "thpower", "1qaz2wsx")
	// Provide CA certificate if server is using self-signed certificate
	// caCert, _ := ioutil.ReadFile("/tmp/ca.crt")
	// jenkins.Requester.CACert = caCert
	_, err := jenkins.Init(ctx)
	//
	if err != nil {
		log.Printf("连接Jenkins失败, %v\n", err)
		return
	}
	log.Println("Jenkins连接成功")
	file, err := os.Open("config.xml")
	if err != nil {
		fmt.Println("读文件失败", err)
		return
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("读内容失败", err)
		return
	}
	fmt.Println(string(content))
	configString := string(content)

	configString = strings.Replace(configString, "##GITURL##", "代码地址", -1)       //代码地址
	configString = strings.Replace(configString, "##MODULENAME##", "模块中文描述", -1) //模块中文描述
	configString = strings.Replace(configString, "##BRANCH##", "代码分支", -1)       //代码分支
	configString = strings.Replace(configString, "##DEPLOY##", "模块英文名称", -1)     //模块英文名称
	configString = strings.Replace(configString, "##ENV##", "环境地址", -1)          //环境地址
	configString = strings.Replace(configString, "##NAMESPACE##", "命名空间", -1)    //命名空间
	configString = strings.Replace(configString, "##IMAGEULR##", "上传镜像地址", -1)   //上传镜像地址

	var del bool
	getjob, err := jenkins.GetJob(ctx, "test-java-java")

	if getjob != nil {
		del, err = jenkins.DeleteJob(ctx, "test-java-java")
		if err != nil && !del {
			panic(err)
		}

	}
	job, err := jenkins.CreateJobInFolder(ctx, configString, "test-java-java")
	if err != nil {
		panic(err)
	}

	if job != nil {
		fmt.Println("Job has been created in child folder")
	}
}
