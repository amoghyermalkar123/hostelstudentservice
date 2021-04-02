package main

import (
	"hserver/grpc/services"
	"hserver/handlers"
	"hserver/proto/api"
	"net"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/grpc"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := &handlers.Handler{}
	// Routes
	e.GET("/", h.GetStudentDetails)
	e.POST("/post_grievance/", h.PostGrievance)
	// start grpc server
	go grpcClient()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func grpcClient() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	grpcConn, err := grpc.Dial(":3032", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	serverObject := grpc.NewServer()

	adminService := api.NewAdminMessageServiceClient(grpcConn)

	service := services.NewStudentAdminMessageService(adminService)

	api.RegisterStudentAdminMessageServiceServer(serverObject, service)

	// Serve
	if e := serverObject.Serve(listener); e != nil {
		panic(e)
	}
}
