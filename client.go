package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//主动连接服务端
	conn, err1 := net.Dial("tcp", ":666")
	if err1 != nil {
		fmt.Println("net.dial.error:", err1)
	}

	defer conn.Close()

	//提示输入文件
	fmt.Println("请输入文件名称")
	var path string
	fmt.Scan(&path)

	//获取文件名
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.stat.error:", err)
		return
	}

	//给接收方发送文件名
	_, err2 := conn.Write([]byte(info.Name()))
	if err2 != nil {
		fmt.Println("conn.write.error:", err2)
		return
	}

	//接收对方回复，如果是ok就开始传输文件
	var n int
	buf := make([]byte, 1024)
	n, err3 := conn.Read(buf)
	if err3 != nil {
		fmt.Println("conn.read.error", err3)
		return
	}

	if "ok" == string(buf[:n]) {
		//发送文件内容
		Send(path, conn)
	}

}
func Send(path string, conn net.Conn) {
	//只读方式打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.open.error:", err)
		return
	}

	defer f.Close()

	buf := make([]byte, 1024*4)

	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件传输完成")
			} else {
				fmt.Println("f.read.error:", err)
			}
			return
		}
		conn.Write(buf[:n])

	}

}
