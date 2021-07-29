package api

import (
	"crudrpc/api/mcrsv"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func RunServer() {

	var err error

	var server = mcrsv.Server{}

	err = godotenv.Load()

	if err != nil {
		log.Fatalf("env error : %v", err)
	}
//DB_DRIVER,MYSQL_DB_USER, MYSQL_DB_PASSWORD,DB_PORT, MYSQL_DB_HOST,MYSQL_DB
	server.InitDB(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	grpcServer := grpc.NewServer()

	lis, _ := net.Listen("tcp", ":8081")

	mcrsv.RegisterRpcAppServer(grpcServer, &server)

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatal(err)
	}

}