package controller

import (
	"github.com/listenGrey/lucianagRpcPKG/ask"
	"google.golang.org/grpc"
	"lucianaLLMServer/conf"
	service "lucianaLLMServer/grpc"
	"net"
)

func Response() error {
	/*creds, err := credentials.NewServerTLSFromFile(conf.CertFile, conf.KeyFile) // crt,key
	if err != nil {
		return err
	}*/
	listen, err := net.Listen("tcp", conf.GrpcServerAddress) //local ip and port
	if err != nil {
		return err
	}

	//初始化 gRpc server
	server := grpc.NewServer(
	//grpc.Creds(creds)
	)

	ask.RegisterRequestServer(server, &service.RequestServer{})

	if err = server.Serve(listen); err != nil {
		return err
	}

	return nil
}
