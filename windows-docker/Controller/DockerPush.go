package Controller

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

type DockerPush struct {
}

func (dockerPush *DockerPush) Router(engine *gin.Engine) {
	engine.GET("/api/push", dockerPush.GetToken)
}

func (dockerPush *DockerPush) GetToken(context *gin.Context) {

	imageurl := context.Query("imageurl")
	harborurl := strings.Replace(imageurl, "192.168.48.36", "172.20.4.89:8899", -1)
	//
	Command := fmt.Sprintf("docker pull %s ", imageurl)
	output, err := Cmd("cmd", "-c", Command).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
	//
	Command = fmt.Sprintf("docker tag %s  %s", imageurl, harborurl)
	output, err = Cmd("cmd", "-c", Command).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
	//
	Command = fmt.Sprintf("docker push %s", harborurl)
	output, err = Cmd("cmd", "-c", Command).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
	//
	Command = fmt.Sprintf("docker rmi %s", harborurl)
	output, err = Cmd("cmd", "-c", Command).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
	//
	Command = fmt.Sprintf("docker rmi %s", imageurl)
	output, err = Cmd("cmd", "-c", Command).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}


func Cmd(commandName string, params []string) (string, error) {
    cmd := exec.Command(commandName, params...)
    fmt.Println("Cmd", cmd.Args)
    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = os.Stderr
    err := cmd.Start()
    if err != nil {
        return "", err
    }
    err = cmd.Wait()
    return out.String(), err
