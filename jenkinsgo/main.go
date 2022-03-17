package main

import (
	"context"
	"log"

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
	Job, err := jenkins.CopyJob(ctx, "email", "email-test")
	//
	if err != nil {
		log.Printf("Job 创建失败, %v\n", err)
		return
	}
	log.Printf("Job 创建成功, %v\n", Job)

}
