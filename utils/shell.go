package utils

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"io"
	"os/exec"
)

func ExecCommand(commandName string, params ...string) (commanStr string, err error) {
	//函数返回一个*Cmd，用于E使用给出的参数执行name指定的程序
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	log.Println(cmd.Args)

	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		return
	}

	cmd.Start()
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)

	commanStr = ""
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		log.Println(line)
		if commanStr != "" {
			commanStr += "\n"
		}
		commanStr += line
	}

	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
	cmd.Wait()
	return
}
