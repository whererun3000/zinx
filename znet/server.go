package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Server struct {
	Name    string
	NetWork string
	IP      string
	Port    int
}

func (s *Server) Start() {
	go func() {
		addr := fmt.Sprintf("%s:%d", s.IP, s.Port)
		fmt.Printf("zinx server start, listen in %s\n", addr)

		ln, err := net.Listen(s.NetWork, addr)
		if err != nil {
			fmt.Printf("zinx server listen err: %v\n", err)
			return
		}

		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Printf("zinx listener accept err: %v\n", err)
				return
			}

			go func(conn net.Conn) {
				for {
					buf := make([]byte, 512)
					n, err := conn.Read(buf)
					if err != nil {
						fmt.Printf("zinx server read from conn err: %v\n", err)
						return
					}

					fmt.Printf("zinx server read from client: %s\n", buf[:n])

					_, err = conn.Write(buf[:n])
					if err != nil {
						fmt.Printf("zinx server echo err: %v\n", err)
						return
					}
				}
			}(conn)
		}
	}()
}

func (s *Server) Stop() {
	panic("implement me")
}

func (s *Server) Serve() {
	s.Start()
	select {}
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:    name,
		NetWork: "tcp",
		IP:      "0.0.0.0",
		Port:    8999,
	}
}
