package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	pb "example.com/grpcpb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type UserManagementServer struct { //create struct of grpc
	pb.UnimplementedUserManagementServer
}

func (*UserManagementServer) CreateMsg(req *pb.Msg, stream pb.UserManagement_CreateMsgServer) error {

	name := req.GetTkn()
	fmt.Println("name %v", name)
	//use wait group to allow process to be concurrent
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("i %v", i)
			result := "Hello " + name + " number " + strconv.Itoa(i)
			res := &pb.Returned{Str: result}
			stream.Send(res)
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //open listener connection to client
	pb.RegisterUserManagementServer(s, &UserManagementServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
