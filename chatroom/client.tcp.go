package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 与server建立链接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
    if err != nil {
        fmt.Println("dial 127.0.0.1:20000 err: ", err)
        return
    }
    // 发送数据
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("msg: ")
        msg, _ := reader.ReadString('\n')
        if msg == "exit" {
            break
        }
        conn.Write([]byte(msg))
    }
    conn.Close()
}