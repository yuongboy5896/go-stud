package Controller

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

type DockerPush struct {
}

func (dockerPush *DockerPush) Router(engine *gin.Engine) {
	engine.GET("/api/push:imageurl", dockerPush.GetToken)
}

func (dockerPush *DockerPush) GetToken(context *gin.Context) {
	imageurl := context.Param("imageurl")
	harborurl := strings.Replace(imageurl, "192.168.48.36", "172.20.4.89:8899", -1)
	//
	Command := fmt.Sprintf("docker pull %s ", imageurl)
	output, err := exec.Command("cmd", "-c", Command).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
	//
	Command = fmt.Sprintf("docker tag %s  %s", imageurl, harborurl)
	output, err = exec.Command("cmd", "-c", Command).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
	//
	Command = fmt.Sprintf("docker push %s", harborurl)
	output, err = exec.Command("cmd", "-c", Command).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
	//
	Command = fmt.Sprintf("docker rmi %s", harborurl)
	output, err = exec.Command("cmd", "-c", Command).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
	//
	Command = fmt.Sprintf("docker rmi %s", imageurl)
	output, err = exec.Command("cmd", "-c", Command).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}
