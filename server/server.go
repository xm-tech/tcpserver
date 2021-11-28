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

// 主要的网络事件循环
func (ch *connHanlder) Handle() {
	go func() {
		defer ch.conn.Close()

		// business logic
		fmt.Println("connHanlder.Handle")
	}()
}

type wrappedConn struct {
	net.Conn
}

func (wc *wrappedConn) Close() {
	wc.Close()
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
