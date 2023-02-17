package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	for {
		con, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error!!!")
			continue
		}
		go process(con)

	}
}
func process(conn net.Conn) {
	defer conn.Close()
	fmt.Println("服务端:%T\n", conn)
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed err:", err)
			break
		}
		receStr := string(buf[:n])
		fmt.Println("服务端收到客户端发来的数据：", receStr)
		inputReader := bufio.NewReader(os.Stdin)
		s, _ := inputReader.ReadString('\n')
		t := strings.Trim(s, "\r\n")
		conn.Write([]byte(t))
	}
}
