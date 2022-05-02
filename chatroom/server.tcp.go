package main

import (
	"bufio"
	"fmt"
	"net"
)

//tcp server
func main() {
    // 本地端口启动服务
    listener, err := net.Listen("tcp", "127.0.0.1:20000")
    if err != nil {
        fmt.Println("start server on 127.0.0.1:20000 failed, err: ", err)
        return
    }
    // 等待别人建立链接
    for {
        conn, err := listener.Accept() // 这里会阻塞
        if err != nil {
            fmt.Println("accept failed, err: ", err)
            return
        }
        go serverProcessor(conn)
    }
}

func serverProcessor(conn net.Conn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)
    // 与客户端通信
    for {
        msg, err := Decode(reader)
        if err != nil {
            fmt.Print(conn.RemoteAddr().String(), "err: ", err)
        }
        fmt.Print(conn.RemoteAddr().String(), ": ", string(msg))
    }
}