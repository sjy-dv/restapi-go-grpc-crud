package main

import (
	"context"
	"crudrpc/api/mcrsv"
	"log"

	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn

	conn, _ = grpc.Dial(":8081", grpc.WithInsecure())

	defer conn.Close()

	c := mcrsv.NewRpcAppClient(conn)

	postman := mcrsv.ProtoUser {
		Idx: 101,
		Username: "들어가주라",
		UserId: "들어가라라랏!",
		Password: "근데 이거 msa말곤 쓰기 힘들겠당..",
	} 

	res, _ := c.CreateUser(context.Background(), &postman)

	log.Print(res.Res)
}