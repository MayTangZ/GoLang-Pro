package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	//在线用户列表
	onLineMap map[string]*User
	mapLock   sync.RWMutex

	//消息广播的channel
	Message chan string
}

// 创建一个服务器的接口

func NewServer(ip string, port int) *Server {

	server := &Server{
		Ip:        ip,
		Port:      port,
		onLineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

// 监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线User
func (s *Server) ListenMessager() {
	for {
		msg := <-s.Message
		//将msg 发送给全部的在线user
		s.mapLock.Lock()
		for _, user := range s.onLineMap {
			user.C <- msg
		}

		s.mapLock.Unlock()

	}

}

func (s *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	s.Message <- sendMsg
}

func (s *Server) Handler(conn net.Conn) {

	//fmt.Println("连接建立成功")
	//用户上线 ，将用户加入到onlineMap中

	user := NewUser(conn)
	s.mapLock.Lock()
	s.onLineMap[user.Name] = user
	s.mapLock.Unlock()
	//广播当前用户上线消息
	s.BroadCast(user, "I am on line")

	//当前handler阻塞
	select {}
}

// 启动服务器的接口
func (s *Server) Start() {
	fmt.Println("服务器启动成功")

	//socket listen

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	//close listen socket
	defer listener.Close()

	//启动监听message的goroutine
	go s.ListenMessager()

	if err != nil {
		fmt.Println("net.Listen err:", err)
	}
	for {

		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue

		}

		//do handler
		go s.Handler(conn)

	}

}
