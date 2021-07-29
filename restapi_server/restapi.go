package main

import (
	"context"
	"crudrpc/api/mcrsv"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

type Request struct {
	Idx      int    `json:"idx"`
	Username string `json:"username"`
	UserId   string `json:"userid"`
	Password string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	req := Request{}

	_ = json.Unmarshal(body, &req)
	c := mcrsv.NewRpcAppClient(conn)

	postman := mcrsv.ProtoUser{
		Idx:      int32(req.Idx),
		Username: req.Username,
		UserId:   req.UserId,
		Password: req.Password,
	}

	rows, _ := c.CreateUser(context.Background(), &postman)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rows)
}

func main() {
	// 그럼 별도 서버는 어떻게 연결하지? 프로젝트는 공유하는건가?
	// 어차피 이렇게해도 나눠서만 배포하면 msa형식이 되니까?
	conn, _ = grpc.Dial(":8081", grpc.WithInsecure())

	defer conn.Close()

	//c := mcrsv.NewRpcAppClient(conn)

	r := mux.NewRouter()

	r.HandleFunc("/create", CreateUser).Methods("POST")

	http.ListenAndServe(":8082", r)

}
