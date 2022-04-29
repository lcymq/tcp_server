package main

import (
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
    // 与客户端通信
    var msg = [128]byte{}
    for {
        n, err := conn.Read(msg[:]) // 这里会阻塞
        if err != nil {
            fmt.Println("read from conn failed, err: ", err)
            return
        }
        fmt.Print(conn.RemoteAddr().String(), ": ", string(msg[:n]))
    }
}