package main

import (
	"log"
	"net"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)

	// 连接服务端并获得1个 conn
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Println("net.Dial Fail")
		return
	}

	// 给 conn 发数据
	for {
		_, err := conn.Write([]byte("hello v0.1"))
		if err != nil {
			log.Println("conn.Write Error, err=", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			log.Println("conn.Read fail, err=", err)
			return
		}

		log.Println("conn.Write succ, Get Data from Server: ", string(buf), ",cnt:", cnt)
		time.Sleep(1 * time.Second)
	}
}
