package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//监听
	listen, err := net.Listen("tcp", ":666")
	if err != nil {
		fmt.Println("net.listen.error:", err)
		return
	}
	defer listen.Close()
	//阻塞 等待用户连接
	conn, err1 := listen.Accept()
	if err1 != nil {
		fmt.Println("listen.accept.error:", err1)
		return
	}
	//读取对方发送的文件名
	defer conn.Close()
	buf := make([]byte, 1024)
	var n int
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("conn.read.error:", err2)
		return
	}

	fileName := string(buf[:n])
	//回复ok
	conn.Write([]byte("ok"))
	//接收文件内容
	Recv(fileName, conn)
}
func Recv(fileName string, conn net.Conn) {
	//新建文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.create.error:", err)
		return
	}

	buf := make([]byte, 1024*4)
	//接收文件
	for {
		n, err := conn.Read(buf) //接收
		if err != nil {
			if err == io.EOF {
				fmt.Println("接收完成")
			} else {
				fmt.Println("conn.read.err:", err)
			}
			return
		}
		f.Write(buf[:n])
	}

}
