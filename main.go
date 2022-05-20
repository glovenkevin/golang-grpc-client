package main

import (
	"context"
	"fmt"
	"log"
	"os"

	chat_pb "grpc-client/pb/chat"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	log := log.New(os.Stdout, "grpc skeleton : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	serverAddr := ":8080"
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := chat_pb.NewChatServiceClient(conn)
	resp, err := client.CheckMyData(ctx, &chat_pb.CheckMyDataRequest{
		Name:      "Kevin",
		Age:       30,
		Height:    170.5,
		IsMarried: false,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Message)
}
