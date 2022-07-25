package server

import (
	"fmt"
	"log"
	"net"

	"github.com/xm-tech/tcpserver/iface"
)

// implements the IServer interface
type Server struct {
	IPVersion string

	Name string

	IP string

	Port int
}

func NewServer() iface.IServer {
	s := &Server{
		IPVersion: "tcp4",
		Name:      "TcpServer",
		IP:        "127.0.0.1",
		Port:      9999,
	}
	return s
}

func (self *Server) Start() {
	// 获取1 TCP 地址
	tcpAddr, err := net.ResolveTCPAddr(self.IPVersion, fmt.Sprintf("%s:%d", self.IP, self.Port))
	if err != nil {
		log.Fatalf("Server Start Fail,Resolve Tcp Addr Error:%s", err)
	}

	// 监听服务端套接字
	listener, err := net.ListenTCP(self.IPVersion, tcpAddr)
	if err != nil {
		log.Fatalf("Server.Start fail, ListenTcp Error:%s", err)
	}

	log.Println("Server.Start,listen succ,tcpAddr=", tcpAddr)

	// 阻塞等待客户端链接，处理链接业务
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Println("Server.Start fail, AcceptTCP Error:", err)
			panic(err)
		}

		log.Println("Server.Start, accepted one conn:", conn)

		go func() {
			for {
				// a simple echo
				buf := make([]byte, 512)
				cnt, err := conn.Read(buf)
				if err != nil {
					log.Println("conn Read Fail,err=", err)
					conn.Close()
					break
				}

				log.Println("Read data:", string(buf), ",cnt:", cnt, " from client:", conn)

				_, err = conn.Write(buf[:cnt])
				if err != nil {
					log.Println("conn Write fail,err=", err)
					conn.Close()
					break
				}
			}
		}()
	}
}

func (self *Server) Stop() {

}

func (self *Server) Run() {
	go self.Start()
	// TODO 做些服务启动后的额外动作
	select {}
}
