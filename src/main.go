package main

import (
	fmt "fmt"
	protocol "im_protocol"
	net "net" // 网络
	"os"

	proto "github.com/golang/protobuf/proto"
)

func login(userName string, pwd string, ip string) {
	// 字符串转IP地址
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ip)
	checkErr(err)

	// 连接目标服务器
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	// 延迟函数，在进程退出后执行
	defer conn.Close()
	checkErr(err)

	// 登录请求
	loginReq := &protocol.IMLoginReq{
		UserName: userName,
		Password: pwd,
	}
	data, _ := proto.Marshal(loginReq)
	conn.Write(data) // 发送

	// 回复
	var buf [1024]byte
	var len int
	len, err = conn.Read(buf[0:])

	fmt.Printf("收到来自服务器的回复，长度：%d", len)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s \n", err.Error())
		os.Exit(1)
	}
}

func main() {
	login("xuyc", "12345", "192.168.100.185:8000")
}
