package main

import (
	"context"
	"encoding/json"
	"fmt"
	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	meteorite "yudhistiraaryarukmana.org/demo/meteorite/grpc/server/yudhistiraaryarukmana.org/demo/meteorite/grpc/api"
)

const meteoriteJsonFile = "y77d-th95.json"
var meteorites []meteorite.MeteoriteLanding

func init() {
	currentDir, err := os.Getwd()
	handleError(err)

	path := filepath.Join(currentDir, "data", meteoriteJsonFile)
	log.Println(path)

	bytes, err := ioutil.ReadFile(path)
	handleError(err)

	handleError(json.Unmarshal(bytes, &meteorites))
	log.Printf("GRPC server initialization complete")
}

type MeteoriteService struct {
	meteorite.UnimplementedMeteoriteLandingsServiceServer
}

func (s *MeteoriteService) GetMeteorite(ctx context.Context, empty *empty.Empty) (*meteorite.MeteoriteLandingList, error) {
	ptrs := make([]*meteorite.MeteoriteLanding, len(meteorites))
	for i, v := range meteorites {
		ptrs[i] = &v
	}
	return &meteorite.MeteoriteLandingList{MeteoriteLanding: ptrs}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8091))
	handleError(err)
	grpcServer := grpc.NewServer()
	meteoriteService := &MeteoriteService{}
	meteorite.RegisterMeteoriteLandingsServiceServer(grpcServer, meteoriteService)
	log.Printf("Starting grpc server at 8091")
	handleError(grpcServer.Serve(lis))
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
