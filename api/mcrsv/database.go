package mcrsv

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



func (server *Server) InitDB(DB_DRIVER, DB_USER, DB_PASSWORD, DB_PORT, DB_HOST, DB_NAME string) {
	var err error

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME, "Asia%2FSeoul")
	server.DB, err = gorm.Open(DB_DRIVER, DBURL)

	if err != nil {
		log.Fatal("db error : ", err)
	}

	//not migrate
	server.DB.SingularTable(true)
	server.DB.DB().SetMaxIdleConns(10)
	server.DB.DB().SetMaxOpenConns(300)
	server.DB.DB().SetConnMaxIdleTime(10)
	server.DB.DB().SetConnMaxLifetime(6000)
}

/*
func (server *Server) Run(port string) {
	fmt.Printf("server on port %s", port)

	grpcServer := grpc.NewServer()

	lis, _ := net.Listen("tcp", port)

	RegisterRpcAppServer(grpcServer, server)

	err := grpcServer.Serve(lis)

	if err != nil {
		log.Fatal(err)
	}

}
*/