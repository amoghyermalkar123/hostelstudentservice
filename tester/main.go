package main

import (
	"context"
	"hserver/proto/api"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	grpcConn, err := grpc.Dial(":3000", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	studentAdminMsgService := api.NewStudentAdminMessageServiceClient(grpcConn)

	ctx := context.Background()
	stream, err := studentAdminMsgService.GetStudentAdminMessages(ctx, &api.StudentAdminMessageRequest{})

	if err != nil {
		log.Println("client tester GetStudentAdminMessages failed", err)
		return
	}

	for {
		res, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			}

			log.Println("This", err)
			break
		}

		log.Println("wohoo")
		log.Println(res.Text)
	}
}
