package controller

import (
	"github.com/listenGrey/lucianagRpcPKG/ask"
	"google.golang.org/grpc"
	service "lucianaLLMServer/grpc"
	"net"
	"os"
)

func Response() error {
	/*creds, err := credentials.NewServerTLSFromFile(conf.CertFile, conf.KeyFile) // crt,key
	if err != nil {
		return err
	}*/
	listen, err := net.Listen("tcp", ":"+os.Getenv("LLMS_PORT")) //local ip and port
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
