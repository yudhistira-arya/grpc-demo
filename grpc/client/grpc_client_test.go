package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"sync"
	"testing"
	"time"
	meteorite "yudhistiraaryarukmana.org/demo/meteorite/grpc/client/yudhistiraaryarukmana.org/demo/meteorite/grpc/api"
)

var result *meteorite.MeteoriteLandingList

func Benchmark1GrpcRequest(b *testing.B) {
	var conn = CreateConnection()
	defer conn.Close()

	client := meteorite.NewMeteoriteLandingsServiceClient(conn)

	var response *meteorite.MeteoriteLandingList
	var err error
	for n := 0; n < b.N; n++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		defer cancel()

		response, err  = client.GetMeteorite(ctx, &empty.Empty{})
		handleError(err)
	}
	result = response
}

func Benchmark10GrpcRequest(b *testing.B) {
	var conn = CreateConnection()
	defer conn.Close()

	concurrentRequest := 10

	client := meteorite.NewMeteoriteLandingsServiceClient(conn)
	var response *meteorite.MeteoriteLandingList
	var err error
	for n := 0; n < b.N; n++ {
		var wg sync.WaitGroup
		wg.Add(concurrentRequest)
		for c := 0; c < concurrentRequest; c++ {
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
				defer cancel()

				defer wg.Done()
				response, err  = client.GetMeteorite(ctx, &empty.Empty{})
				handleError(err)
			}()
		}
		wg.Wait()
	}
	result = response
}

func Benchmark100GrpcRequest(b *testing.B) {
	var conn1 = CreateConnection()
	defer conn1.Close()

	concurrentRequest := 100
	client1 := meteorite.NewMeteoriteLandingsServiceClient(conn1)

	var response *meteorite.MeteoriteLandingList
	var err error
	for n := 0; n < b.N; n++ {
		var wg sync.WaitGroup
		wg.Add(concurrentRequest)
		for c := 0; c < concurrentRequest; c++ {
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
				defer cancel()
				defer wg.Done()

				response, err  = client1.GetMeteorite(ctx, &empty.Empty{})
				handleError(err)
			}()
		}
		wg.Wait()
	}
	result = response
}


// 1
// grpc:    1069319 ns /op
// restful: 4814045 ns /op

// 10
// grpc:	 3676586 ns/op
// restful: 13565108 ns/op

// 100
// grpc:     28657121 ns/op
// restful: 112159437 ns/op


