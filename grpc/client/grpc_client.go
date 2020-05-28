package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"time"
	meteorite "yudhistira.dev/demo/meteorite/grpc/client/yudhistira.dev/demo/meteorite/grpc/api"
)

func main() {
	conn := CreateConnection()
	defer conn.Close()

	client := meteorite.NewMeteoriteLandingsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	meteorites, err := client.GetMeteorite(ctx, &empty.Empty{})
	handleError(err)
	log.Println(meteorites)
}

func CreateConnection() *grpc.ClientConn {
	dialOptions := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("localhost:8091", dialOptions...)
	handleError(err)
	return conn
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
