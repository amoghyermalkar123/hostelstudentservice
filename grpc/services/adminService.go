package services

import (
	"context"
	"hserver/proto/api"
	"io"
	"log"
)

type StudentAdminMessageService struct {
	admin api.AdminMessageServiceClient
}

func NewStudentAdminMessageService(ad api.AdminMessageServiceClient) *StudentAdminMessageService {
	return &StudentAdminMessageService{
		admin: ad,
	}
}

func (s *StudentAdminMessageService) GetStudentAdminMessages(req *api.StudentAdminMessageRequest, srv api.StudentAdminMessageService_GetStudentAdminMessagesServer) error {
	ctx := context.Background()
	stream, err := s.admin.MonitorAdminMessages(ctx, &api.StudentAdminMessageRequest{})

	if err != nil {
		log.Println(err)
		return err
	}

	for {
		res, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("StudentAdminMessageService.GetStudentAdminMessages failed", err)
			return err
		}

		log.Println("res:", res)

		err = srv.Send(res)

		if err != nil {
			return err
		}
	}
	return nil
}
