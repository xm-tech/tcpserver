package server

import (
	"fmt"
	"net"
)

type Server struct {
	listener net.Listener
}

type connHanlder struct {
	conn *wrappedConn
}

// 单 TCP 连接网络事件循环处理逻辑
func (ch *connHanlder) Handle() {
	// business logic TODO
	fmt.Println("connHanlder.Handle")
	go func() {
		defer ch.conn.Close()
		for {
			// 读数据
			buff := make([]byte, 512)
			ch.conn.Read(buff)

			// 解码
			fmt.Println(string(buff))
		}
	}()
}

type wrappedConn struct {
	net.Conn
}

func (wc *wrappedConn) Close() {
	wc.Conn.Close()
}

func (self *Server) Exec() {
	self.Listen()
}

func (self *Server) Listen() {
	for {
		conn, err := self.listener.Accept()
		if err != nil {
			panic(err)
		}
		ch := connHanlder{
			conn: &wrappedConn{Conn: conn},
		}
		ch.Handle()
	}
}
