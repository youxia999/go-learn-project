package main

import (
	"context"
	"go-framework-01/google_grpc_basic_example/proto"
	"go-framework-01/google_grpc_etcd_example/registercenter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Greeter  struct{}

func (greeter *Greeter) Helloservice(ctx context.Context, req *protocol.RequestMessage) (*protocol.ResponseMessage, error) {
	log.Print("received ")
	responseMessage:=protocol.ResponseMessage{Msg:"hello "+req.Name}
	return &responseMessage,nil
}
func main() {
	lis,err:=net.Listen("tcp","10.3.20.57:8084")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go registercenter.Register("192.168.172.5:2379;192.168.172.6:2379;192.168.172.7:2379","GreeterService","10.3.20.57:8084",10)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		registercenter.UnRegister("GreeterService", "10.3.20.57:8084")

		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}
	}()
	s:=grpc.NewServer()
	greeter:=Greeter{}
	protocol.RegisterGreeterServer(s,&greeter)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

