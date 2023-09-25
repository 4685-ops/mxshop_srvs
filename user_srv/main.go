package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"mxshop_srvs/user_srv/handle"
	"mxshop_srvs/user_srv/proto"
	"net"
)

func main() {
	// 动态的接收ip和端口
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("PORT", 50051, "端口号")
	flag.Parse()
	fmt.Println("IP：" + *IP)
	fmt.Println(*Port)
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handle.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic(err.Error())
	}
	err = server.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
