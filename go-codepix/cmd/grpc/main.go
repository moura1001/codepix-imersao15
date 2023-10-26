package main

import (
	"github.com/moura1001/codepix/infra/db"
	"github.com/moura1001/codepix/service/grpc"
)

/*
*

	evans example:
		evans -r repl
		package github.com.moura1001.codepix
		service PixService
		call RegisterPixKe

*
*/
func main() {
	database := db.GetDBConnection()

	grpc.StartGrpcServer(database, 50051)
}
