package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"kitex-multi-protocol/internal/protocol"
	"kitex-multi-protocol/utils"
)

func main() {
	// 创建 UserService 实例
	service := &utils.UserServiceImpl{}

	// 监听端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080...")

	for {
		// 接受新连接
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		// 这里简单假设前 4 个字节能判断协议，实际情况需要更复杂的逻辑
		buf := make([]byte, 4)
		_, err = conn.Read(buf)
		if err != nil {
			log.Printf("Failed to read protocol bytes: %v", err)
			conn.Close()
			continue
		}

		var protocolStr string
		// 简单示例，根据字节内容判断协议
		if string(buf) == "HTTP" {
			protocolStr = "HTTP"
		} else {
			// 假设其他情况为 Thrift
			protocolStr = "Thrift"
		}

		// 创建处理程序
		handler := protocol.CreateHandler(protocolStr, service)
		if handler == nil {
			log.Printf("Unsupported protocol: %s", protocolStr)
			conn.Close()
			continue
		}

		// 处理连接
		go func() {
			defer conn.Close()
			if err := handler.Handle(context.Background(), conn); err != nil {
				log.Printf("Error handling connection: %v", err)
			}
		}()
	}
}
