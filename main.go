package main

import (
	"github.com/spf13/viper"
	"go-libr-category/configs"
	"go-libr-category/databases/connection"
	"go-libr-category/databases/migration"
	"go-libr-category/modules/category"
	"go-libr-category/utils/logger"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// initiate file configuration
	configs.Initiator()

	// initiate logger
	logger.Initiator()

	// initiate database connection
	dbConnection, _ := connection.Initiator()
	defer dbConnection.Close()

	// initiate sql migration
	migration.Initiator(dbConnection)

	// initiate rabbitmq publisher
	//rabbitMqConn := rabbitmq.Initiator()
	//defer rabbitMqConn.Channel.Close()
	//defer rabbitMqConn.Conn.Close()

	// initiate rabbitmq consumer
	//_ = rabbitMqConn.Consume()

	// initiate
	lis, err := net.Listen("tcp", viper.GetString("app.base_url")+":"+viper.GetString("app.port"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	category.Inititator(grpcServer, dbConnection)

	log.Printf("gRPC server listening on %v", lis.Addr())
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
