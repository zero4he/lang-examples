package main

import (
	"path/filepath"
	"os"
	"fmt"
	"os/exec"
	"syscall"
)

func main() {

	jreDir := filepath.Dir(filepath.Join(filepath.Dir(os.Args[0]), "jre"))
	jreInfo, err := os.Stat(jreDir)
	if err != nil {
		fmt.Println("jre不存在")
	}
	fmt.Println(jreDir)
	fmt.Println(jreInfo.IsDir())
	//cmd, err := exec.Command("java","-jar","hello.jar").Output()
	// 为了隐藏命令窗口, 编译的时候需要添加下面的参数
	// go build -ldflags -H=windowsgui java_launcher.go
	cmd := exec.Command("java","-jar","hello.jar")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} //隐藏窗口这句也是必须的
	err = cmd.Start()
	if err != nil {
		fmt.Println("error: ",err.Error())
	} else {
		fmt.Println("OK")
	}
	//cmd:=exec.Command("javacs -jar ./hello.jar")
	/*	stdout, err := cmd.StdoutPipe()

		if err != nil {
			fmt.Println(err)
		}

		cmd.Start()
		fmt.Println(cmd.Args)

		reader := bufio.NewReader(stdout)

		//实时循环读取输出流中的一行内容
		for {
			line, err2 := reader.ReadString('\n')
			if err2 != nil || io.EOF == err2 {
				break
			}
			fmt.Println(line)
		}

		cmd.Wait()*/
}
