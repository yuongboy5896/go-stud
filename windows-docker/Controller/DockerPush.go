package Controller

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	Commandarry := []string{Command}
	err := Cmd("cmd", Commandarry)
	if err != nil {
		fmt.Println(err)
		return
	}

	//
	Command = fmt.Sprintf("docker tag %s  %s", imageurl, harborurl)
	Commandarry[0] = Command
	err = Cmd("cmd", Commandarry)
	if err != nil {
		fmt.Println(err)
		return
	}

	//
	Command = fmt.Sprintf("docker push %s", harborurl)
	Commandarry[0] = Command
	err = Cmd("cmd", Commandarry)
	if err != nil {
		fmt.Println(err)
		return
	}
	//
	Command = fmt.Sprintf("docker rmi %s", harborurl)
	Commandarry[0] = Command
	err = Cmd("cmd", Commandarry)
	if err != nil {
		fmt.Println(err)
		return
	}

	//
	Command = fmt.Sprintf("docker rmi %s", imageurl)
	Commandarry[0] = Command
	err = Cmd("cmd", Commandarry)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func Cmd(commandName string, params []string) error {
	cmd := exec.Command(commandName, params...)
	dir := GetCurrentDirectory()
	fmt.Println("CmdAndChangeDirToFile", dir, cmd.Args)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("cmd.StdoutPipe: ", err)
		return err
	}
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	err = cmd.Start()
	if err != nil {
		return err
	}
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}
	err = cmd.Wait()
	return err
}

func CmdAndChangeDirToShow(dir string, commandName string, params []string) error {
	cmd := exec.Command(commandName, params...)
	fmt.Println("CmdAndChangeDirToFile", dir, cmd.Args)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("cmd.StdoutPipe: ", err)
		return err
	}
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	err = cmd.Start()
	if err != nil {
		return err
	}
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}
	err = cmd.Wait()
	return err
}
func GetCurrentDirectory() string {
	//返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	//将\替换成/
	return strings.Replace(dir, "\\", "/", -1)
}
